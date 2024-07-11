package routers

import (
	"net/http"

	_ "github.com/1stguardleft/discipline-gbackend/docs"
	"github.com/1stguardleft/discipline-gbackend/routers/api"
	"github.com/gin-gonic/gin"

	"github.com/1stguardleft/discipline-gbackend/pkg/export"
	"github.com/1stguardleft/discipline-gbackend/pkg/qrcode"
	"github.com/1stguardleft/discipline-gbackend/pkg/upload"
)

func InitRouter() *gin.Engine {
	engine := gin.New()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())

	engine.StaticFS("/export", http.Dir(export.GetExcelFullPath()))
	engine.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	engine.StaticFS("/qrcode", http.Dir(qrcode.GetQrCodeFullPath()))

	engine.POST("/auth", api.GetAuth)
	// engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	/*
		apiv1 := engine.Group("/api/v1")
			apiv1.Use(jwt.JWT())
			{
				apiv1.GET("/articles", v1.GetArticles)
				apiv1.GET("/articles/:id", v1.GetArticleß)
			}
	*/
	return engine
}
