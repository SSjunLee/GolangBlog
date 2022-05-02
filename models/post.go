package models

import (
	"Myblog/cmd"
	"Myblog/common"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"strings"
	"time"
)

type Post struct {
	Id       int       `gorm:"primary key;auto increment;"json:"id"`
	Kind     int       `json:"kind"`
	Status   int       `gorm:"type:int(2);"json:"status"` //发布状态：1.草稿，2.发布
	CatId    int       `json:"cate_id"`
	Title    string    `gorm:"type:varchar(100);" json:"title"`
	Summary  string    `gorm:"type:varchar(100);" json:"summary"`
	MarkDown string    `gorm:"type:text;" json:"markdown"`
	RichText string    `gorm:"type:mediumtext" json:"richtext"`
	Created  time.Time `json:"created"`
	Updated  time.Time `json:"updated"`
	Path     string    `json:"path" gorm:"type:varchar(100)"`
	Tags     []Tag     `json:"tags"gorm:"many2many:post_tag"`
	Allow    bool      `json:"allow" gorm:"TINYINT(4)"` //允许评论
	//Cate     Cate      `json:"cate" gorm:"foreignKey:CatId"`
}

type Naver struct {
	Prev, Next string
}

func (p *Post) PathExits() bool {
	cnt := int64(0)
	Db.Session(&gorm.Session{}).Where("path = ?", p.Path).Table("posts").Count(&cnt)
	return cnt > 0
}

func (p *Post) HandleRichText() {
	if strings.Contains(p.RichText, "<!--more-->") {
		p.Summary = strings.Split(p.RichText, "<!--more-->")[0]
	}
	p.RichText = common.GetTocHtml(p.RichText)
}

func (p *Post) GetNav() *Naver {
	naver := &Naver{}
	next := &Post{}
	prev := &Post{}
	var err error
	err = Db.Session(&gorm.Session{}).
		Where("kind = ? AND status = 2 AND created>?",
			KindArticle,
			p.Created.Format(common.StdDateTime)).Order("created asc").First(next).Error

	if err == nil && next.Path != "" {
		naver.Next = `<a href = "/post/` + next.Path + `.html" class="next"` + `>` + next.Title + `</a>`
	}
	err = Db.Session(&gorm.Session{}).
		Where("kind = ? AND status = 2 AND created<?",
			KindArticle,
			p.Created.Format(common.StdDateTime)).Order("created desc").First(prev).Error
	if err == nil && prev.Path != "" {
		naver.Prev = `<a href = "/post/` + prev.Path + `.html" class="prev"` + `>` + prev.Title + `</a>`
	}
	return naver
}

func GetPostByPath(path string) *Post {
	return getPostByPathImpl(path, KindArticle)
}

func GetPageByPath(path string) *Post {
	return getPostByPathImpl(path, KindPage)
}

func getPostByPathImpl(path string, kind int) *Post {
	p := &Post{}
	if err := Db.Session(&gorm.Session{}).Preload("Tags").Where("path = ? AND kind = ?", path, kind).First(p).Error; err != nil {
		log.Println(err)
		return nil
	}
	return p
}

const (
	KindArticle = 1 //文章
	KindPage    = 2 //页面
)

func init() {
	migrateList = append(migrateList, []interface{}{
		&Post{},
	}...)
}

func PostGetPageWithTags(kind, catId, page, pageSize int, cols ...string) (posts []Post, err error) {
	r := Db.Session(&gorm.Session{}).Scopes(PageHelper(page, pageSize))
	if kind > 0 {
		r.Where("kind = ?", kind)
	}
	if catId > 0 {
		r.Where("cat_id = ?", catId)
	}
	if len(cols) > 0 {
		r.Select(cols)
	}
	err = r.Preload("Tags").Find(&posts).Error
	return
}

func PostGetPage(kind, catId, page, pageSize int, cols ...string) []Post {
	var posts []Post
	r := Db.Session(&gorm.Session{}).Scopes(PageHelper(page, pageSize))
	if kind > 0 {
		r.Where("kind = ?", kind)
	}
	if catId > 0 {
		r.Where("cat_id = ?", catId)
	}
	if len(cols) > 0 {
		r.Select(cols)
	}
	err := r.Find(&posts).Error
	if err != nil {
		log.Println(err)
		return nil
	}
	return posts
}

func PostCount(kind, catId int) int {
	r := Db.Session(&gorm.Session{}).Table("posts")
	if kind > 0 {
		r.Where("kind = ?", kind)
	}
	if catId > 0 {
		r.Where("cat_id = ?", catId)
	}
	var count int64
	r.Count(&count)
	return int(count)
}

func PostAdd(post *Post) error {
	return Db.Transaction(func(tx *gorm.DB) error {
		return tx.Create(post).Error
	})
}

func PostEdit(post *Post) error {
	return Db.Transaction(func(tx *gorm.DB) error {
		return tx.Table("posts").Where("id=?", post.Id).
			Updates(post).Error
	})
}

func PostDrop(id int) error {
	return Db.Transaction(func(tx *gorm.DB) error {
		return tx.Select("Tags").Delete(&Post{Id: id}).Error
	})
}

func PostGet(id int) *Post {
	p := &Post{}
	err := Db.Session(&gorm.Session{}).Preload("Tags").First(p, id).Error
	if err != nil {
		log.Println(err)
		return nil
	}
	return p
}

type Page struct {
	Pid   int `json:"pid" form:"pid"`
	Psize int `json:"psize" form:"psize"`
	Cond  int `json:"cond" form:"cond"`
}

func (p *Page) Check() error {
	if p.Psize > cmd.Config.PageMax {
		return errors.New("pageSize 过大")
	} else if p.Psize < cmd.Config.PageMin {
		return errors.New("pageSize 过小")
	}
	return nil
}

func BuildPageFromHttpParams(c *gin.Context) Page {
	in := Page{}
	err := c.BindQuery(&in)
	if err != nil {
		panic(err)
	}
	err = in.Check()
	if err != nil {
		panic(err)
	}
	return in
}
