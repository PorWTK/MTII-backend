package migrations

import (
	"mtii-backend/entities"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	tables := []interface{}{
		entities.User{},
		entities.Platform{},
		entities.Status{},
		entities.PaymentMethod{},
		entities.SalePerson{},
		entities.Channel{},
		entities.Bank{},
		entities.Receiver{},
		entities.Income{},
		entities.Detail{},
	}

	for _, table := range tables {
		if !db.Migrator().HasTable(table) {
			if err := db.AutoMigrate(table); err != nil {
				return err
			}
		}
	}

	return nil
}
