package cac

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Add(companyVanityName string, userID int64) error {
	return nil
}
