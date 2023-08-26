package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
)

var Db *sql.DB

func InitDB() {
	// Use root:dbpass@tcp(172.17.0.2)/hackernews, if you're using Windows.
	db, err := sql.Open("mysql", "dbuser:dbpass@tcp(db:3306)/hackernews")
	if err != nil {
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
 		log.Panic(err)
	}
	Db = db
}

func CloseDB() error {
	return Db.Close()
}

func Migrate() {
	if err := Db.Ping(); err != nil {
		log.Fatal(err)
	}
	driver, err := mysql.WithInstance(Db, &mysql.Config{})
	if err != nil {
		log.Fatalf("Failed to create driver: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://internal/pkg/db/migrations/mysql",
		"mysql",
		driver,
	)
	if err != nil {
		log.Fatalf("Failed to create instance: %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}

}