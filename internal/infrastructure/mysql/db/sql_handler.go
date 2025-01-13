package db

import (
	"api-buddy/config"
	"fmt"
	"time"

	"github.com/Fukuemon/go-pkg/query"
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

func ApplyFiltersAndSort(dbQuery *gorm.DB, filters []query.Filter, sort query.SortOption, mappings map[string]query.RelationMapping) *gorm.DB {
	q := query.NewQuery()

	// フィルタを適用
	for _, filter := range filters {
		filter.Apply(q)
	}

	// リレーションに基づくフィルタリング
	for _, mapping := range mappings {
		if value, exists := q.Filters[mapping.FilterField]; exists {
			dbQuery = dbQuery.Joins("JOIN "+mapping.TableName+" ON "+mapping.JoinKey).
				Where(mapping.TableName+"."+mapping.FilterField+" = ?", value)
		}
	}

	// 単純なフィルタリング
	for key, value := range q.Filters {
		if _, isRelationField := mappings[key]; !isRelationField {
			dbQuery = dbQuery.Where(key+" = ?", value)
		}
	}

	// ソートオプションの適用
	if sort.Field != "" {
		if mapping, exists := mappings[sort.Field]; exists {
			dbQuery = dbQuery.Joins("JOIN " + mapping.TableName + " ON " + mapping.JoinKey).
				Order(mapping.TableName + "." + mapping.FilterField + " " + sort.Order)
		} else {
			dbQuery = dbQuery.Order(sort.Field + " " + sort.Order)
		}
	} else {
		dbQuery = dbQuery.Order("created_at DESC") // デフォルトのソート条件
	}

	return dbQuery
}
