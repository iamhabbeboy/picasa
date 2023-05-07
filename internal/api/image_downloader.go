package api

import "log"

type IimageDownloader interface {
	GetImages(q string)
}

func NewImageDownload(svc string) IimageDownloader {
	if svc == "" {
		log.Fatal("API Service not found")
	}
	app := map[string]IimageDownloader{
		"unsplash": NewUnleaseService(),
		"pixabay":  NewPixabayService(),
	}
	return app[svc]
}
