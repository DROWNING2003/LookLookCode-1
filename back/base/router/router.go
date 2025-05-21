package router

import (
	"base/controller"
	"base/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Router struct {
	helloController *controller.HelloController
	authController  *controller.AuthController
	apiController   *controller.ApiController
}

func NewRouter() *Router {
	return &Router{
		helloController: controller.NewHelloController(),
	}
}

func (r *Router) SetupRouter() *gin.Engine {
	// 创建Gin路由
	router := gin.Default()

	// 配置CORS中间件
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	config.AllowCredentials = true
	router.Use(cors.New(config))

	// 创建路由分组
	helloGroup := router.Group("/")
	{
		helloGroup.GET("", r.helloController.Index)
		helloGroup.GET("hello", r.helloController.Hello)
	}

	// GitHub OAuth相关路由分组
	apiGroup := router.Group("/api")
	{
		//大模型
		apiGroup.POST("/chat", r.apiController.ChatHandler)
		apiGroup.POST("/chatOllama", r.apiController.Chat)
		apiGroup.POST("/ana", r.apiController.Analisy)
		// 认证相关路由
		authGroup := apiGroup.Group("/auth")
		{
			authGroup.GET("/github", r.authController.GithubAuth)

			authGroup.GET("/github/callback", r.authController.GithubCallback)
		}

		// 需要JWT认证的路由
		githubGroup := apiGroup.Group("/github", middleware.JWTAuthMiddleware())
		{
			githubGroup.GET("/user", r.authController.GithubUserInfo)
		}

		return router
	}
}
