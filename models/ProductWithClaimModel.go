package models

import (
	"time"
)

// ProductWithClaimModel s
type ProductWithClaimModel struct {
	ID          uint
	Claim       int
	Description string
	Price       float64
	FinishDate  time.Time
}
