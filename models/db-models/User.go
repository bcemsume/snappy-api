package dbmodels

import "github.com/jinzhu/gorm"

// User s
type User struct {
	*gorm.Model
	UserName, LastName, Name, SocialToken, Password string
	SocialTokenType                                 int16
	IsActive, IsDeleted                             bool
}
