package configs

import (
	"database/sql"
	"log"
	"sync"

	"github.com/yash91989201/go_cart/internal/database"
)

var (
	db     *database.Queries
	dbOnce sync.Once
)

func createDBConnection() *database.Queries {
	DATABASE_URL := GetEnv().DATABASE_URL
	dbOnce.Do(func() {
		var err error
		dbConn, err := sql.Open("postgres", DATABASE_URL)
		if err != nil {
			log.Fatalf("Database connection failed: %v", err)
		}
		if err = dbConn.Ping(); err != nil {
			log.Fatalf("Database connection failed: %v", err)
		}
		log.Print("Database connection successful")
	})

	return db
}

func GetDB() *database.Queries {
	if db == nil {
		db = connectDB()
		return db
	}
	return db
}
