package controller

import (
	"dz_admin/dal/model"
	. "dz_admin/web/context"
)

func init() {
	V1.GET("/user", _user.Info)
}

type userController struct{}

var _user = userController{}

func (this *userController) Info(ctx *Context) {
	u, err := model.UserModel().InfoById(ctx.DBCtx, "1")
	if nil != err {
		ctx.String(500, err.Error())
		return
	}

	ctx.String(200, u.Name)
}
