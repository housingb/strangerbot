package main

import (
	"time"

	"github.com/didi/gendry/manager"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

type Database struct {
	Host            string
	Port            int
	Username        string
	Password        string
	DBName          string
	Charset         string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime int64
}

func InitDB(cfg Database) error {

	// db conn
	pDB, err := manager.New(
		cfg.DBName,
		cfg.Username,
		cfg.Password,
		cfg.Host,
	).Port(cfg.Port).Open(false)
	if err != nil {
		return err
	}
	pDB.SetMaxOpenConns(cfg.MaxOpenConns)
	pDB.SetMaxIdleConns(cfg.MaxIdleConns)
	pDB.SetConnMaxLifetime(time.Duration(cfg.ConnMaxLifetime) * time.Second)
	if err := pDB.Ping(); err != nil {
		return err
	}

	// gorm db
	pGorm, err := gorm.Open("mysql", pDB)
	if err != nil {
		return err
	}

	DB = pGorm
	DB = DB.LogMode(true)

	return nil
}
