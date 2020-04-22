package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go_blog/pkg/setting"
	"log"
)

var db *gorm.DB

type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
}

func init() {
	var (
		err                                               error
		dbType, dbName, user, password, host, tablePrefix string
	)

	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		log.Fatal(2, "Fail to get section 'database':%v", err)
	}

	dbType = sec.Key("DB_TYPE").MustString("mysql")
	dbName = sec.Key("NAME").MustString("go_blog")
	user = sec.Key("USER").MustString("root")
	password = sec.Key("PASSWORD").MustString("")
	host = sec.Key("HOST").MustString("localhost:3306")
	tablePrefix = sec.Key("TABLE_PREFIX").MustString("blog_")

	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))

	if err != nil {
		log.Println(err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}

	db.SingularTable(true)
	db.LogMode(true)
	db.DB().SetConnMaxLifetime(10)
	db.DB().SetConnMaxLifetime(100)
}

func CloseDB() {
	defer db.Close()
}
