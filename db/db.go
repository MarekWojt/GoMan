package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/MarekWojt/GoMan/config"
)

// DB is the gorm database
var DB *gorm.DB = mustOpenDb()

func mustOpenDb() *gorm.DB {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       config.DbDsn, // data source name
		DefaultStringSize:         255,          // default size for string fields
		DisableDatetimePrecision:  true,         // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,         // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,         // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false,        // auto configure based on currently MySQL version
	}), &gorm.Config{})
	if err != nil {
		panic("Database connection failed: " + err.Error())
	}

	return db
}
