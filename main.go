package main

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
	common "github.com/u2u-labs/go-layerg-common/masterdb"
)

func main() {

	localDbUrl := "postgres://postgres:password@localhost:5432/layerg-masterdb?sslmode=disable"

	localDb, err := sql.Open("postgres", localDbUrl)
	if err != nil {
		log.Fatalf("Failed to connect to local database: %v", err)
	}

	// Initialize database configuration
	masterDbClient, err := common.NewMasterDbConfig(
		localDb,
		"https://0c4d-14-241-247-139.ngrok-free.app",
		true,
	)
	if err != nil {
		log.Fatalf("Failed to initialize database configuration: %v", err)
	}

	// Get the database URL

	// Create a new asset query builder using the configuration
	asset := masterDbClient.CreateQueryBuilder().
		WithChainId(1).
		WithCollectionId("1:0x0091BD12166d29539Db6bb37FB79670779aBf266").
		// WithTokenIds([]string{"1", "2", "3"}).
		// WithOwner("0x821dAb5C6fffD8183d4E3e4A5C1725c847c36789").
		WithCreatedAtFrom(time.Date(2025, 1, 5, 0, 0, 0, 0, time.UTC)).
		WithCreatedAtTo(time.Date(2025, 1, 10, 0, 0, 0, 0, time.UTC)).
		WithPage(1).
		WithLimit(10).
		Build()

	_, err = asset.GetPaginatedAsset()
	if err != nil {
		log.Fatalf("Error getting paginated assets: %v", err)
	}

}
