// package config

// import (
// 	"fmt"
// 	"os"

// 	"github.com/joho/godotenv"
// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// )

// func SetUpDatabaseConnection() *gorm.DB {
// 	if os.Getenv("APP_ENV") != "Production" {
// 		err := godotenv.Load(".env")
// 		if err != nil {
// 			fmt.Println(err)
// 			panic(err)
// 		}
// 	}

// 	dbUser := os.Getenv("DB_USER")
// 	dbPass := os.Getenv("DB_PASS")
// 	dbHost := os.Getenv("DB_HOST")
// 	dbName := os.Getenv("DB_NAME")
// 	dbPort := os.Getenv("DB_PORT")

// 	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v TimeZone=Asia/Jakarta", dbHost, dbUser, dbPass, dbName, dbPort)

// 	db, err := gorm.Open(postgres.New(postgres.Config{
// 		DSN:                  dsn,
// 		PreferSimpleProtocol: true,
// 	}), &gorm.Config{})

// 	if err != nil {
// 		fmt.Println(err)
// 		panic(err)
// 	}

// 	fmt.Println("Database Connected")
// 	return db
// }

//	func ClosDatabaseConnection(db *gorm.DB) {
//		dbSQL, err := db.DB()
//		if err != nil {
//			fmt.Println(err)
//			panic(err)
//		}
//		dbSQL.Close()
//	}
package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetUpDatabaseConnection() *gorm.DB {
	// Load local environment variables if not in production
	if os.Getenv("APP_ENV") != "Production" {
		err := godotenv.Load(".env")
		if err != nil {
			fmt.Println("Error loading .env file:", err)
			panic(err)
		}
	}

	// Check for DATABASE_URL (Railway connection string)
	dbURL := os.Getenv("DATABASE_URL")
	var dsn string

	if dbURL != "" {
		// If DATABASE_URL is set, use it directly.
		dsn = dbURL
	} else {
		// Otherwise, build the DSN from individual environment variables.
		dbUser := os.Getenv("DB_USER")
		dbPass := os.Getenv("DB_PASS")
		dbHost := os.Getenv("DB_HOST")
		dbName := os.Getenv("DB_NAME")
		dbPort := os.Getenv("DB_PORT")
		dsn = fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v TimeZone=Asia/Jakarta", dbHost, dbUser, dbPass, dbName, dbPort)
	}

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to connect to database:", err)
		panic(err)
	}

	fmt.Println("Database Connected")
	return db
}

func ClosDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	dbSQL.Close()
}
