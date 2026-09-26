package platform

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Models is every persisted type in the monolith — one database, every context.
var Models []any

// Register adds a context's models to the shared migration set.
func Register(models ...any) { Models = append(Models, models...) }

// Open connects the ONE shared database and migrates every context's tables.
func Open(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if err := db.AutoMigrate(Models...); err != nil {
		return nil, err
	}
	return db, nil
}
