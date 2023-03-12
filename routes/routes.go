package routes

import (
	"ezTikTok/controller"
	"ezTikTok/logger"
	"ezTikTok/settings"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Setup(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	// 注册业务路由
	r.POST("/signup", controller.SignupHandler)

	// 登录业务路由
	r.POST("/login", controller.LoginHandler)

	r.GET("/version", func(c *gin.Context) {
		c.String(http.StatusOK, settings.Conf.Version)
	})
	return r
}
