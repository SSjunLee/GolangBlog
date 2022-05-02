package models

import (
	"Myblog/cmd"
	_ "github.com/jinzhu/gorm"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

var (
	Db          *gorm.DB
	migrateList []interface{}
)

func clearDb(name string) {
	Db.Raw(`SELECT concat('DROP TABLE IF EXISTS ', table_name, ';')
FROM information_schema.tables
WHERE table_schema = 'mydb';`)
}

func AutoMigrate(name string) {
	var err error
	clearDb(name)
	err = Db.AutoMigrate(migrateList...)
	if err != nil {
		panic(err)
	}

}

func DbInit() {
	var err error
	c := &cmd.Config
	url := c.DbUrl
	url = "root:123@tcp(localhost:3306)/mydb?charset=utf8&loc=Local&parseTime=True"
	logLevel := logger.Info
	if !c.EnableSqlLog {
		logLevel = logger.Warn
	}
	Db, err = gorm.Open(mysql.Open(url), &gorm.Config{
		SkipDefaultTransaction: true,
		Logger:                 logger.Default.LogMode(logLevel),
	})
	if err != nil {
		log.Println(err)
		panic(err)
	}
}

func PageHelper(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page <= 0 {
			page = 1
		}
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
