package services

type ImageServicer interface {
	GetImage()
	Download()
	ProcessImage()
}
