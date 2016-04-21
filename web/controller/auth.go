package controller

import (
	. "dz_admin/web/context"
)

func init() {
	V1.GET("/auth", _auth.Auth)
}

type authController struct{}

var _auth = authController{}

func (this *authController) Auth(ctx *Context) {
	ctx.String(200, "love %s", "z")
}
