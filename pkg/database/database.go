package database

import (
	"fmt"
	"log"

	"github.com/hosting-engine/envoy-proxy-manager-api/pkg/config"

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

// New ...
func New() *gorm.DB {
	var db *gorm.DB
	var err error

	if config.S.DatabaseType == TypeSQLite {
		db, err = gorm.Open(sqlite.Open(config.S.DatabaseDSN), &gorm.Config{})
		if err != nil {
			log.Fatalf("%v", fmt.Errorf("not possible to connect to SQLite: %w", err))
		}
	} else if config.S.DatabaseType == TypeMySQL {
		db, err = gorm.Open(mysql.Open(config.S.DatabaseDSN), &gorm.Config{})
		if err != nil {
			log.Fatalf("%v", fmt.Errorf("not possible to connect to MySQL: %w", err))
		}
	} else if config.S.DatabaseType == TypePostgres {
		db, err = gorm.Open(postgres.Open(config.S.DatabaseDSN), &gorm.Config{})
		if err != nil {
			log.Fatalf("%v", fmt.Errorf("not possible to connect to Postgres: %w", err))
		}
	} else if config.S.DatabaseType == TypeSQLServer {
		db, err = gorm.Open(sqlserver.Open(config.S.DatabaseDSN), &gorm.Config{})
		if err != nil {
			log.Fatalf("%v", fmt.Errorf("not possible to connect to SQLServer: %w", err))
		}
	} else {
		log.Fatalf("DATABASE_TYPE '%s' not supported. Please configure env DATABASE_TYPE with one of possible types: '%s', '%s', '%s', '%s'",
			config.S.DatabaseType, TypeSQLite, TypeMySQL, TypePostgres, TypeSQLServer)
	}

	return db
}
