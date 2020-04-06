package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type CampaignModel struct {
	*gorm.Model
	ProductID  uint
	Claim      int
	FinishDate time.Time
}
