package user

import "time"

type User struct {
	Id          string    `json:"id,omitempty" bson:"_id,omitempty"`
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	UserName    string    `json:"userName" validate:"required,min=2"`
	DateOfBirth time.Time `json:"dateOfBirth,omitempty"`
	Email       string    `json:"email"`
	Phone       string    `json:"phone,omitempty"`
	Password    []byte    `json:"password,omitempty"`
	IpAddress   string    `json:"ipAddress"`
	Active      bool      `json:"active,omitempty"`
}
