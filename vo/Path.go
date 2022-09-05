package vo

import "github.com/jinzhu/gorm"

type Path struct {
	gorm.Model
	Start    int   `json:"start"`
	Location []int `json:"location"`
	End      int   `json:"end"`
}
