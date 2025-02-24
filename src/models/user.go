package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
	Pincode  int
	City     string
	State    string
	Country  string
}
