package models

type RestaurantModel struct {
	ID uint
	Logo, Title, WorkingHours, Address, Email, Phone, PaymentMethods, WorkingDays string
	Lang, Long                                                              uint32
	IsActive, IsDeleted, IsPromo                                            bool
}