package models

import "time"

// ProductModel s
type ProductModel struct {
	RestaurantID, ID uint
	Description      string
	Price            float64
	FinishDate       time.Time
}
