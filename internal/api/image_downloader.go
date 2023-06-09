package api

import "log"

type IimageDownloader interface {
	GetImages()
}

func NewImageDownload(svc string) IimageDownloader {
	if svc == "" {
		log.Fatal("API Service not found")
	}
	app := map[string]IimageDownloader{
		"unsplash": NewUnsplashService(),
		"pixabay":  NewPixabayService(),
	}
	return app[svc]
}
