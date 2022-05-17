package user

import "go.mongodb.org/mongo-driver/bson/primitive"

type Service struct {
	r *Repository
}

func NewService(r *Repository) *Service {
	return &Service{
		r: r,
	}
}

func (s *Service) Create(u *CreateUserDto, ip string) (interface{}, error) {
	return s.r.Create(u, ip)
}

func (s *Service) GetByEmail(e string) (User, error) {
	return s.r.GetByEmail(e)
}

func (s *Service) GetByUserName(u string) (User, error) {
	return s.r.GetByUserName(u)
}

func (s *Service) GetById(id primitive.ObjectID) (User, error) {
	return s.r.GetById(id)
}

func (s *Service) CheckPassword(u *User, p string) (bool, error) {
	return s.r.CheckPassword(u, p)
}
