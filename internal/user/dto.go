package user

import "time"

type CreateUserDto struct {
	FirstName   string    `json:"firstName" validate:"required,min=2"`
	LastName    string    `json:"lastName" validate:"required,min=2"`
	UserName    string    `json:"userName" validate:"required,min=2"`
	DateOfBirth time.Time `json:"dateOfBirth,omitempty" validate:"required"`
	Email       string    `json:"email" validate:"required,email,min=6,max=50"`
	Phone       string    `json:"phone,omitempty" validate:"required,min=10,max=10"`
	Password    string    `json:"password,omitempty" validate:"required,min=8,max=16"`
	ipAddress   string
}

func (u *CreateUserDto) setIp(ip string) {
	u.ipAddress = ip
}

type LoginUserDto struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}
