package main

import (
	"fmt"

	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name    string `gorm:"column:name"`
	Age     int    `gorm:"column:age"`
	Address string `gorm:"column:address"`
}

// hooks
func (u *Student) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Println("Student BeforeCreate")
	return nil
}

func (u *Student) AfterCreate(tx *gorm.DB) (err error) {
	fmt.Println("Student AfterCreate")
	return nil
}
