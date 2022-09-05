package controller

import (
	"STU/response"
	"STU/util"
	"STU/vo"
	"log"

	"github.com/gin-gonic/gin"
)

var update bool = true

func PathController(ctx *gin.Context) {
	if update {
		update = false
		util.InitPath()
	}

	var path vo.Path
	if err := ctx.ShouldBind(&path); err != nil {
		log.Print(err.Error())
		response.Fail(ctx, nil, "数据格式错误")
		return
	}

	ResPoint := []int{}
	ResEdge := []int{}

	length := util.DFS(&ResPoint, &ResEdge, &path)

	// 成功
	response.Success(ctx, gin.H{"Points": ResPoint, "Edges": ResEdge, "length": length}, "成功")
}
