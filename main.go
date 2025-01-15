package main

import (
	"asset-query/pkg/query"
	"fmt"
	"log"
	"time"
)

func main() {
	// Initialize database configuration
	masterDbClient, err := query.NewMasterDbConfig(
		"postgresql://localhost:5432/localdb",
		"https://44b3-14-241-247-139.ngrok-free.app",
		true,
	)
	if err != nil {
		log.Fatalf("Failed to initialize database configuration: %v", err)
	}

	// Get the database URL
	dbUrl := masterDbClient.GetDbUrl()
	fmt.Printf("Using database: %s\n\n", dbUrl)

	// Create a new asset query builder using the configuration
	asset := masterDbClient.CreateQueryBuilder().
		WithChainId(1).
		WithCollectionId("1:0x0091BD12166d29539Db6bb37FB79670779aBf266").
		WithTokenIds([]string{"1", "2", "3"}).
		WithOwner("0x821dAb5C6fffD8183d4E3e4A5C1725c847c36789").
		WithCreatedAtFrom(time.Date(2025, 1, 5, 0, 0, 0, 0, time.UTC)).
		WithCreatedAtTo(time.Date(2025, 1, 10, 0, 0, 0, 0, time.UTC)).
		WithPage(1).
		WithLimit(10).
		WithOffset(0).
		Build()

	_, err = asset.GetPaginatedAsset()
	if err != nil {
		log.Fatalf("Error getting paginated assets: %v", err)
	}

}
