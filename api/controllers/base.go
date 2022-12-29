package controllers

import "gorm.io/gorm"

type BaseController struct {
	DB *gorm.DB
}
