package handlers

import "gorm.io/gorm"

// access from handler directly to database object
// dependency injection
type handler struct {
	DB *gorm.DB
}

// new instance
func New(db *gorm.DB) handler {
	return handler{
		DB: db,
	}
}
