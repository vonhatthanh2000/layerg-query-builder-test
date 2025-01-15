package interfaces

import (
	"time"
)

type AssetQueryBuilder interface {
	WithChainId(chainId int32) AssetQueryBuilder
	WithCollectionId(collectionId string) AssetQueryBuilder
	WithTokenId(tokenId string) AssetQueryBuilder
	WithOwner(owner string) AssetQueryBuilder
	WithCreatedAtFrom(createdAtFrom time.Time) AssetQueryBuilder
	WithCreatedAtTo(createdAtTo time.Time) AssetQueryBuilder
	WithPage(page int) AssetQueryBuilder
	WithLimit(limit int) AssetQueryBuilder
	WithOffset(offset int) AssetQueryBuilder
	Build() AssetQueryFunction
}

type AssetQueryFunction interface {
	GetAssetQueryBuilder() error
	GetPaginatedAsset() (any, error)
}
