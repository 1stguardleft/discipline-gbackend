package main

import (
	"github.com/1stguardleft/discipline_gbackend/pkg/logging"
	"github.com/1stguardleft/discipline_gbackend/pkg/setting"
	"github.com/gin-gonic/gin"
)

func init() {
	setting.Setup()
	logging.Setup()
}

func main() {
	gin.SetMode(setting.ServerSetting.RunMode)

}
