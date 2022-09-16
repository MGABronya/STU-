// @Title  PathController
// @Description  该文件中含返回所请求的最短路径的函数
// @Author  MGAronya（张健）
// @Update  MGAronya（张健）  2022-9-16 0:35
package controller

import (
	"STU/response"
	"STU/util"
	"STU/vo"
	"log"

	"github.com/gin-gonic/gin"
)

var update bool = true

// @title    PathController
// @description   根据上下文中传入的path结构体，求出对应的最短路径
// @auth      MGAronya（张健）             2022-9-16 10:13
// @param     ctx        *gin.Context    入参为路由接收的上下文
// @return    void       void            没有返回值
func PathController(ctx *gin.Context) {
	// TODO 检查是否更新了数据库中的图，如果更新了，更新链式前向星
	if update {
		update = false
		util.InitPath()
	}

	var path vo.Path

	// TODO 从上下文中取出path
	if err := ctx.ShouldBind(&path); err != nil {
		log.Print(err.Error())
		response.Fail(ctx, nil, "数据格式错误")
		return
	}

	// TODO 对必经点个数进行限制
	if len(path.Location) > 6 {
		response.Fail(ctx, nil, "必经点不得大于6个")
	}

	// TODO 调用DFS计算出最短路径
	length, ResPoint, ResEdge := util.DFS(&path)

	// TODO 成功
	response.Success(ctx, gin.H{"Points": ResPoint, "Edges": ResEdge, "length": length}, "成功")
}
