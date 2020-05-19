package models

import (
	"strings"
	"time"
)

// UserProfileModel s
type UserProfileModel struct {
	ID                                                     uint
	UserName, LastName, Name, SocialToken, Password, Email string
	BirthDay                                               SpecialDate
	Gender                                                 byte
	SocialTokenType, UserType                              int16
	IsActive, IsDeleted                                    bool
}

type SpecialDate struct {
	time.Time
}

func (sd *SpecialDate) UnmarshalJSON(input []byte) error {
	strInput := string(input)
	strInput = strings.Trim(strInput, `"`)
	newTime, err := time.Parse("2006-01-02", strInput)
	if err != nil {
		return err
	}

	sd.Time = newTime
	return nil
}
