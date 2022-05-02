// @Author ljn 2022/4/28 19:16:00
package view

import (
	"Myblog/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

/*
{
Id:1
SiteUrl:https://blog.zxysilent.com
LogoUrl:/static/logo.png
Author: Title:zxysilent
Keywords:zxysilent,zxyslt,zxy
Description:zxysilent;zxysilent blog;zxyslt;zxyslt blog;
FaviconUrl:/favicon.ico
BeianMiit:蜀ICP备16011344号-2
BeianNism:
Copyright: SiteJs:console.log("https://blog.zxysilent.com")
SiteCss:
PageSize:6
Analytic:<script async src="//busuanzi.ibruce.info/busuanzi/2.3/busuanzi.pure.mini.js"></script>
Comment:{"clientID": "2d028c155cbc14d10f53","clientSecret": "e503c3d371fb046b2ec9ca99253c10b320be0052","repo": "comments","owner": "zxysilent","admin":["zxysilent"],"distractionFreeMode":true,"githubUserName":"zxysilent"}
GithubUrl:https://github.com/zxysilent
WeiboUrl:https://weibo.com/u/7072792903}

*/

func render(c *gin.Context, tmpl string, m gin.H) {
	m["meta"] = models.MetaInfo
	c.HTML(200, tmpl, m)
}

func Index(c *gin.Context) {
	ps := models.MetaInfo.PageSize
	pi, _ := strconv.Atoi(c.Query("page"))
	if pi == 0 {
		pi = 1
	}
	modes := models.PostGetPage(models.KindArticle, -1, pi, ps, "id", "title", "path", "created", "summary")
	if modes == nil {
		modes = make([]models.Post, 0)
	}
	total := models.PostCount(models.KindArticle, -1)
	naver := models.Naver{}
	if pi > 1 {
		naver.Prev = "/?page=" + strconv.Itoa(pi-1)
	}
	if total > (pi-1)*ps {
		naver.Next = "/?page=" + strconv.Itoa(pi+1)
	}
	render(c, "index.html", gin.H{
		"Posts": modes,
		"Naver": naver,
	})

}
