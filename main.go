package main

import (
	"fileshare-server/gobalConfig"
	"fileshare-server/model"
	"fileshare-server/router"
	"fileshare-server/util"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron"
	"log"
)

func main() {
	log.Println("正在连接数据库...")
	util.InitDB()
	log.Println("正在检查表结构...")
	model.InitAutoMigrateDB()
	gin.SetMode(gobalConfig.GinMode)
	r := gin.Default()
	if gobalConfig.FrontMode {
		log.Println("已开启前后端整合模式！")
		gobalConfig.UseFrontMode(r)
	}
	router.RegRouter(r)
	c := cron.New()
	c.AddFunc("@every 30s", model.AutoDelFile)
	c.Start()
	log.Println("定时任务启动成功,服务启动成功,当前使用端口：", gobalConfig.ServerPort)
	r.Run(":" + gobalConfig.ServerPort)
}
