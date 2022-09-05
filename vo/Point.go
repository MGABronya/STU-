package vo

import "github.com/jinzhu/gorm"

type Point struct {
	gorm.Model
	Name   string `json:"name" gorm:"type:varchar(20);"`
	Remark string `json:"remark" gorm:"type:varchar(100);"`
}
