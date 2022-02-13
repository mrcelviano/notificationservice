package postgres

import (
	"fmt"
	goCraft "github.com/gocraft/dbr"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	maxOpenConnections = 6
	maxIdleConnections = 2

	sslMode = "disable"
)

func NewGoCraftDBConnectionPG(host, port, user, password, dbName string) (*goCraft.Connection, error) {
	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s password=%s",
		host, port, user, dbName, sslMode, password)

	conn, err := goCraft.Open("postgres", connectionString, nil)
	if err != nil {
		return nil, err
	}

	conn.SetMaxOpenConns(maxOpenConnections)
	conn.SetMaxIdleConns(maxIdleConnections)

	return conn, nil
}
