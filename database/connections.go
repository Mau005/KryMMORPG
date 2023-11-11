package database

import (
	"fmt"
	"log"

	"github.com/mau005/ServerMMORPG/configuration"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func connectionSqlite(database configuration.Database, debugMode bool) error {
	var err error

	logDebug := logger.Silent
	if debugMode {
		logDebug = logger.Warn
	}

	DB, err = gorm.Open(sqlite.Open(database.SqlitePath), &gorm.Config{
		Logger: logger.Default.LogMode(logDebug),
	})
	if err != nil {
		return err
	}
	log.Println(configuration.CONNECTION_SQLITE)
	return nil
}

/*Connection Mysql, require dates*/
func connectionMysql(database configuration.Database, debugMode bool) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		database.User, database.Password, database.Host, database.Port, database.Name)

	logDebug := logger.Silent
	if debugMode {
		logDebug = logger.Warn
	}
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logDebug),
	})
	if err != nil {
		return err
	}
	log.Println(configuration.CONNECTION_MYSQL)
	return nil
}
