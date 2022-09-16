// @Title  routes
// @Description  程序的路由均集中在这个文件里
// @Author  MGAronya（张健）
// @Update  MGAronya（张健）  2022-9-16 0:50
package main

import (
	"STU/controller"
	"STU/middleware"

	"github.com/gin-gonic/gin"
)

// @title    CollectRoute
// @description   给gin引擎挂上路由监听
// @auth      MGAronya（张健）             2022-9-16 10:52
// @param     r *gin.Engine			gin引擎
// @return    r *gin.Engine			gin引擎
func CollectRoute(r *gin.Engine) *gin.Engine {
	// TODO 添加中间件
	r.Use(middleware.CORSMiddleware(), middleware.RecoveryMiddleware())

	// TODO 添加边的路由
	r.POST("/add/edge", controller.AddEdgeController)

	// TODO 添加点的路由
	r.POST("/add/point", controller.AddPointController)

	// TODO 删除边的路由
	r.DELETE("/delete/edge/:id", controller.DeleteEdgeController)

	// TODO 删除点的路由
	r.DELETE("/delete/point/:id", controller.DeletePointController)

	// TODO 返回所有边的路由
	r.GET("/show/edges", controller.ShowEdgesController)

	// TODO 返回所有点的路由
	r.GET("/show/points", controller.ShowPointsController)

	// TODO 返回指定边的路由
	r.GET("/show/edge/:id", controller.ShowEdgeController)

	// TODO 返回指定点的路由
	r.GET("/show/point/:id", controller.ShowPointController)

	// TODO 改变指定边的路由
	r.POST("/change/edge/:id", controller.ChangeEdgeController)

	// TODO 改变指定点的路由
	r.POST("/change/point/:id", controller.ChangePointController)

	// TODO 查询最短路径的路由
	r.POST("/path", controller.PathController)

	return r
}
