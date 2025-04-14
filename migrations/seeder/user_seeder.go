package seeder

import (
	"encoding/json"
	"io"
	"mtii-backend/entities"
	"mtii-backend/helpers"
	"os"
	"time"

	"gorm.io/gorm"
)

func UserSeeder(db *gorm.DB) error {
	file, err := os.Open("./migrations/json/user.json")
	if err != nil {
		return err
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	var users []entities.User
	if err := json.Unmarshal(byteValue, &users); err != nil {
		return err
	}

	hasTable := db.Migrator().HasTable(&entities.User{})
	if !hasTable {
		if err := db.Migrator().CreateTable(&entities.User{}); err != nil {
			return err
		}
	}

	id := 1
	for _, user := range users {
		user.Id = id
		user.Password, err = helpers.HashPassword(user.Password)
		if err != nil {
			return err
		}
		user.CreatedAt = time.Now()
		user.UpdatedAt = time.Now()

		if err := db.Create(&user).Error; err != nil {
			return err
		}
		id++
	}

	return nil
}
