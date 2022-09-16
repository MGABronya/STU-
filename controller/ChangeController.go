// @Title  ChangeController
// @Description  该文件中含对边、点信息进行更改的函数
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

// @title    ChangeEdgeController
// @description   更改数据库中存储的边的数据
// @auth      MGAronya（张健）             2022-9-16 10:13
// @param     ctx        *gin.Context    入参为路由接收的上下文
// @return    void       void            没有返回值
func ChangeEdgeController(ctx *gin.Context) {

	var requestEdge vo.Edge
	// TODO 数据验证
	if err := ctx.ShouldBind(&requestEdge); err != nil {
		log.Print(err.Error())
		response.Fail(ctx, nil, "数据验证错误")
		return
	}

	// TODO 获取path中的id
	EdgeId := ctx.Params.ByName("id")
	db := common.GetDB()

	var edge vo.Edge

	// TODO 查看数据库中该边的id是否存在
	if db.Where("id = ?", EdgeId).First(&edge).RecordNotFound() {
		response.Fail(ctx, nil, "边不存在")
		return
	}

	// TODO 尝试更新边，如果更新失败，返回更新失败错误
	if err := db.Model(&edge).Update(requestEdge).Error; err != nil {
		response.Fail(ctx, nil, "更新失败")
		return
	}

	response.Success(ctx, gin.H{"edge": edge}, "更新成功")
	update = true
}

// @title    ChangePointController
// @description   更改数据库中存储的点的数据
// @auth      MGAronya（张健）             2022-9-16 10:13
// @param     ctx        *gin.Context    入参为路由接收的上下文
// @return    void       void            没有返回值
func ChangePointController(ctx *gin.Context) {
	var requestPoint vo.Edge
	// TODO 数据验证
	if err := ctx.ShouldBind(&requestPoint); err != nil {
		log.Print(err.Error())
		response.Fail(ctx, nil, "数据验证错误")
		return
	}

	// TODO 获取path中的id
	PointId := ctx.Params.ByName("id")
	db := common.GetDB()

	var point vo.Point

	// TODO 查看该点ID在数据库中是否存在
	if db.Where("id = ?", PointId).First(&point).RecordNotFound() {
		response.Fail(ctx, nil, "点不存在")
		return
	}

	// TDOD 尝试更新该点
	if err := db.Model(&point).Update(requestPoint).Error; err != nil {
		response.Fail(ctx, nil, "更新失败")
		return
	}

	response.Success(ctx, gin.H{"point": point}, "更新成功")
	update = true
}
