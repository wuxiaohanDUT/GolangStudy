package model

import "github.com/jinzhu/gorm"

type Person struct {
	gorm.Model
	Name string `gorm:"column:person_name" json:"personName"`
	Age  int    `gorm:"column:person_age" json:"personAge"`
}
