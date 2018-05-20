package main

import (
	"time"
	// you will need to update this if you aren't using postgres
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"go-api-starter/models"
)

// Config is where we can globally change any configurate data
var Config = ConfigContainer{
	// DatabaseDialect is the GORM dialect. see gorm.io for more information.  If this changes, you will need to import it above!
	DatabaseDialect: "postgres",
	// GetEnv tries to get an environment variable called DATABASE_URL. If it's not found, we use the second parameter.  see: getEnv.go
	DatabaseURL: GetEnv("DATABASE_URL", "postgres://postgres:admin@localhost:5432/test?sslmode=disable"),
	// MigrationModels automigrates your database for you based on the models provided.  Add more models to the array if you wish.
	MigrationModels: []interface{}{
		&models.User{},
	},

	/// Rate Limit

	// The time duration to limit requests by
	RateLimiterPeriod: 1 * time.Hour,
	// The number of requests allowed in RateLimiterPeriod
	RateLimiterLimit: 1000,
}

////////////////////////////////////////////////
/// NO NEED TO EDIT BELOW THIS LINE
////////////////////////////////////////////////

// ConfigContainer is simply a container for the config data.  No need to change this unless you want to add more config attributes
type ConfigContainer struct {
	DatabaseURL       string
	DatabaseDialect   string
	MigrationModels   []interface{}
	RateLimiterPeriod time.Duration
	RateLimiterLimit  int64
}