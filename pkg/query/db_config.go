package query

import (
	"errors"
	"fmt"
	"net/url"
)

type masterDbConfig struct {
	localUrl    string
	masterDbUrl string
	useMasterDb bool
}

// NewMasterDbConfig creates a new instance of masterDbConfig with validation
func NewMasterDbConfig(
	localUrl string,
	masterDbUrl string,
	useMasterDb bool,
) (*masterDbConfig, error) {
	// Validate URLs
	if err := validateDbUrl(localUrl); err != nil {
		return nil, fmt.Errorf("invalid local URL: %w", err)
	}
	if err := validateDbUrl(masterDbUrl); err != nil {
		return nil, fmt.Errorf("invalid master URL: %w", err)
	}

	return &masterDbConfig{
		localUrl:    localUrl,
		masterDbUrl: masterDbUrl,
		useMasterDb: useMasterDb,
	}, nil
}

// GetDbUrl returns the appropriate database URL based on configuration
func (c *masterDbConfig) GetDbUrl() string {
	if c.useMasterDb {
		return c.masterDbUrl
	}
	return c.localUrl
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
