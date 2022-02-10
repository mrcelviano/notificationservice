package commons

import (
	"fmt"
	gocraft "github.com/gocraft/dbr"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func InitGocraftDBRConnectionPG() *gocraft.Connection {
	connectionString := "host=" + settings.DataBase.Host + " port=" + settings.DataBase.Port + " user=" + settings.DataBase.User + " dbname=" + settings.DataBase.DBName + " sslmode=disable password=" + settings.DataBase.Password
	conn, err := gocraft.Open("postgres", connectionString, nil)
	if err != nil {
		fmt.Printf("Error dbr: %v \n", err)
	}
	conn.SetMaxOpenConns(6)
	conn.SetMaxIdleConns(2)
	return conn
}
