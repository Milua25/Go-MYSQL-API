package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	database, err := gorm.Open("mysql", "ayomide:voshu2-hytsit-zypMaf@tcp(test-database.cmkkt3mlvh2g.us-west-2.rds.amazonaws.com:3306)/test_database?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	//defer db.Close()
	db = database
}

func GetDB() *gorm.DB {
	return db
}
