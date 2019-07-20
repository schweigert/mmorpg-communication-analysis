package repositories

// ClearDB USE ONLY ON TESTS
func ClearDB() {
	migrate()
	db.DropTableIfExists(availableModels()...)
	migrate()
}
