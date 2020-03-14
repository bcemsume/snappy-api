package dbmodels

import "github.com/jinzhu/gorm"

type User struct {
	UserName, LastName, Name, SocialToken, Password string
	CreatedBy, UpdatedBy, TitleID, TokenType        int16
	IsActive, IsDeleted                             bool
	*gorm.Model
}
