package dbmodels

import (
	"github.com/jinzhu/gorm"
)

// RestaurantUser s
type RestaurantUser struct {
	*gorm.Model
	UserName, LastName, Name, Password, Email string
	IsActive                                  bool
	Restaurants                               []*Restaurant `gorm:"many2many:user_restaurants;"`
}
