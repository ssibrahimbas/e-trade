package auth

import (
	"errors"

	"go.mongodb.org/mongo-driver/mongo"
	"ssibrahimbas.com/e-trade/internal/user"
	"ssibrahimbas.com/e-trade/services"
)

type Service struct {
	us  *user.Service
	jwt *services.Jwt
}

func NewService(us *user.Service) *Service {
	jwt := &services.Jwt{}
	return &Service{us: us, jwt: jwt}
}

func (s *Service) Login(u *user.LoginUserDto) (string, error) {
	user, err := s.us.GetByEmail(u.Email)
	if err != nil {
		return "", err
	}
	res, err := s.us.CheckPassword(&user, u.Password)
	if err != nil {
		return "", err
	}
	if res == false {
		return "", errors.New("Wrong password")
	}
	token, err := s.jwt.Sign(user)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *Service) Register(u *user.CreateUserDto, ip string) (any, error) {
	err := s.checkEmailIsExists(u.Email)
	if err != nil {
		return nil, err
	}
	err = s.checkUsernameIsExists(u.UserName)
	if err != nil {
		return nil, err
	}
	user, err := s.us.Create(u, ip)
	return user, err
}

func (s *Service) ForgotPassword() {

}

func (s *Service) ForgotUsername() {

}

func (s *Service) checkEmailIsExists(e string) error {
	_, err := s.us.GetByEmail(e)
	if err != mongo.ErrNoDocuments {
		return errors.New("Found account registered with this email.")
	}
	return nil
}

func (s *Service) checkUsernameIsExists(u string) error {
	_, err := s.us.GetByUserName(u)
	if err != mongo.ErrNoDocuments {
		return errors.New("Found account registered with this username.")
	}
	return nil
}
