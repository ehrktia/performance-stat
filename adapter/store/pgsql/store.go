package pgsql

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	_ "github.com/lib/pq"
)

const (
	USERNAME = "DB_USERNAME"
	PASSWORD = "DB_PASSWORD"
	HOST     = "DB_HOST"
	DBNAME   = "DB_NAME"
)

var (
	DEFAULTUSER     = "postgres"
	DEFAULTPASSWORD = "postgres"
	DEFAULTHOST     = "localhost"
	DEFAULTDB       = "dev"
)

type pgStore struct {
	connection *sql.DB
}

func isEmpty(in string) bool {
	return strings.EqualFold(in, "")

}

func getConnFromEnv() string {
	u := os.Getenv(USERNAME)
	if isEmpty(u) {
		u = DEFAULTUSER
	}
	pwd := os.Getenv(PASSWORD)
	if isEmpty(pwd) {
		pwd = DEFAULTPASSWORD
	}
	h := os.Getenv(HOST)
	if isEmpty(h) {
		h = DEFAULTHOST

	}
	d := os.Getenv(DBNAME)
	if isEmpty(d) {
		d = DEFAULTDB
	}
	connectionString := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", u, pwd, h, d)
	return connectionString

}

func New() (*pgStore, error) {
	s := &pgStore{}
	connStr := getConnFromEnv()
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return s, err
	}
	if err := db.Ping(); err != nil {
		return s, err
	}
	s.connection = db
	return s, nil
}
