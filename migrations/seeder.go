package migrations

import (
	"mtii-backend/migrations/seeder"

	"gorm.io/gorm"
)

func Seeder(db *gorm.DB) error {
	if err := seeder.UserSeeder(db); err != nil {
		return err
	}
	return nil
}
