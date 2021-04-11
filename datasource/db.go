package datasource

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

var Db *gorm.DB
var err error

func init() {
	uri := os.Getenv("DB_URL")

	config := gorm.Config{
		PrepareStmt:                              true,
		SkipDefaultTransaction:                   true,
		DisableForeignKeyConstraintWhenMigrating: true,
	}

	if uri != "" {
		if Db, err = gorm.Open(mysql.Open(uri), &config); err != nil {
			log.Fatalln(err)
		}
	} else {
		if Db, err = gorm.Open(sqlite.Open("shoppe.db"), &config); err != nil {
			log.Fatalln(err)
		}
	}

	Db.Use(
		dbresolver.Register(dbresolver.Config{ /* xxx */ }).
			SetConnMaxIdleTime(time.Hour).
			SetConnMaxLifetime(24 * time.Hour).
			SetMaxIdleConns(100).
			SetMaxOpenConns(200),
	)
}
