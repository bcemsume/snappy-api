package dbmodels

import "github.com/jinzhu/gorm"

// Image s
type Image struct {
	*gorm.Model
	Restaurant  Restaurant `gorm:"association_autoupdate:false"`
	ImageURL    string
	Order, Type byte
}
