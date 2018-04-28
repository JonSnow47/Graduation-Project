/*
 * Revision History:
 *     Initial: 2018/04/26        Chen Yanchen
 */

package mysql

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

const DatabaseShop = "shop"

func InitMysql(dbname string) *gorm.DB {
	db, err := gorm.Open("mysql",
		fmt.Sprintf("root:0000000000@/%s?charset=utf8&parseTime=True&loc=Local", dbname))
	if err != nil {
		panic(err)
	}

	db.DB().SetMaxIdleConns(20)
	db.DB().SetMaxOpenConns(100)
	log.Println("Mysql init success!")
	return db
}

func CreateTable(db *gorm.DB, v interface{}) *gorm.DB {
	if !db.HasTable(&v) {
		err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&v).Error
		if err != nil {
			panic(err)
		}
	}
	log.Println("Create table success!")
	return db
}
