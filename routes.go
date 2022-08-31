package main

import (
	"Essential/controller"
	"Essential/middleware"

	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.Use(middleware.CORSMiddleware(), middleware.RecoveryMiddleware())
	r.POST("/regist", controller.Register)
	r.POST("/login", controller.Login)

	r.GET("/personal", middleware.AuthMiddleware(), controller.PersonalPage)
	r.PUT("/personal", middleware.AuthMiddleware(), controller.PersonalUpdate)
	r.POST("/personal", middleware.AuthMiddleware(), controller.PersonalIcon)
	r.GET("/personal/:id", middleware.AuthMiddleware(), controller.PersonalShow)

	articleRoutes := r.Group("/article")
	articleRoutes.Use(middleware.AuthMiddleware())
	articleController := controller.NewArticleController()
	articleRoutes.POST("", articleController.Create)
	articleRoutes.PUT("/:id", articleController.Update) //替换
	articleRoutes.GET("/:id", articleController.Show)
	articleRoutes.DELETE("/:id", articleController.Delete)
	articleRoutes.POST("/pagelist", articleController.PageList)

	postRoutes := r.Group("/post")
	postRoutes.Use(middleware.AuthMiddleware())
	postController := controller.NewPostController()
	postRoutes.POST("", postController.Create)
	postRoutes.PUT("/:id", postController.Update) //替换
	postRoutes.GET("/:id", postController.Show)
	postRoutes.DELETE("/:id", postController.Delete)
	postRoutes.POST("/pagelist", postController.PageList)

	postRoutes.POST("/:id", postController.ThreadCreate)

	return r
}
