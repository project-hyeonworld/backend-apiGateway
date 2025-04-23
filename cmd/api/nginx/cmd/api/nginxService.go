package nginx

type Service struct {
	biz OsBusiness
}

func NewService(biz OsBusiness) *Service {
	return &Service{biz: biz}
}

func (s *Service) Restart() error {
	if err := s.biz.RestartNginx(); err != nil {
		return err
	}
	return nil
}
