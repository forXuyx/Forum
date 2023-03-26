package routes

import (
	"ezTikTok/controller"
	_ "ezTikTok/docs"
	"ezTikTok/logger"
	"ezTikTok/middlewares"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Setup(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.LoadHTMLFiles("./templates/index.html")
	r.Static("/static", "./static")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	v1 := r.Group("/api/v1")

	// 注册业务路由
	v1.POST("/signup", controller.SignupHandler)
	// 登录业务路由
	v1.POST("/login", controller.LoginHandler)

	// 获取帖子列表
	v1.GET("/posts", controller.GetPostListHandler)
	// 根据时间或分数获取帖子列表
	v1.GET("/posts2", controller.GetPostListHandler2)

	v1.Use(middlewares.JWTAuthMiddleware()) // 应用中间件
	{
		// 获取社区列表
		v1.GET("/community", controller.CommunityHandler)
		// 获取社区详情
		v1.GET("community/:id", controller.CommunityDetailHandler)

		// 创建帖子
		v1.POST("/post", controller.CreatePostHandler)
		// 获取帖子详情
		v1.GET("/post/:id", controller.GetPostDetailHandler)

		// 投票功能
		v1.POST("/vote", controller.PostVoteHandler)
	}

	pprof.Register(r)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})

	return r
}
