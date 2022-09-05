package vo

import "github.com/jinzhu/gorm"

type Edge struct {
	gorm.Model
	PointA int   `json:"pointa" gorm:"type:int;"`
	PointB int   `json:"pointb" gorm:"type:int;"`
	Length float64 `json:"length" gorm:"type:double;"`
	Name   string `json:"name" gorm:"type:varchar(20);"`
	Remark string `json:"remark" gorm:"type:varchar(100);"`
}
