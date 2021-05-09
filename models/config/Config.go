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
	TestEnv    bool
}

type JWTConfig struct {
	SecretKey  string
	ExpireTime int //Second
}

var once sync.Once
var dbConfig *DbConfig
var jwtConfig *JWTConfig

func (db *DbConfig) CreateDatabase() error {
	if nil != db.Database {
		_, err := db.Database.Exec(`
			CREATE TABLE IF NOT EXISTS users (
				id UUID NOT NULL,
				username CHAR(50) NOT NULL,
				password CHAR(50) NOT NULL,
				max_todo INTEGER DEFAULT 5 NOT NULL,
				CONSTRAINT users_PK PRIMARY KEY (id)
			);

			INSERT INTO users(id, username, password, max_todo)
				SELECT 'e52e46ad-c655-4011-871c-dce601a986de', 'firstUser', 'example', 5
			WHERE NOT EXISTS (
				SELECT 1 FROM users WHERE username = 'firstUser' AND password = 'example'
			);		
			
			CREATE TABLE IF NOT EXISTS tasks (
				id UUID NOT NULL,
				content TEXT NOT NULL,
				user_id UUID NOT NULL,
				created_date DATE NOT NULL,
				CONSTRAINT tasks_PK PRIMARY KEY (id),
				CONSTRAINT tasks_FK FOREIGN KEY (user_id) REFERENCES users(id)
			);
		`)
		return err
	}
	return nil
}

func NewDbConfig(db *DbConfig) *DbConfig {
	once.Do(func() {
		dbInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", db.Hostname, db.Port, db.UserName, db.Password, db.DbName)
		var sqlDB *sql.DB
		var err error

		if db.TestEnv {
			sqlDB, err = sql.Open("sqlite3", ":memory:")
		} else {
			sqlDB, err = sql.Open(db.Drivername, dbInfo)
		}

		if err != nil {
			panic(err)
		}

		err = sqlDB.Ping()

		if err != nil {
			panic(err)
		}

		dbConfig = db
		dbConfig.Database = sqlDB
		dbConfig.CreateDatabase()
	})

	return dbConfig
}

func GetDbConfg() *DbConfig {
	return dbConfig
}

func NewJWTConfig(jwtInfo *JWTConfig) *JWTConfig {
	jwtConfig = jwtInfo
	return jwtConfig
}

func GetJWTConfig() *JWTConfig {
	return jwtConfig
}
