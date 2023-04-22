package services

type UnleaseService struct {
	config *ConfigService
}

func NewUnleaseService() *UnleaseService {
	return &UnleaseService{
		config: NewConfigService(),
	}
}

func (u *UnleaseService) GetImage() {
}

func (u *UnleaseService) Download() {
}

func (u *UnleaseService) ProcessImage() {
}
