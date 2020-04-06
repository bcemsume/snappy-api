package models

// ImageModel s
type ImageModel struct {
	ID, RestaurantID uint
	ImageURL         string
	Order, Type      byte
}
