package service

import (
	"dz_admin/service/config"
	"dz_admin/service/db"
	"dz_admin/service/log"
	"dz_admin/service/web"
	_ "dz_admin/web/controller"
)

func Init() {
	config.Init()
	log.Init()
	db.Init()
	web.Start()
}
