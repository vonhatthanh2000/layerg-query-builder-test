package query

import (
	"database/sql"
	"errors"
	"fmt"
	"net/url"
)

type masterDbConfig struct {
	localDb     *sql.DB
	masterDbUrl string
	useMasterDb bool
}

// NewMasterDbConfig creates a new instance of masterDbConfig with validation
func NewMasterDbConfig(
	localDb *sql.DB,
	masterDbUrl string,
	useMasterDb bool,
) (*masterDbConfig, error) {

	if err := validateDbUrl(masterDbUrl); err != nil {
		return nil, fmt.Errorf("invalid master URL: %w", err)
	}

	return &masterDbConfig{
		localDb:     localDb,
		masterDbUrl: masterDbUrl,
		useMasterDb: useMasterDb,
	}, nil
}

// CreateQueryBuilder creates a new AssetQueryBuilder instance
func (c *masterDbConfig) CreateQueryBuilder() AssetQueryBuilder {
	return NewAssetQueryBuilder(c)
}

// validateDbUrl checks if the provided URL is valid
func validateDbUrl(dbUrl string) error {
	if dbUrl == "" {
		return errors.New("database URL cannot be empty")
	}

	_, err := url.Parse(dbUrl)
	if err != nil {
		return fmt.Errorf("invalid URL format: %w", err)
	}

	return nil
}
