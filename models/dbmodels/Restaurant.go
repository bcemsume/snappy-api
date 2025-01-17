package dbmodels

import "github.com/jinzhu/gorm"

// Restaurant s
type Restaurant struct {
	*gorm.Model
	Title, WorkingHours, Address, Email, Phone, PaymentMethods, WorkingDays, Lang, Long string
	IsActive, IsDeleted, IsPromo                                                        bool
	Products                                                                            []Product
	Images                                                                              []*Image          `gorm:"auto_preload"`
	Users                                                                               []*RestaurantUser `gorm:"many2many:user_restaurants;auto_preload"`
}
