package test

import (
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/stretchr/testify/assert"
)

func NewDbConnection() *gorm.DB {
	err := godotenv.Load("../config.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
		panic(err)
	}

	dbHost := os.Getenv("DATABASE_HOST")
	dbUser := os.Getenv("DATABASE_USER")
	dbPassword := os.Getenv("DATABASE_PASSWORD")
	dbName := os.Getenv("DATABASE_NAME")
	dbPort := os.Getenv("DATABASE_PORT")
	dbSslMode := os.Getenv("DATABASE_SSL_MODE")
	dbTimeZone := os.Getenv("DATABASE_TIME_ZONE")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", dbHost, dbUser, dbPassword, dbName, dbPort, dbSslMode, dbTimeZone)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)
	sqlDB.SetConnMaxIdleTime(5 * time.Minute)

	return db
}

func TestDbConnetion(t *testing.T) {
	db := NewDbConnection()
	assert.Nil(t, db)

}

func TestDbConnetionFailure(t *testing.T) {
	db := NewDbConnection()
	assert.NotNil(t, db)

}
