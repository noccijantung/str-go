package config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

type PushConfig struct {
	Path         string
	ApiToken     string
	ApiUrl       string
	PushOnlyFile bool
}

var StrConfig *PushConfig
var Scon *sql.DB
var isDb = false

func SetConfig(path string, token string, url string, onlyfile bool, pushto string) {
	StrConfig = &PushConfig{Path: path, ApiToken: token, ApiUrl: url, PushOnlyFile: onlyfile}
	if pushto == "db" {
		Scon = Makeconn()
		isDb = true
	}
}

func Makeconn() *sql.DB {
	conn, err := sql.Open("mysql", "apppooluser:1qazxc3e2w@tcp(127.0.0.1:3306)/pool_bridge")
	if err != nil {
		log.Printf("Sql Database error: %s", err)
		os.Exit(1)
	}
	isDb = true
	return conn
}

func Closedb() bool {
	if isDb {
		db := Scon
		if db == nil {
			return true
		} else {
			defer func(db *sql.DB) {
				_ = db.Close()
			}(db)
			return true
		}
	}
	return true
}
