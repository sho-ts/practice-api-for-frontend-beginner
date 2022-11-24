package db

import (
	"gorm.io/gorm"
	"note-app/infrastructure/db/schema"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(
		&schema.Note{},
	)
}
