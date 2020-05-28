package models

// RestaurantModel asd
type RestaurantModel struct {
	ID                                                                                        uint
	Logo, Title, WorkingHours, Address, Email, Phone, PaymentMethods, WorkingDays, Lang, Long string
	IsActive, IsDeleted, IsPromo                                                              bool
	Images                                                                                    []Image
	Campaigns                                                                                 []CampaingModel
	Products                                                                                  []ProductModel
}
