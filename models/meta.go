// @Author ljn 2022/5/2 8:40:00
package models

import (
	"gorm.io/gorm"
	"log"
)

type Meta struct {
	Id          int    `json:"id"gorm:"primary key;auto increment;"`
	FaviconUrl  string `json:"favicon_url"gorm:"type:varchar(200)"`
	Keywords    string `json:"keywords"gorm:"type:varchar(200)"`
	Description string `json:"description"gorm:"type:varchar(200)"`
	Css         string `json:"css"gorm:"type:varchar(200)"`
	Js          string `json:"js"gorm:"type:varchar(200)"`
	GithubUrl   string `json:"github_url"gorm:"type:varchar(200)"`
	WeiboUrl    string `json:"weibo_url"gorm:"type:varchar(200)"`
	Title       string `json:"title"gorm:"type:varchar(200)"`
	LogoUrl     string `json:"logo_url"gorm:"type:varchar(200)"`
	Author      string `json:"author"gorm:"type:varchar(200)"`
	SiteUrl     string `json:"site_url"gorm:"type:varchar(200)"`
	Analytic    string `json:"analytic"gorm:"type:varchar(200)"`
	PageSize    int    `json:"page_size"gorm:"type:varchar(200)"`
}

func MetaEdit(meta *Meta) error {
	return Db.Transaction(func(tx *gorm.DB) error {
		return tx.Table("meta").Where("id=?", meta.Id).
			Updates(meta).Error
	})
}

func MetaAdd(meta *Meta) error {
	return Db.Transaction(func(tx *gorm.DB) error {
		return tx.Create(meta).Error
	})
}

func MetaGet(arg interface{}) *Meta {
	meta := &Meta{}
	var err error
	ss := Db.Session(&gorm.Session{})
	switch v := arg.(type) {
	case int:
		err = ss.First(meta, v).Error
		break
	case string:
		err = ss.Where("name = ?", v).First(meta).Error
		break
	default:
		panic("参数类型错误")
	}
	if err != nil {
		log.Println(err)
		return nil
	}
	return meta
}

var MetaInfo Meta

func (m *Meta) Load() {
	p := MetaGet(1)
	if p == nil {
		_ = MetaAdd(&Meta{})
		return
	}
	*m = *p
}

func init() {
	migrateList = append(migrateList, []interface{}{
		&Meta{},
	}...)
}
