package model

import "gorm.io/gorm"

type Crunchy struct {
	Name string `json:"name"`
	gorm.Model
}
