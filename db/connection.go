package db

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

// Connection is global variable that contains pointer to database connection
var Connection *gorm.DB

// NewConnection opens new connection to database
func NewConnection() *gorm.DB {
	dsn := fmt.Sprintf(
		"postgresql://%v:%v@%v:%v/%v",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect database")
	}

	sqlDB, err := connection.DB()

	if err != nil {
		panic("Failed to get database instance")
	}

	// Configure connection pooling
	sqlDB.SetMaxOpenConns(100)  // Maximum number of open connections
	sqlDB.SetMaxIdleConns(10)   // Maximum number of idle connections
	sqlDB.SetConnMaxLifetime(0) // Maximum amount of time a connection

	return connection
}

// NewMockConnection creates new mock database. Returns pointer to mocked database and mock function for mimic
// this database answers
func NewMockConnection() (*gorm.DB, sqlmock.Sqlmock) {
	dbMock, mock, err := sqlmock.New()
	if err != nil {
		panic("Failed to create mock database: " + err.Error())
	}
	dbConn, openErr := gorm.Open(postgres.New(postgres.Config{
		Conn: dbMock,
	}), &gorm.Config{})
	if openErr != nil {
		panic("Failed to create mock database: " + openErr.Error())
	}
	return dbConn, mock
}
