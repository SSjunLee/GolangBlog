package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id       int    `gorm:"primary_key;auto increment;"`
	Name     string `gorm:"type:varchar(50);"`
	Password string `gorm:"type:varchar(50);"`
	Created  time.Time
	Roles    []Role `gorm:"many2many:user_role;"`
	Menus    []Menu `gorm:"many2many:user_menu"`
}
type Role struct {
	Id          int          `gorm:"primary_key;auto increment;"`
	Name        string       `gorm:"type:varchar(20);"`
	Permissions []Permission `gorm:"many2many:role_permission;"`
}
type Permission struct {
	Id   int    `gorm:"primary key;auto increment;"`
	Name string `gorm:"type:varchar(20);"`
	Url  string `gorm:"type:varchar(100);"`
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

func FetchUserWithRoleByName(username string) *User {
	user := &User{}
	if Db.Session(&gorm.Session{SkipHooks: true}).Where("name = ?", username).First(user).Error != nil {
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
