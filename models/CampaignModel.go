package models

import "time"

// CampaingModel s
type CampaingModel struct {
	ID, ProductID uint
	Claim         int
	Description   string
	FinishDate    time.Time
}
