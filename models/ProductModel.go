package models

import (
	"time"
)

type ProductModel struct {
	ID          uint
	Claim       int
	Description string
	FinishDate  time.Time
}
