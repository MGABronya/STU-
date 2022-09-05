package controller

import (
	"STU/common"
	"STU/response"
	"STU/vo"
	"log"

	"github.com/gin-gonic/gin"
)

func AddEdgeController(ctx *gin.Context) {
	db := common.GetDB()
	var edge vo.Edge
	if err := ctx.ShouldBind(&edge); err != nil {
		log.Print(err.Error())
		response.Fail(ctx, nil, "数据格式错误")
		return
	}

	if db.Where("id = ?", edge.PointA).First(&vo.Point{}).RecordNotFound() {
		response.Fail(ctx, nil, "端点A不存在")
		return
	}
	if db.Where("id = ?", edge.PointB).First(&vo.Point{}).RecordNotFound() {
		response.Fail(ctx, nil, "端点B不存在")
		return
	}

	// 插入数据
	if err := db.Create(&edge).Error; err != nil {
		response.Fail(ctx, nil, "系统错误")
		return
	}

	// 成功
	response.Success(ctx, nil, "创建成功")
	update = true
}

func AddPointController(ctx *gin.Context) {
	db := common.GetDB()
	var point vo.Point
	if err := ctx.ShouldBind(&point); err != nil {
		log.Print(err.Error())
		response.Fail(ctx, nil, "数据格式错误")
		return
	}

	// 插入数据
	if err := db.Create(&point).Error; err != nil {
		response.Fail(ctx, nil, "系统错误")
		return
	}

	// 成功
	response.Success(ctx, nil, "创建成功")
	update = true
}
