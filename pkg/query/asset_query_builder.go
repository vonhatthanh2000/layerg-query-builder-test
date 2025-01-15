package query

import (
	"asset-query/internal/response"
	"context"
	"errors"
	"fmt"
	"time"

	masterDbCommon "github.com/u2u-labs/go-layerg-common/masterdb"
)

type Pagination[T any] struct {
	Page       int    `json:"page"`              // Current page number
	Limit      int    `json:"limit"`             // Number of items per page
	TotalItems int64  `json:"totalItems"`        // Total number of items available
	TotalPages int64  `json:"totalPages"`        // Total number of pages
	Holders    *int64 `json:"holders,omitempty"` // Optional holder field
	Data       []T    `json:"data"`              // The paginated items (can be any type)
}

type assetQueryBuilderParam struct {
	chainId       int32
	collectionId  *string
	tokenIds      *[]string
	owner         *string
	createdAtFrom *time.Time
	createdAtTo   *time.Time
	page          *int
	limit         *int
	offset        *int
	config        *masterDbConfig
}

func (b *assetQueryBuilderParam) getHttpClient() *HttpClient {
	client := NewHttpClient(b.config.masterDbUrl)
	return client
}

// WithChainId implements AssetQueryBuilder.
func (b *assetQueryBuilderParam) getCollectionType() masterDbCommon.CollectionType {
	client := b.getHttpClient()

	var response response.HTTPResponse[masterDbCommon.CollectionResponse]
	path := fmt.Sprintf("/chain/%d/collection/%s", b.chainId, *b.collectionId)
	err := client.DoRequest(context.Background(), "GET", path, nil, &response)
	if err != nil {
		fmt.Println("err", err)
		return masterDbCommon.CollectionType("")
	}

	return response.Data.Type
}

// WithChainId implements AssetQueryBuilder.
func (b *assetQueryBuilderParam) WithChainId(chainId int32) AssetQueryBuilder {
	b.chainId = chainId
	return b
}

// WithCollectionId implements AssetQueryBuilder.
func (b *assetQueryBuilderParam) WithCollectionId(collectionId string) AssetQueryBuilder {
	b.collectionId = &collectionId
	return b
}

// WithCreatedAtFrom implements AssetQueryBuilder.
func (b *assetQueryBuilderParam) WithCreatedAtFrom(createdAtFrom time.Time) AssetQueryBuilder {
	b.createdAtFrom = &createdAtFrom
	return b
}

// WithCreatedAtTo implements AssetQueryBuilder.
func (b *assetQueryBuilderParam) WithCreatedAtTo(createdAtTo time.Time) AssetQueryBuilder {
	b.createdAtTo = &createdAtTo
	return b
}

// WithLimit implements AssetQueryBuilder.
func (b *assetQueryBuilderParam) WithLimit(limit int) AssetQueryBuilder {
	b.limit = &limit
	return b
}

// WithOwner implements AssetQueryBuilder.
func (b *assetQueryBuilderParam) WithOwner(owner string) AssetQueryBuilder {
	b.owner = &owner
	return b
}

// WithPage implements AssetQueryBuilder.
func (b *assetQueryBuilderParam) WithPage(page int) AssetQueryBuilder {
	b.page = &page
	return b
}

// WithTokenId implements AssetQueryBuilder.
func (b *assetQueryBuilderParam) WithTokenIds(tokenIds []string) AssetQueryBuilder {
	b.tokenIds = &tokenIds
	return b
}

type AssetQueryFunction interface {
	GetAssetQueryBuilder() (*assetQueryBuilderParam, error)
	GetPaginatedAsset() (any, error)
}

type AssetQueryBuilder interface {
	WithChainId(chainId int32) AssetQueryBuilder
	WithCollectionId(collectionId string) AssetQueryBuilder
	WithTokenIds(tokenIds []string) AssetQueryBuilder
	WithOwner(owner string) AssetQueryBuilder
	WithCreatedAtFrom(createdAtFrom time.Time) AssetQueryBuilder
	WithCreatedAtTo(createdAtTo time.Time) AssetQueryBuilder
	WithPage(page int) AssetQueryBuilder
	WithLimit(limit int) AssetQueryBuilder
	Build() AssetQueryFunction
}

func (b *assetQueryBuilderParam) Build() AssetQueryFunction {
	// Set default values if not provided
	defaultPage := 1
	defaultLimit := 10

	if b.page != nil {
		if *b.page < 1 {
			return nil
		}
		defaultPage = *b.page
	}

	if b.limit != nil {
		if *b.limit > 100 {
			return nil
		}
		defaultLimit = *b.limit
	}

	// Calculate offset
	offset := (defaultPage - 1) * defaultLimit
	b.offset = &offset
	return b
}

func (b *assetQueryBuilderParam) GetAssetQueryBuilder() (*assetQueryBuilderParam, error) {
	if b.chainId == 0 {
		return nil, errors.New("chainId is required")
	}
	return b, nil
}

func (b *assetQueryBuilderParam) getLocalAssetQuery() (any, error) {
	collectionType := b.getCollectionType()
	fmt.Println("collectionType", collectionType)

	filterConditions := b.getFilterConditions()

	switch collectionType {
	case masterDbCommon.CollectionTypeERC721:
		totalAssets, _, err := CountItemsWithFilter(b.config.localDb, "erc_721_collection_assets", filterConditions)
		if err != nil {
			return nil, err
		}

		assets, _ := QueryWithDynamicFilter[masterDbCommon.Erc721CollectionAssetResponse](b.config.localDb, "erc_721_collection_assets", *b.limit, *b.offset, filterConditions)
		return Pagination[masterDbCommon.Erc721CollectionAssetResponse]{
			Page:       *b.page,
			Limit:      *b.limit,
			TotalItems: int64(totalAssets),
			TotalPages: (int64(totalAssets) + int64(*b.limit) - 1) / int64(*b.limit),
			Data:       assets,
		}, nil

	case masterDbCommon.CollectionTypeERC1155:
		totalAssets, _, err := CountItemsWithFilter(b.config.localDb, "erc_1155_collection_assets", filterConditions)
		if err != nil {
			return nil, err
		}

		assets, _ := QueryWithDynamicFilter[masterDbCommon.Erc1155CollectionAssetResponse](b.config.localDb, "erc_1155_collection_assets", *b.limit, *b.offset, filterConditions)
		return Pagination[masterDbCommon.Erc1155CollectionAssetResponse]{
			Page:       *b.page,
			Limit:      *b.limit,
			TotalItems: int64(totalAssets),
			TotalPages: (int64(totalAssets) + int64(*b.limit) - 1) / int64(*b.limit),
			Data:       assets,
		}, nil

	case masterDbCommon.CollectionTypeERC20:
		totalAssets, _, err := CountItemsWithFilter(b.config.localDb, "erc_20_collection_assets", filterConditions)
		if err != nil {
			return nil, err
		}

		assets, _ := QueryWithDynamicFilter[masterDbCommon.Erc20CollectionAssetResponse](b.config.localDb, "erc_20_collection_assets", *b.limit, *b.offset, filterConditions)
		return Pagination[masterDbCommon.Erc20CollectionAssetResponse]{
			Page:       *b.page,
			Limit:      *b.limit,
			TotalItems: int64(totalAssets),
			TotalPages: (int64(totalAssets) + int64(*b.limit) - 1) / int64(*b.limit),
			Data:       assets,
		}, nil
	}

	return nil, nil
}

func (b *assetQueryBuilderParam) getMasterDbAsset() (any, error) {
	httpClient := b.getHttpClient()
	collectionType := b.getCollectionType()
	fmt.Println("collectionType", collectionType)

	requestBody := map[string]interface{}{
		"chainId":       b.chainId,
		"collectionId":  b.collectionId,
		"tokenIds":      b.tokenIds,
		"owner":         b.owner,
		"createdAtFrom": b.createdAtFrom,
		"createdAtTo":   b.createdAtTo,
		"page":          b.page,
		"limit":         b.limit,
		"offset":        b.offset,
	}

	switch collectionType {
	case masterDbCommon.CollectionTypeERC721:
		var response response.HTTPResponse[Pagination[masterDbCommon.Erc721CollectionAssetResponse]]
		err := httpClient.DoRequest(context.Background(), "POST", "/query-builder", requestBody, &response)
		if err != nil {
			fmt.Println("err", err)
			return nil, err
		}
		fmt.Println("response", response.Data)
		return response.Data, nil
	case masterDbCommon.CollectionTypeERC1155:
		var response response.HTTPResponse[Pagination[masterDbCommon.Erc1155CollectionAssetResponse]]
		err := httpClient.DoRequest(context.Background(), "POST", "/query-builder", requestBody, &response)
		if err != nil {
			fmt.Println("err", err)
			return nil, err
		}
		fmt.Println("response", response.Data)
		return response.Data, nil
	case masterDbCommon.CollectionTypeERC20:
		var response response.HTTPResponse[Pagination[masterDbCommon.Erc20CollectionAssetResponse]]
		err := httpClient.DoRequest(context.Background(), "POST", "/query-builder", requestBody, &response)
		if err != nil {
			fmt.Println("err", err)
			return nil, err
		}
		fmt.Println("response", response.Data)
		return response.Data, nil
	}

	return nil, nil
}

func (b *assetQueryBuilderParam) GetPaginatedAsset() (any, error) {

	if !b.config.useMasterDb {
		return b.getLocalAssetQuery()
	}
	return b.getMasterDbAsset()

}

func NewAssetQueryBuilder(config *masterDbConfig) AssetQueryBuilder {
	return &assetQueryBuilderParam{config: config}
}

func (b *assetQueryBuilderParam) getFilterConditions() map[string][]string {
	filterConditions := make(map[string][]string)

	if b.collectionId != nil {
		filterConditions["collection_id"] = []string{*b.collectionId}
	}

	if b.tokenIds != nil && len(*b.tokenIds) > 0 {
		filterConditions["token_id"] = *b.tokenIds
	}

	if b.owner != nil {
		filterConditions["owner"] = []string{*b.owner}
	}

	if b.createdAtFrom != nil {
		filterConditions["created_at_from"] = []string{b.createdAtFrom.Format(time.RFC3339)}
	}

	if b.createdAtTo != nil {
		filterConditions["created_at_to"] = []string{b.createdAtTo.Format(time.RFC3339)}
	}

	return filterConditions
}
