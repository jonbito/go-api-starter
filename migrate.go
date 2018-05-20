package main

// Migrate automigrates the database
// Put all your models that you want automigrated here
func Migrate() {
	DB.AutoMigrate(Config.MigrationModels...)
}
