package services

import "log"

type ImageServicer interface {
	GetImages(q string)
}

func NewImageServicer(svc string) ImageServicer {
	if svc == "" {
		log.Fatal("Service not found")
	}
	app := map[string]ImageServicer{
		"unsplash": NewUnleaseService(),
		"pic":      NewPicServer(),
	}
	return app[svc]
}
