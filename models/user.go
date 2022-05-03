package models

import (
	"gorm.io/gorm"
	"log"
	"time"
)

type User struct {
	Id       int    `json:"id"gorm:"primary_key;auto increment;"`
	Name     string `json:"name"gorm:"type:varchar(50);"`
	Password string `json:"password"gorm:"type:varchar(50);"`
	Created  time.Time
	Roles    []Role `json:"roles"gorm:"many2many:user_role;"`
	Menus    []Menu `json:"menus"gorm:"many2many:user_menu"`
}
type Role struct {
	Id          int          `json:"id"gorm:"primary_key;auto increment;"`
	Name        string       `json:"name"gorm:"type:varchar(20);"`
	Permissions []Permission `json:"permissions"gorm:"many2many:role_permission;"`
}
type Permission struct {
	Id   int    `json:"id"gorm:"primary key;auto increment;"`
	Name string `json:"name"gorm:"type:varchar(20);"`
	Url  string `json:"url"gorm:"type:varchar(100);"`
}

type Menu struct {
	Id   int    `gorm:"primary key;auto increment;"`
	Name string `gorm:"type:varchar(20);"`
	Icon string `gorm:"type:varchar(50);"`
}

func init() {
	migrateList = append(migrateList, []interface{}{
		&Role{}, &Permission{}, &User{}, &Menu{},
	}...)
}

func FetchUser(arg interface{}) *User {
	s := Db.Session(&gorm.Session{SkipHooks: true})
	var err error
	user := &User{}
	switch arg.(type) {
	case string:
		err = s.Where("name=?", arg.(string)).First(user).Error
		break
	case int:
		err = s.First(user, arg.(int)).Error
	}
	if err != nil {
		log.Println(err)
		return nil
	}
	return user
}

func FetchMenuByUser(id int) []Menu {
	var ids []int
	_ = Db.Session(&gorm.Session{SkipHooks: true}).Table("user_menu").Select("menu_id").Where("user_id =?", 1).Scan(&ids)
	var menus []Menu
	Db.Find(&menus, ids)
	return menus
}
