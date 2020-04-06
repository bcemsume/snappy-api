package dbmodels

import "github.com/jinzhu/gorm"

// Image s
type Image struct {
	*gorm.Model
	RestaurantID uint
	ImageURL     string
	Order, Type  byte
	Restaurant   Restaurant
}
