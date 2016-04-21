package controller

import (
	. "dz_admin/web/context"

	"dz_admin/dw"

	"dz_admin/dal/model"
	"dz_admin/service/log"

	"github.com/gin-gonic/gin"
)

func init() {
	V1.GET("/post", _post.Index)
	V1.POST("/post", _post.Do)
}

type postController struct{}

var _post = postController{}

func (this *postController) Index(ctx *Context) {
	ctx.HTML(200, "post.html", gin.H{"title": "发布文章"})
}

func (this *postController) Do(ctx *Context) {
	title := ctx.PostForm("title")
	content := ctx.PostForm("content")
	kind := ctx.PostForm("kind")

	post, err := model.PostModel().Add(ctx.DBCtx, title, content)
	if nil != err {
		log.Errorln(err)
		ctx.String(500, err.Error())
		return
	}

	dw.Add(post.Id, title+content+kind)
	ctx.String(200, "add ok")
	//	resp := dw.Query(key)
	//	ctx.String(200,resp.Tokens[0])
}
