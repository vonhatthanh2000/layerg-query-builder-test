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

// WithChainId implements AssetQueryBuilder.
func (b *assetQueryBuilderParam) getCollectionType() (*HttpClient, masterDbCommon.CollectionType) {
	client := NewHttpClient(b.config.GetDbUrl())

	fmt.Println("b.config.GetDbUrl()", b.config.GetDbUrl())

	var response response.HTTPResponse[masterDbCommon.CollectionResponse]
	path := fmt.Sprintf("/chain/%d/collection/%s", b.chainId, *b.collectionId)
	err := client.DoRequest(context.Background(), "GET", path, nil, &response)
	if err != nil {
		fmt.Println("err", err)
		return nil, masterDbCommon.CollectionType("")
	}

	return client, response.Data.Type
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

// WithOffset implements AssetQueryBuilder.
func (b *assetQueryBuilderParam) WithOffset(offset int) AssetQueryBuilder {
	b.offset = &offset
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
	WithOffset(offset int) AssetQueryBuilder
	Build() AssetQueryFunction
}

func (b *assetQueryBuilderParam) Build() AssetQueryFunction {
	return b
}

func (b *assetQueryBuilderParam) GetAssetQueryBuilder() (*assetQueryBuilderParam, error) {
	fmt.Println("GetAssetQueryBuilder")

	if b.chainId == 0 {
		return nil, errors.New("chainId is required")
	}
	return b, nil
}

func (b *assetQueryBuilderParam) GetPaginatedAsset() (any, error) {
	httpClient, collectionType := b.getCollectionType()
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

func NewAssetQueryBuilder(config *masterDbConfig) AssetQueryBuilder {
	return &assetQueryBuilderParam{config: config}
}
