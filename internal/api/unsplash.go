package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"

	log "github.com/sirupsen/logrus"
)

type UnleaseService struct {
	config *ConfigService
}

type Image struct {
	Urls struct {
		Raw  string `json:"raw"`
		Full string `json:"full"`
	} `json:"urls"`
	Links struct {
		Download string `json:"download"`
	} `json:"links"`
}

func NewUnsplashService() *UnleaseService {
	return &UnleaseService{
		config: &ConfigService{},
	}
}

func (u *UnleaseService) GetImages() error {
	c, _ := u.config.GetItem("picasa")
	// apiUrl := c.APIUrl       //u.config.Get("api.url")
	// query := c.Query         //u.config.Get("api.query")
	// maxImage := c.MaxImage   //u.config.Get("config.max_image")
	// accessKey := c.AccessKey //u.config.Get("api.access_key")
	// imagePath := c.ImagePath //u.config.Get("config.image_path")
	fmt.Println(c)
	// url := fmt.Sprintf("%s/photos/random?client_id=%s&count=%s&orientation=landscape&query=%s", apiUrl, accessKey, maxImage, query)
	// fmt.Println(url)
	// result, err := getImage(url)
	// if err != nil {
	// 	return err
	// }
	// var wg sync.WaitGroup
	// for key, v := range result {
	// 	wg.Add(1)
	// 	go download(v.Urls.Full, key, &wg, imagePath)
	// }
	// wg.Wait()
	return nil
}

func getImage(url string) ([]Image, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	var p []Image
	if err := json.NewDecoder(resp.Body).Decode(&p); err != nil {
		return nil, err
	}
	return p, nil
}

func download(image string, index int, wg *sync.WaitGroup, IMAGE_DIR string) {
	defer wg.Done()
	resp, err := http.Get(image)
	if err != nil {
		log.Fatal(err)
	}
	info := fmt.Sprintf("Downloading: %s/%v", IMAGE_DIR, index)
	fmt.Println(info)

	defer resp.Body.Close()
	f, err := os.Create(fmt.Sprintf("%s/%v.jpg", IMAGE_DIR, index))
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	_, err = io.Copy(f, resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Downloaded to: ", f.Name())
}
