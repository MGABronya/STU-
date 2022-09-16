// @Title  Edge
// @Description  边以及其信息
// @Author  MGAronya（张健）
// @Update  MGAronya（张健）  2022-9-16 0:51
package vo

import "github.com/jinzhu/gorm"

// Edge				定义了边的基本信息
type Edge struct {
	gorm.Model         // gorm的模板
	PointA     int     `json:"pointa" gorm:"type:int;"`          // 端点1的编号
	PointB     int     `json:"pointb" gorm:"type:int;"`          // 端点2的编号
	Length     float64 `json:"length" gorm:"type:double;"`       // 边的长度
	Name       string  `json:"name" gorm:"type:varchar(20);"`    // 边的名称
	Remark     string  `json:"remark" gorm:"type:varchar(100);"` // 边的备注
}
