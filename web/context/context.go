package context

import (
	"dz_admin/dal/context"

	"github.com/gin-gonic/gin"
)

type Context struct {
	*gin.Context
	DBCtx *context.Context
}

func NewContext(ctx *gin.Context) *Context {
	return &Context{
		Context: ctx,
		DBCtx:   context.NewContext(),
	}
}
