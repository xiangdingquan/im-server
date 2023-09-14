package service

type Service struct {
}

func New() *Service {
	s := new(Service)
	return s
}

func (s *Service) Close() {
}
