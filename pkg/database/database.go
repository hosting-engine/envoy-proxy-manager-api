package database

import (
	"fmt"
	"github.com/hosting-engine/envoy-proxy-manager-api/pkg/config"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

/* Basic supported databases */
const (
	TypeSQLite    = "sqlite"
	TypeMySQL     = "mysql"
	TypePostgres  = "postgres"
	TypeSQLServer = "sqlserver"
)

// NewDatabaseConnection ...
func NewDatabaseConnection(conf *config.Settings) *gorm.DB {
	var db *gorm.DB
	var err error

	if conf.DatabaseType == TypeSQLite {
		db, err = gorm.Open(sqlite.Open(conf.DatabaseDSN), &gorm.Config{})
		if err != nil {
			log.Fatalf("%v", fmt.Errorf("not possible to connect to SQLite: %w", err))
		}
	} else if conf.DatabaseType == TypeMySQL {
		db, err = gorm.Open(mysql.Open(conf.DatabaseDSN), &gorm.Config{})
		if err != nil {
			log.Fatalf("%v", fmt.Errorf("not possible to connect to MySQL: %w", err))
		}
	} else if conf.DatabaseType == TypePostgres {
		db, err = gorm.Open(postgres.Open(conf.DatabaseDSN), &gorm.Config{})
		if err != nil {
			log.Fatalf("%v", fmt.Errorf("not possible to connect to Postgres: %w", err))
		}
	} else if conf.DatabaseType == TypeSQLServer {
		db, err = gorm.Open(sqlserver.Open(conf.DatabaseDSN), &gorm.Config{})
		if err != nil {
			log.Fatalf("%v", fmt.Errorf("not possible to connect to SQLServer: %w", err))
		}
	} else {
		log.Fatalf("DATABASE_TYPE '%s' not supported. Please configure env DATABASE_TYPE with one of possible types: '%s', '%s', '%s', '%s'",
			conf.DatabaseType, TypeSQLite, TypeMySQL, TypePostgres, TypeSQLServer)
	}

	return db
}
