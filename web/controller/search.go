package controller

import (
	. "dz_admin/web/context"

	"github.com/gin-gonic/gin"
)

func init() {
	V1.GET("/search", _search.Index)
	V1.POST("/search", _search.Do)
}

type searchController struct{}

var _search = searchController{}

func (this *searchController) Index(ctx *Context) {
	ctx.HTML(200, "search.html",gin.H{"title":"搜索"})
}

func (this *searchController) Do(ctx *Context) {

	ctx.HTML(200, "search_result.html",gin.H{"title":"搜索结果","result":"我爱你"})
}
