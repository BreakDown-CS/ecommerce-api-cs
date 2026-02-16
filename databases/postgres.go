package databases

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	ConnPostgres *gorm.DB
	oncePostgres sync.Once
)

func PostgresConn() (*gorm.DB, error) {
	var err error
	oncePostgres.Do(func() {
		host := os.Getenv("DB_POSTGRES_URL")
		portStr := os.Getenv("DB_POSTGRES_PORT")
		dbName := os.Getenv("DB_POSTGRES_DB_NAME")
		user := os.Getenv("DB_POSTGRES_USER")
		pass := os.Getenv("DB_POSTGRES_PASSWORD")

		port, convErr := strconv.Atoi(portStr)
		if convErr != nil {
			log.Fatalf("Invalid Postgres port: %v", convErr)
		}

		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Bangkok",
			host, user, pass, dbName, port,
		)

		ConnPostgres, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})

		if err != nil {
			log.Fatalf("Failed to connect to Postgres DB: %v", err)
		}

		log.Println("Connected to Postgres DB successfully")
	})

	return ConnPostgres, err

}

func ClosePostgres() {
	if ConnPostgres != nil {
		sqlDB, err := ConnPostgres.DB()
		if err != nil {
			log.Printf("Error getting database instance: %v", err)
			return
		}
		sqlDB.Close()
		log.Println("Postgres DB connection closed")
	}
}
