package config

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

type DbConfig struct {
	Hostname   string
	UserName   string
	Password   string
	DbName     string
	Port       int
	Drivername string
	Sslmodel   bool
	Database   *sql.DB
}

var once sync.Once
var dbConfig *DbConfig

func NewDbConfig(db *DbConfig) *DbConfig {
	once.Do(func() {
		dbInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", db.Hostname, db.Port, db.UserName, db.Password, db.DbName)
		sqlDB, err := sql.Open(db.Drivername, dbInfo)
		if err != nil {
			panic(err)
		}

		dbConfig = db
		dbConfig.Database = sqlDB
	})

	return dbConfig
}

func GetDbConfg() *DbConfig {
	return dbConfig
}
