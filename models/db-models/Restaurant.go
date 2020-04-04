package dbmodels

import "github.com/jinzhu/gorm"

// Restaurant s
type Restaurant struct {
	*gorm.Model
	Title, WorkingHours, Address, Email, Phone, PaymentMethods, WorkingDays string
	Lang, Long                                                              uint32
	IsActive, IsDeleted, IsPromo                                            bool
}
