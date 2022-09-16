// @Title  Path
// @Description  从前端接收的起点终点以及必经点
// @Author  MGAronya（张健）
// @Update  MGAronya（张健）  2022-9-16 0:51
package vo

import "github.com/jinzhu/gorm"

// Path				定义了接收计算请求的基本信息
type Path struct {
	gorm.Model       //gorm的模板
	Start      int   `json:"start"`    //起点
	Location   []int `json:"location"` //必经点
	End        int   `json:"end"`      //终点
}
