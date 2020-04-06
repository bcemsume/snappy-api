package models

import "time"

// RestaurantModel s
type ProductModel struct {
	RestaurantID, ID uint
	Description      string
	Price            float64
	FinishDate       time.Time
}
