package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"strings"
)

func GetConnection() *gorm.DB {
	var db *gorm.DB
	var err error

	dsn := strings.Join([]string{
		os.Getenv("MYSQL_USER") + ":",
		os.Getenv("MYSQL_PASSWORD"),
		"@tcp(" + os.Getenv("MYSQL_HOST") + ")/",
		os.Getenv("MYSQL_DB_NAME"),
		"?parseTime=true",
	}, "")

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	Migrate(db)

	if err != nil {
		panic(err)
	}

	return db
}
