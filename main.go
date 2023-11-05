package main

import (
	"github.com/futuregadgetlabx/bit-particle-cannon/config"
	bcron "github.com/futuregadgetlabx/bit-particle-cannon/cron"
	"github.com/futuregadgetlabx/bit-particle-cannon/registry"
	"github.com/futuregadgetlabx/bit-particle-cannon/route"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron"
)

func main() {
	config.Load()

	r := gin.Default()
	route.Init(r)
	registry.Load()
	c := cron.New()
	// 添加定时任务
	err := c.AddFunc("0 0 10 * * ? ", bcron.Notification)
	if err != nil {
		panic(err)
	}

	// 启动定时器
	c.Start()
	err = r.Run(":8081")
	if err != nil {
		panic(err)
	}
}
