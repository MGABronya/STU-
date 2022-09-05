package controller

import (
	"STU/common"
	"STU/response"
	"STU/vo"
	"log"

	"github.com/gin-gonic/gin"
)

func ChangeEdgeController(ctx *gin.Context) {

	var requestEdge vo.Edge
	// 数据验证
	if err := ctx.ShouldBind(&requestEdge); err != nil {
		log.Print(err.Error())
		response.Fail(ctx, nil, "数据验证错误")
		return
	}

	// 获取path中的id
	EdgeId := ctx.Params.ByName("id")
	db := common.GetDB()

	var edge vo.Edge
	if db.Where("id = ?", EdgeId).First(&edge).RecordNotFound() {
		response.Fail(ctx, nil, "边不存在")
		return
	}

	// 更新文章
	if err := db.Model(&edge).Update(requestEdge).Error; err != nil {
		response.Fail(ctx, nil, "更新失败")
		return
	}

	response.Success(ctx, gin.H{"edge": edge}, "更新成功")
	update = true
}

func ChangePointController(ctx *gin.Context) {
	var requestPoint vo.Edge
	// 数据验证
	if err := ctx.ShouldBind(&requestPoint); err != nil {
		log.Print(err.Error())
		response.Fail(ctx, nil, "数据验证错误")
		return
	}

	// 获取path中的id
	PointId := ctx.Params.ByName("id")
	db := common.GetDB()

	var point vo.Point
	if db.Where("id = ?", PointId).First(&point).RecordNotFound() {
		response.Fail(ctx, nil, "点不存在")
		return
	}

	// 更新文章
	if err := db.Model(&point).Update(requestPoint).Error; err != nil {
		response.Fail(ctx, nil, "更新失败")
		return
	}

	response.Success(ctx, gin.H{"point": point}, "更新成功")
	update = true
}