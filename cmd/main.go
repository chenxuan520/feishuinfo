package main

import (
	"fmt"

	"github.com/chenxuan520/feishuinfo/internal/config"
	"github.com/chenxuan520/feishuinfo/internal/logger"
	"github.com/chenxuan520/feishuinfo/internal/tools"
	"github.com/chenxuan520/feishuinfo/internal/view"
	"github.com/gin-gonic/gin"
)

func main() {
	err := config.InitConfig()
	if err != nil {
		panic(err)
	}
	err = logger.InitLog("Debug", "console", "[dian-feishu]", "logs", false, "LowercaseLevelEncoder", true)
	if err != nil {
		panic(err)
	}
	err = tools.InitLarkClient(config.GlobalConfig.Feishu.AppID, config.GlobalConfig.Feishu.AppSecret)
	if err != nil {
		panic(err)
	}

	g := gin.Default()
	g.Use(gin.Recovery())
	view.InitGin(g)
	g.Run(fmt.Sprintf("%s:%d", config.GlobalConfig.Server.Host, config.GlobalConfig.Server.Port))
}
