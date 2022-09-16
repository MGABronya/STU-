// @Title  DeleteController
// @Description  该文件中含删除边、点的函数
// @Author  MGAronya（张健）
// @Update  MGAronya（张健）  2022-9-16 0:39
package controller

import (
	"STU/common"
	"STU/response"
	"STU/vo"

	"github.com/gin-gonic/gin"
)

// @title    DeleteEdgeController
// @description   删除数据库中对应的边
// @auth      MGAronya（张健）             2022-9-16 10:13
// @param     ctx        *gin.Context    入参为路由接收的上下文
// @return    void       void            没有返回值
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

// @title    DeletePointController
// @description   删除数据库中对应的点
// @auth      MGAronya（张健）             2022-9-16 10:13
// @param     ctx        *gin.Context    入参为路由接收的上下文
// @return    void       void            没有返回值
func DeletePointController(ctx *gin.Context) {
	// TODO 获取path中的id
	PointId := ctx.Params.ByName("id")
	db := common.GetDB()
	var point vo.Point

	// TODO 查看数据库中该点是否存在
	if db.Where("id = ?", PointId).First(&point).RecordNotFound() {
		response.Fail(ctx, nil, "该点不存在")
		return
	}

	// TODO  删除与该点相关的边
	db.Where("pointa = ?", PointId).Delete(vo.Edge{})
	db.Where("pointb = ?", PointId).Delete(vo.Edge{})

	// TODO 删除该点
	db.Delete(&point)

	response.Success(ctx, gin.H{"point": point}, "删除成功")
	update = true
}
