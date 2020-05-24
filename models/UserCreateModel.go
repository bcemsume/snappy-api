package models

type UserCreateModel struct {
	UserName, PhoneNumber, Password, DeviceID, FCMToken string
}
