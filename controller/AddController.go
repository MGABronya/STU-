// @Title  AddController
// @Description  该文件中含对边、点进行增加的函数
// @Author  MGAronya（张健）
// @Update  MGAronya（张健）  2022-9-16 0:35
package controller

import (
	"STU/common"
	"STU/response"
	"STU/vo"
	"log"

	"github.com/gin-gonic/gin"
)

// @title    AddEdgeController
// @description   向数据库中添加边以及边的信息
// @auth      MGAronya（张健）             2022-9-16 10:10
// @param     ctx        *gin.Context    入参为路由接收的上下文
// @return    void       void            没有返回值
func AddEdgeController(ctx *gin.Context) {
	db := common.GetDB()
	var edge vo.Edge
	// TODO  尝试从上下文中取出edge，如果未能取出，返回数据格式错误的错误信息
	if err := ctx.ShouldBind(&edge); err != nil {
		log.Print(err.Error())
		response.Fail(ctx, nil, "数据格式错误")
		return
	}

	// TODO  从数据库中查看是否存在端点1，不存在，返回端点不存在的错误信息
	if db.Where("id = ?", edge.PointA).First(&vo.Point{}).RecordNotFound() {
		response.Fail(ctx, nil, "端点A不存在")
		return
	}

	// TODO  从数据库中查看是否存在端点2，不存在，返回端点不存在的错误信息
	if db.Where("id = ?", edge.PointB).First(&vo.Point{}).RecordNotFound() {
		response.Fail(ctx, nil, "端点B不存在")
		return
	}

	// TODO 插入数据
	if err := db.Create(&edge).Error; err != nil {
		response.Fail(ctx, nil, "系统错误")
		return
	}

	// TODO 成功
	response.Success(ctx, nil, "创建成功")
	update = true
}

// @title    AddPointController
// @description   向数据库中添加点以及点的信息
// @auth      MGAronya（张健）             2022-9-16 10:12
// @param     ctx        *gin.Context    入参为路由接收的上下文
// @return    void       void            没有返回值
func AddPointController(ctx *gin.Context) {
	db := common.GetDB()
	var point vo.Point

	// TODO  尝试从上下文中取出point，如果未能取出，返回数据格式错误的错误信息
	if err := ctx.ShouldBind(&point); err != nil {
		log.Print(err.Error())
		response.Fail(ctx, nil, "数据格式错误")
		return
	}

	// TODO 插入数据
	if err := db.Create(&point).Error; err != nil {
		response.Fail(ctx, nil, "系统错误")
		return
	}

	// TODO 成功
	response.Success(ctx, nil, "创建成功")
	update = true
}
