package main

import (
	"Myblog/cmd"
	"Myblog/core"
	"Myblog/models"
	"Myblog/routers"
	"log"
)

func Init() {
	cmd.ConfigInit()
	models.DbInit()
	core.InitFileUploader()
	cmd.InstallCmd("run", "", func(strings []string) {
		models.MetaInfo.Load()
		routers.RouterApp()
	})

	cmd.InstallCmd("migrate", "migrate dbname", func(args []string) {
		if len(args) != 1 {
			panic("[app] migrate dbname...\n")
		}
		models.AutoMigrate(args[0])
		log.Println("数据迁移完成...")
	})
	log.Println("初始化完成.....")

}

func main() {
	Init()
	//routers.RouterApp()
	cmd.Exec()
}
