package main

import (
	"STU/controller"
	"STU/middleware"

	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.Use(middleware.CORSMiddleware(), middleware.RecoveryMiddleware())

	r.POST("/add/edge", controller.AddEdgeController)
	r.POST("/add/point", controller.AddPointController)

	r.DELETE("/delete/edge/:id", controller.DeleteEdgeController)
	r.DELETE("/delete/point/:id", controller.DeletePointController)

	r.GET("/show/edges", controller.ShowEdgesController)
	r.GET("/show/points", controller.ShowPointsController)

	r.GET("/show/edge/:id", controller.ShowEdgeController)
	r.GET("/show/point/:id", controller.ShowPointController)

	r.POST("/change/edge/:id", controller.ChangeEdgeController)
	r.POST("/change/point/:id", controller.ChangePointController)

	r.POST("/path", controller.PathController)

	return r
}
