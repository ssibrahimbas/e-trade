package product

type Service struct {
	r *Repository
}

func NewService(r *Repository) *Service {
	return &Service{
		r: r,
	}
}

func (s *Service) Create(p *CreateProductDto) (interface{}, error) {
	return s.r.Create(p)
}

func (s *Service) GetAll() ([]Product, error) {
	return s.r.GetAll()
}

func (s *Service) GetDetail(p *GetDetailProductDto) (Product, error) {
	return s.r.GetDetail(p)
}