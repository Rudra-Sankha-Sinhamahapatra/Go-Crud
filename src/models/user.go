package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `gorm:"uniqueIndex" json:"email"`
	Password string `json:"password"`
	Pincode  int    `json:"pincode"`
	City     string `json:"city"`
	State    string `json:"state"`
	Country  string `json:"country"`
}
