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

func NewUnleaseService() *UnleaseService {
	return &UnleaseService{
		config: NewConfigService(),
	}
}

func (u *UnleaseService) GetImages() {
	apiUrl := u.config.Get("api.url")
	query := u.config.Get("api.query")
	maxImage := u.config.Get("config.max_image")
	accessKey := u.config.Get("api.access_key")
	imagePath := u.config.Get("config.image_path")

	url := fmt.Sprintf("%s/photos/random?client_id=%s&count=%s&orientation=landscape&query=%s", apiUrl, accessKey, maxImage, query)
	result := getImage(url)
	var wg sync.WaitGroup
	for key, v := range result {
		wg.Add(1)
		go download(v.Urls.Full, key, &wg, imagePath)
	}
	wg.Wait()
}

func getImage(url string) []Image {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("Unable to connect to the internet.")
	}
	var p []Image
	if err := json.NewDecoder(resp.Body).Decode(&p); err != nil {
		log.Fatal(err)
	}
	return p
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
