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

	// for feishu check
	// g.POST("/api/event", func(c *gin.Context) {
	// 	type ChallengeReq struct {
	// 		Challenge string `json:"challenge"`
	// 	}
	// 	req := ChallengeReq{}
	// 	err := c.BindJSON(&req)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		c.JSON(400, gin.H{
	// 			"code": 1,
	// 			"msg":  err.Error(),
	// 		})
	// 		return
	// 	}
	// 	c.JSON(200, gin.H{
	// 		"challenge": req.Challenge,
	// 	})
	// })
	view.InitGin(g)
	g.Run(fmt.Sprintf("%s:%d", config.GlobalConfig.Server.Host, config.GlobalConfig.Server.Port))
}
