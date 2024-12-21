package api

import "log"

type ImageDownloader interface {
	GetImages(img ImageConfig) error
}

func NewImageDownload(svc string, path string, apikey string) ImageDownloader {
	if svc == "" {
		log.Fatal("API Service not found")
	}
	app := map[string]ImageDownloader{
		"unsplash": NewUnsplashService(apikey, path),
		"pixabay":  NewPixabayService(),
	}
	return app[svc]
}
