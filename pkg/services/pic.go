package services

import "fmt"

type PicService struct {
}

func NewPicServer() *PicService {
	return &PicService{}
}

func (p *PicService) GetImages(q string) {
	fmt.Println("New Pic Service")
}
