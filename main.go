package main

import (
	"github.com/1stguardleft/discipline-gbackend/models"
	"github.com/1stguardleft/discipline-gbackend/pkg/gredis"
	"github.com/1stguardleft/discipline-gbackend/pkg/logging"
	"github.com/1stguardleft/discipline-gbackend/pkg/setting"
	"github.com/1stguardleft/discipline-gbackend/routers"
	"github.com/gin-gonic/gin"
)

func init() {
	setting.Setup()
	models.Setup()
	logging.Setup()
	gredis.Setup()
}

func main() {
	gin.SetMode(setting.ServerSetting.RunMode)

	routersInit := routers.InitRouter()
}
