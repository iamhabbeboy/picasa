package api

type PixabayService struct {
}

func NewPixabayService() *PixabayService {
	return &PixabayService{}
}

func (p *PixabayService) GetImages(img ImageConfig) error {
	// TODO: Implementation here
	return nil
}
