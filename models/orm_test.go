package models

import (
	"fmt"
	"gorm.io/gorm/clause"
	"log"
	"testing"
	"time"
)

func try(f func() (interface{}, error)) {
	if _, err := f(); err != nil {
		fmt.Print(err)
	}
}

func TestAutoMigrate(t *testing.T) {
	//AutoMigrate()
}

func TestCrate(t *testing.T) {
	DbInit()
	user := &User{Name: "ljn", Created: time.Now()}
	res := Db.Create(user)
	fmt.Println(res)
}

func TestCrate2(t *testing.T) {
	DbInit()
	r := []Role{{Name: "admin"}}
	user := &User{Name: "ljn", Created: time.Now(), Roles: r}
	res := Db.Create(user)
	fmt.Println(res)
}

func TestCrate3(t *testing.T) {
	DbInit()
	r := []Role{{Name: "admin"}, {Name: "boss"}}
	r[0].Permissions = []Permission{
		{Name: "update", Url: "http://"},
	}
	user := &User{Name: "ljn", Created: time.Now(), Roles: r}
	res := Db.Create(user)
	fmt.Println(res)
}

func TestCrateBatch(t *testing.T) {
	DbInit()
	users := []User{{Name: "ljn22", Created: time.Now()}, {Name: "wj", Created: time.Now()}}
	fmt.Println(Db.Create(users))
}

func TestSelect1(t *testing.T) {
	DbInit()
	users := make([]User, 1)
	Db.Where("name like ?", "ljn%").Find(&users).Order("created desc")
	fmt.Println(users)
}

func TestSelect2(t *testing.T) {
	DbInit()
	u := User{}
	Db.First(&u)
	fmt.Println(u)
}

func insertTestCase() {

	ps := []Permission{
		{Id: 1, Name: "mainPage", Url: "/page/mainPage"},
		{Id: 2, Name: "deleteUser", Url: "/control/deleteUser"},
		{Id: 3, Name: "addUser", Url: "/control/addUser"},
	}

	roles := []Role{
		{Id: 1, Name: "admin", Permissions: ps[:2]}, {Id: 2, Name: "root", Permissions: ps}, {Id: 3, Name: "user", Permissions: ps[:1]},
	}

	users := []User{
		{Name: "ljn", Roles: roles[:1], Created: time.Now()}, {Name: "wj", Roles: roles[1:2], Created: time.Now()},
	}
	log.Println(Db.Create(users))
}

func TestSelectPreload(t *testing.T) {
	//AutoMigrate()
	//insertTestCase()
	var users []User
	Db.Preload("Roles.Permissions").Preload("Roles").Find(&users)
	fmt.Println(users)
}

func TestJPreloadAll(t *testing.T) {
	DbInit()
	//AutoMigrate()
	//insertTestCase()
	var users []User
	//预加载全部不会加载嵌套关联
	Db.Preload(clause.Associations).Find(&users)
	fmt.Println(users)
	Db.Preload("Roles.Permissions").Preload(clause.Associations).Find(&users)
	fmt.Println(users)
}

func TestMenu(t *testing.T) {
	DbInit()
	var ids []int32
	_ = Db.Table("user_menu").Select("menu_id").Where("user_id =?", 1).Scan(&ids)
	var menus []Menu
	Db.Find(&menus, ids)
	fmt.Println(menus)
}

func TestPostAdd(t *testing.T) {
	DbInit()
	for i := 0; i < 10; i++ {
		_ = PostAdd(&Post{
			Id:       0,
			CatId:    2,
			Title:    fmt.Sprintf("title %d", i),
			Summary:  fmt.Sprintf("summery %d", i),
			MarkDown: fmt.Sprintf("### 大标题 %d\n\n", i),
			RichText: "<p></p>",
			Created:  time.Now(),
			Updated:  time.Now(),
		})
	}
}

func TestPostEdit(t *testing.T) {
	DbInit()
	_ = PostEdit(&Post{
		Id:       2,
		Title:    "ttt222",
		Summary:  "ttt",
		MarkDown: "ttt",
		RichText: "ttt",
		Created:  time.Now().Add(10 * time.Minute),
		Updated:  time.Now().Add(10 * time.Minute),
	})
}

func TestCateGet(t *testing.T) {
	DbInit()
	r := CateGet(100)
	log.Println(r)
	p := PostGet(100)
	log.Println(p)
	p = GetPostByPath("blog1")
	log.Println(p)
}

func TestPostArchive(t *testing.T) {
	DbInit()
	r, _ := PostArchive()
	log.Println(r)
}

func TestGetAllTagState(t *testing.T) {
	DbInit()
	r := GetAllTagState()
	log.Println(r)
}

func TestTagPostPage(t *testing.T) {
	DbInit()
	r := TagPostPage(1, 1, 5)
	log.Println(r)
}

func TestMeta(t *testing.T) {
	DbInit()
	meta := Meta{
		FaviconUrl:  "/static/favicon.ico",
		Keywords:    "博客",
		Description: "博客描述",
		Css:         "/static/css/app.css",
		Js:          "/static/css/app.js",
		GithubUrl:   "http://www.baidu.com",
		WeiboUrl:    "http://www.baidu.com",
		Title:       "Ljn的博客",
		LogoUrl:     "/static/logo.jpg",
		Author:      "Ljn",
		SiteUrl:     "http://www.abc.com",
		Analytic:    `<script async src="//busuanzi.ibruce.info/busuanzi/2.3/busuanzi.pure.mini.js"></script>`,
		PageSize:    8,
	}
	err := MetaAdd(&meta)
	log.Println(err)
}
