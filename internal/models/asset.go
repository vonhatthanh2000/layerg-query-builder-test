package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type AssetResponse interface {
	Pagination[Erc721CollectionAsset] |
		Pagination[Erc1155CollectionAsset] |
		Pagination[Erc20CollectionAsset]
}

type Erc721CollectionAsset struct {
	ID           uuid.UUID      `json:"id"`
	ChainID      int32          `json:"chainId"`
	CollectionID string         `json:"collectionId"`
	TokenID      string         `json:"tokenId"`
	Owner        string         `json:"owner"`
	Attributes   sql.NullString `json:"attributes"`
	CreatedAt    time.Time      `json:"createdAt"`
	UpdatedAt    time.Time      `json:"updatedAt"`
	UpdatedBy    uuid.UUID      `json:"updatedBy"`
	Signature    string         `json:"signature"`
}

type Erc1155CollectionAsset struct {
	ID           uuid.UUID      `json:"id"`
	ChainID      int32          `json:"chainId"`
	CollectionID string         `json:"collectionId"`
	TokenID      string         `json:"tokenId"`
	Owner        string         `json:"owner"`
	Balance      string         `json:"balance"`
	Attributes   sql.NullString `json:"attributes"`
	CreatedAt    time.Time      `json:"createdAt"`
	UpdatedAt    time.Time      `json:"updatedAt"`
	UpdatedBy    uuid.UUID      `json:"updatedBy"`
	Signature    string         `json:"signature"`
}

type Erc20CollectionAsset struct {
	ID           uuid.UUID `json:"id"`
	ChainID      int32     `json:"chainId"`
	CollectionID string    `json:"collectionId"`
	Owner        string    `json:"owner"`
	Balance      string    `json:"balance"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	UpdatedBy    uuid.UUID `json:"updatedBy"`
	Signature    string    `json:"signature"`
}
