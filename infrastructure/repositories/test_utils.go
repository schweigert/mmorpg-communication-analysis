package repositories

// ClearDB USE ONLY ON TESTS
func ClearDB() {
	defer migrate()
	db.DropTableIfExists(availableModels()...)
}
