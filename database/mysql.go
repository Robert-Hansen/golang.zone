package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/steffen25/golang.zone/config"
	"log"
)

type MySQLDB struct {
	*sql.DB
}

type GormMySQLDB struct {
	DB *gorm.DB
}

func NewMySQLDB(dbCfg config.MySQLConfig) (*MySQLDB, error) {
	dataSourceName := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s&parseTime=true",
		dbCfg.Username,
		dbCfg.Password,
		dbCfg.Host,
		dbCfg.Port,
		dbCfg.DatabaseName,
		dbCfg.Encoding)
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatalf("Error when connect database, the error is '%v'", err)
	}

	return &MySQLDB{db}, nil
}

func NewGormMySQLDB(dbCfg config.MySQLConfig) (*GormMySQLDB, error) {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		dbCfg.Username,
		dbCfg.Password,
		dbCfg.Host,
		dbCfg.Port,
		dbCfg.DatabaseName,
		dbCfg.Encoding)
	db, err := gorm.Open("mysql", dataSourceName)

	if err != nil {
		log.Fatalf("Error when connect database, the error is '%v'", err)
	}

	return &GormMySQLDB{db}, nil
}
