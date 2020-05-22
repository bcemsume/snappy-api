package dbmodels

import (
	"time"

	"github.com/jinzhu/gorm"
)

// User s
type User struct {
	*gorm.Model
	UserName, LastName, Name, SocialToken, Password, Email, PhoneNumber, DeviceID string
	BirthDay                                                                      time.Time
	Gender                                                                        byte
	SocialTokenType, UserType                                                     int16
	IsActive, IsDeleted                                                           bool
	ClaimEvents                                                                   []ClaimEvent
}
