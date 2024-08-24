package db

import (
	"api-buddy/config"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type SQLHandler struct {
	DB  *gorm.DB
	Err error
}

var dbConn *SQLHandler

func GetDB() *gorm.DB {
	return dbConn.DB
}

func DBOpen(cnf config.DBConfig) {
	dbConn = NewSQLHandler(cnf)
}

func DBClose() {
	sqlDB, _ := dbConn.DB.DB()
	sqlDB.Close()
}

func NewSQLHandler(cnf config.DBConfig) *SQLHandler {
	user := cnf.User
	password := cnf.Password
	host := cnf.Host
	port := cnf.Port
	dbName := cnf.Name
	fmt.Printf("user: %s, password: %s, host: %s, port: %s, dbname: %s\n", user, password, host, port, dbName)
	dsn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbName + "?parseTime=true&loc=Asia%2FTokyo"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// DB接続数の設定
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(100)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(100 * time.Second)

	// DBインスタンス生成
	sqlHandler := new(SQLHandler)
	db.Logger.LogMode(4)
	sqlHandler.DB = db
	return sqlHandler
}
