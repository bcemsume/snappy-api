package models

// ImageModel s
type ImageModel struct {
	RestaurantID uint
	Images       []Image
}

// Image s
type Image struct {
	ImageURL    string
	ID          uint
	Order, Type byte
}
