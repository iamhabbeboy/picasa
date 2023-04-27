package services

import (
	"testing"

	"github.com/stretchr/testify/mock"
)

type MyMockedObject struct {
	mock.Mock
}

func TestService(t *testing.T) {
	// var buf bytes.Buffer
	// log.SetOutput(&buf)
	// svc := NewImageServicer("unsplash")
	testObj := new(MyMockedObject)
	testObj.On("GetImages", "nigeria").Return("nigeria")
	// assert.Equal(t, svc, &UnsplashService{}, "UnsplashService should be returned")
}
