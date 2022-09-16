// @Title  Point
// @Description  点以及其信息
// @Author  MGAronya（张健）
// @Update  MGAronya（张健）  2022-9-16 0:51
package vo

import "github.com/jinzhu/gorm"

// Point			点对象，定义了点的基本信息
type Point struct {
	gorm.Model        //gorm的模板
	Name       string `json:"name" gorm:"type:varchar(20);"`    //点的名称
	Remark     string `json:"remark" gorm:"type:varchar(100);"` //点的备注
}
