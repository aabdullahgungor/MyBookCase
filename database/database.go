package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	DB_DRIVER = "mysql"
	DB_NAME = "mybookcasedb"
	DB_USER = "root"
	DB_PASSWORD = "mysql123"
)

func GetDB() (*gorm.DB, error) {

	db, err := gorm.Open(DB_DRIVER, DB_USER+":"+DB_PASSWORD+"@/"+DB_NAME+"?charset=utf8&parseTime=True")
	if err != nil {
		return nil, err
	}
	return db, nil
}