// @Title  ShowController
// @Description  该文件中含返回点、边信息的函数
// @Author  MGAronya（张健）
// @Update  MGAronya（张健）  2022-9-16 0:40
package controller

import (
	"STU/common"
	"STU/response"
	"STU/vo"

	"github.com/gin-gonic/gin"
)

// @title    ShowPointsController
// @description   给出数据库中所有点的信息
// @auth      MGAronya（张健）             2022-9-16 10:16
// @param     ctx        *gin.Context    入参为路由接收的上下文
// @return    void       void            没有返回值
func ShowPointsController(ctx *gin.Context) {
	var points []vo.Point
	db := common.GetDB()
	db.Find(&points)

	// TODO 记录的总条数
	var total int
	db.Model(vo.Point{}).Count(&total)
	// TODO 返回数据
	response.Success(ctx, gin.H{"points": points, "total": total}, "成功")
}

// @title    ShowEdgesController
// @description   给出数据库中所有边的信息
// @auth      MGAronya（张健）             2022-9-16 10:16
// @param     ctx        *gin.Context    入参为路由接收的上下文
// @return    void       void            没有返回值
func ShowEdgesController(ctx *gin.Context) {
	var edges []vo.Edge
	db := common.GetDB()
	db.Find(&edges)

	// TODO 记录的总条数
	var total int
	db.Model(vo.Edge{}).Count(&total)
	// TODO 返回数据
	response.Success(ctx, gin.H{"edges": edges, "total": total}, "成功")
}

// @title    ShowPointController
// @description   从上下文中读取点的id，并给出该点在数据库中的所有信息
// @auth      MGAronya（张健）             2022-9-16 10:16
// @param     ctx        *gin.Context    入参为路由接收的上下文
// @return    void       void            没有返回值
func ShowPointController(ctx *gin.Context) {
	db := common.GetDB()

	// TODO 获取path中的id
	Id := ctx.Params.ByName("id")

	var point vo.Point

	// TODO 尝试从数据库中查找该点
	if db.Where("id = ?", Id).First(&point).RecordNotFound() {
		response.Fail(ctx, nil, "该点不存在")
		return
	}

	// TODO 返回数据
	response.Success(ctx, gin.H{"point": point}, "成功")
}

// @title    ShowPointController
// @description   从上下文中读取边的id，并给出该点在数据库中的所有信息
// @auth      MGAronya（张健）             2022-9-16 10:16
// @param     ctx        *gin.Context    入参为路由接收的上下文
// @return    void       void            没有返回值
func ShowEdgeController(ctx *gin.Context) {
	db := common.GetDB()

	// TODO 获取path中的id
	Id := ctx.Params.ByName("id")

	var edge vo.Edge

	// TODO 尝试从数据库中取出该边
	if db.Where("id = ?", Id).First(&edge).RecordNotFound() {
		response.Fail(ctx, nil, "该边不存在")
		return
	}

	// TODO 返回数据
	response.Success(ctx, gin.H{"edge": edge}, "成功")
}
