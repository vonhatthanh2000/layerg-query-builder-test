package models

import (
	"database/sql"
	"time"
)

type CollectionType string

const (
	CollectionTypeERC721  CollectionType = "ERC721"
	CollectionTypeERC1155 CollectionType = "ERC1155"
	CollectionTypeERC20   CollectionType = "ERC20"
)

type Collection struct {
	ID                string         `json:"id"`
	ChainID           int32          `json:"chainId"`
	CollectionAddress string         `json:"collectionAddress"`
	Type              CollectionType `json:"type"`
	CreatedAt         time.Time      `json:"createdAt"`
	UpdatedAt         time.Time      `json:"updatedAt"`
	DecimalData       sql.NullInt16  `json:"decimalData"`
	InitialBlock      sql.NullInt64  `json:"initialBlock"`
	LastUpdated       sql.NullTime   `json:"lastUpdated"`
}
