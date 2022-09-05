package controller

import (
	"STU/common"
	"STU/response"
	"STU/vo"

	"github.com/gin-gonic/gin"
)

func DeleteEdgeController(ctx *gin.Context) {
	// 获取path中的id
	EdgeId := ctx.Params.ByName("id")
	db := common.GetDB()
	var edge vo.Edge
	if db.Where("id = ?", EdgeId).First(&edge).RecordNotFound() {
		response.Fail(ctx, nil, "该边不存在")
		return
	}

	db.Delete(&edge)

	response.Success(ctx, gin.H{"edge": edge}, "删除成功")
	update = true
}

func DeletePointController(ctx *gin.Context) {
	// 获取path中的id
	PointId := ctx.Params.ByName("id")
	db := common.GetDB()
	var point vo.Point
	if db.Where("id = ?", PointId).First(&point).RecordNotFound() {
		response.Fail(ctx, nil, "该点不存在")
		return
	}

	db.Where("pointa = ?", PointId).Delete(vo.Edge{})
	db.Where("pointb = ?", PointId).Delete(vo.Edge{})

	db.Delete(&point)

	response.Success(ctx, gin.H{"point": point}, "删除成功")
	update = true
}
