package vo

import "github.com/jinzhu/gorm"

type Path struct {
	gorm.Model
	Start int   `json:"start"`
	End   []int `json:"end"`
}
