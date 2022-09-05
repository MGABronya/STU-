package controller

import (
	"STU/common"
	"STU/response"
	"STU/vo"

	"github.com/gin-gonic/gin"
)

func ShowPointsController(ctx *gin.Context) {	
	var points []vo.Point
	db := common.GetDB()
	db.Find(&points)
	
	// 记录的总条数
	var total int
	db.Model(vo.Point{}).Count(&total)
	// 返回数据
	response.Success(ctx, gin.H{"points": points, "total": total}, "成功")
}

func ShowEdgesController(ctx *gin.Context) {	
	var edges []vo.Edge
	db := common.GetDB()
	db.Find(&edges)
	
	// 记录的总条数
	var total int
	db.Model(vo.Edge{}).Count(&total)
	// 返回数据
	response.Success(ctx, gin.H{"edges": edges, "total": total}, "成功")
}

func ShowPointController(ctx *gin.Context) {	
	db := common.GetDB()
	// 获取path中的id
	Id := ctx.Params.ByName("id")
	var point vo.Point
	if db.Where("id = ?", Id).First(&point).RecordNotFound() {
		response.Fail(ctx, nil, "该点不存在")
		return
	}
	// 返回数据
	response.Success(ctx, gin.H{"point": point}, "成功")
}

func ShowEdgeController(ctx *gin.Context) {	
	db := common.GetDB()
	// 获取path中的id
	Id := ctx.Params.ByName("id")
	var edge vo.Edge
	if db.Where("id = ?", Id).First(&edge).RecordNotFound() {
		response.Fail(ctx, nil, "该边不存在")
		return
	}
	// 返回数据
	response.Success(ctx, gin.H{"edge": edge}, "成功")
}