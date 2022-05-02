// @Author ljn 2022/4/26 14:35:00
package models

import (
	"gorm.io/gorm"
	"log"
)

type Cate struct {
	Id    int    `json:"id" gorm:"primary key;auto increment;"`
	Name  string `json:"name" gorm:"type:varchar(20)"`
	Intro string `json:"intro" gorm:"type:varchar(200)"`
}

func init() {
	migrateList = append(migrateList, []interface{}{
		&Cate{},
	}...)
}

func GetAllCate() []Cate {
	var cates []Cate
	Db.Session(&gorm.Session{}).Find(&cates)
	return cates
}

func CateGetPage(page, pageSize int) (cates []Cate, err error) {
	r := Db.Session(&gorm.Session{}).Scopes(PageHelper(page, pageSize))
	err = r.Find(&cates).Error
	return
}

func CateCnt() (cnt int64) {
	Db.Session(&gorm.Session{}).Table("cates").Count(&cnt)
	return
}

func CateGet(arg interface{}) *Cate {
	cate := &Cate{}
	var err error
	ss := Db.Session(&gorm.Session{})
	switch v := arg.(type) {
	case int:
		err = ss.First(cate, v).Error
		break
	case string:
		err = ss.Where("name = ?", v).First(cate).Error
		break
	default:
		panic("参数类型错误")
	}
	if err != nil {
		log.Println(err)
		return nil
	}
	return cate
}

func CateEdit(cate *Cate) error {
	return Db.Transaction(func(tx *gorm.DB) error {
		return tx.Table("cates").Where("id=?", cate.Id).
			Updates(cate).Error
	})
}

func CateAdd(cate *Cate) error {
	return Db.Transaction(func(tx *gorm.DB) error {
		return tx.Create(cate).Error
	})
}

func CateDrop(id int) error {
	return Db.Transaction(func(tx *gorm.DB) error {
		return tx.Delete(&Cate{}, id).Error
	})
}
