package model

import (
	. "dz_admin/dal/context"
	"dz_admin/service/log"
)

type userModel struct{}

var _user = &userModel{}

func UserModel() *userModel {
	return _user
}

func (this *userModel) InfoById(ctx *Context, id string) (*User, error) {
	sqlStr := "SELECT * FROM dz.users"
	rows, err := ctx.Query(sqlStr)
	if nil != err {
		log.Errorf("sql:%s,err:%s", sqlStr, err)
		return nil, err
	}

	var u User
	for rows.Next() {
		rows.Scan(&u.Id, &u.Name)
	}

	return &u, nil
}
