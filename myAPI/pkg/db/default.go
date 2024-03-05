package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"myAPI/pkg/reader"
	"os"
)

func Default() (*gorm.DB, error) {
	var (
		username string
		password string
		host     string
		port     string
		dbName   string
	)

	reader.GetEnv(".env")
	username = os.Getenv("MYSQL_USER")
	password = os.Getenv("MYSQL_PASSWORD")
	host = os.Getenv("MYSQL_HOST")
	dbName = os.Getenv("MYSQL_DATABASE")
	port = os.Getenv("MYSQL_PORT")

	dbConn, err := gorm.Open(
		mysql.Open(fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=UTC",
			username,
			password,
			host,
			port,
			dbName,
		)),
		&gorm.Config{
			CreateBatchSize: 500,
		},
	)
	return dbConn, err
}
