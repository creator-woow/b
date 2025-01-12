package db

import (
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

// New opens new connection to database and returns pointer to it's instance
func New(username, password, host string, port int, name string) *gorm.DB {
	dsn := fmt.Sprintf(
		"postgresql://%v:%v@%v:%v/%v",
		username, password, host, port, name,
	)
	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Panicln("Failed to connect database")
	}

	sqlDB, err := connection.DB()

	if err != nil {
		log.Panicln("Failed to get database instance")
	}

	sqlDB.SetMaxOpenConns(100)  // Maximum number of open connections
	sqlDB.SetMaxIdleConns(10)   // Maximum number of idle connections
	sqlDB.SetConnMaxLifetime(0) // Maximum amount of time a connection

	return connection
}

// NewMock creates new mock database for tests. Returns pointer to mocked database and mock function for mimic
// this database answers
func NewMock() (*gorm.DB, sqlmock.Sqlmock, *sql.DB) {
	conn, mock, err := sqlmock.New()
	if err != nil {
		log.Panicln("Failed to create mock database: " + err.Error())
	}
	dbConn, openErr := gorm.Open(
		postgres.New(postgres.Config{Conn: conn}), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)},
	)
	if openErr != nil {
		log.Panicln("Failed to create mock database: " + openErr.Error())
	}
	return dbConn, mock, conn
}
