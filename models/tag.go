// @Author ljn 2022/4/26 15:31:00
package models

import (
	"gorm.io/gorm"
	"log"
)

type Tag struct {
	Id    int    `json:"id" gorm:"primary_key;auto increment;"`
	Name  string `json:"name"`
	Intro string `json:"intro"`
}

type TagState struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
	Intro string `json:"intro"`
}

func GetAllTagState() []TagState {
	ts := make([]TagState, 0)
	err := Db.Session(&gorm.Session{}).
		Table("tags").
		Select("name,count(post_tag.post_id) as count,intro").
		Joins("join post_tag on post_tag.tag_id = tags.id").
		Group("tags.id").
		Having("count>1").
		Scan(&ts).Error
	if err != nil {
		log.Println(err)
		return nil
	}
	return ts
}

func GetAllTags() []Tag {
	var tgs []Tag
	Db.Session(&gorm.Session{}).Find(&tgs)
	return tgs
}

func GetTagsByPostId(postId int) []Tag {
	var tgs []Tag
	Db.Session(&gorm.Session{}).Table("tags").
		Joins("join post_tag on tags.id = post_tag.tag_id").
		Where("post_tag.post_id").Scan(&tgs)
	return tgs
}

func GetTagCnt() (cnt int64) {
	Db.Session(&gorm.Session{}).Table("tags").Count(&cnt)
	return
}

func TagGetPage(page, pageSize int) (tags []Tag, err error) {
	r := Db.Session(&gorm.Session{}).Scopes(PageHelper(page, pageSize))
	err = r.Find(&tags).Error
	return
}

func TagDrop(id int) error {
	return Db.Transaction(func(tx *gorm.DB) error {
		return tx.Delete(&Tag{}, id).Error
	})
}

func TagEdit(tag *Tag) error {
	return Db.Transaction(func(tx *gorm.DB) error {
		return tx.Model(&Tag{}).Where("id=?", tag.Id).
			Updates(tag).Error
	})
}

func TagAdd(tag *Tag) error {
	return Db.Transaction(func(tx *gorm.DB) error {
		return tx.Create(tag).Error
	})
}

func TagGet(arg interface{}) *Tag {
	tag := &Tag{}
	var err error
	ss := Db.Session(&gorm.Session{})
	switch v := arg.(type) {
	case int:
		err = ss.First(tag, v).Error
		break
	case string:
		err = ss.Where("name = ?", v).First(tag).Error
		break
	default:
		panic("参数类型错误")
	}
	if err != nil {
		log.Println(err)
		return nil
	}
	return tag
}

func TagPostPage(tagId int, pi int, ps int, cols ...string) []Post {
	posts := make([]Post, 0)
	r := Db.Session(&gorm.Session{}).
		Table("posts")

	if len(cols) > 0 {
		r.Select(cols)
	}
	err := r.Joins("join post_tag ON posts.id = post_tag.post_id").
		Where("post_tag.tag_id = ? AND kind = ?", tagId, KindArticle).
		Scopes(PageHelper(pi, ps)).
		Scan(&posts).Error
	if err != nil {
		log.Println(err)
		return nil
	}
	return posts
}

func TagPostCount(tagId int) int {
	cnt := int64(0)
	err := Db.Session(&gorm.Session{}).
		Table("post_tag").
		Count(&cnt).
		Where("tag_id = ?", tagId).
		Group("tag_id").
		Error
	if err != nil {
		log.Println(err)
	}
	return int(cnt)
}

func init() {
	migrateList = append(migrateList, []interface{}{
		&Tag{},
	}...)
}
