package repository

// MockRepository is a blank repository that returns nothing for testing
type MockRepository struct{}

// Create creates a new database entry of data
func (r *MockRepository) Create(data interface{}) error {
	return nil
}

// Find retrieves a database entry and places it in data
func (r *MockRepository) Find(out interface{}, where ...interface{}) error {
	return nil
}

// First returns the first entry in the database query
func (r *MockRepository) First(out interface{}, where ...interface{}) error {
	return nil
}
