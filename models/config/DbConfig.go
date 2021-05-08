package config

import "database/sql"

type DbConfig struct {
	UserName   string
	Password   string
	ServerName string
	Port       int
	Database   *sql.DB
}
