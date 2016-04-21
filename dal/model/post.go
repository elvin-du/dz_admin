package model

import (
	. "dz_admin/dal/context"
	"dz_admin/service/log"
	"fmt"
	"time"
)

type postModel struct{}

var _post = &postModel{}

func PostModel() *postModel {
	return _post
}

func (this *postModel) Add(ctx *Context, title, content string) (*Post, error) {
	id := time.Now().Unix()
	sqlStr := fmt.Sprintf("INSERT dz.posts SET id=%d,title='%s',content='%s'", id, title, content)
	_, err := ctx.Exec(sqlStr)
	if nil != err {
		log.Errorf("sql:%s,err:%s", sqlStr, err)
		return nil, err
	}

	var p Post
	p.Id = uint64(id)
	p.Title = title
	p.Content = content

	return &p, nil
}

func (this *postModel) Query(ctx *Context, id string) (*Post, error) {
	sqlStr := "SELECT * FROM dz.posts"
	rows, err := ctx.Query(sqlStr)
	if nil != err {
		log.Errorf("sql:%s,err:%s", sqlStr, err)
		return nil, err
	}

	var p Post
	for rows.Next() {
		rows.Scan(&p.Id, &p.Title, &p.Content)
	}

	return &p, nil
}
