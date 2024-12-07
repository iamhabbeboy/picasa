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

type ImageConfig struct {
	Category           string
	TotalDownloadImage int
}

func NewUnsplashService() *UnleaseService {
	return &UnleaseService{
		config: &ConfigService{},
	}
}

func (u *UnleaseService) GetImages(imgConf ImageConfig) error {
	//c, _ := u.config.GetItem("picasa")

	// access_key: Nw5jS2P4zr_oO_qbFt_39zyj7QTIMI49vYx5lCzxujY
	// api_key: pseMeAYqR4G1I8cx8vbwkm4HTs1o56NzW6ZiKGHCMNs
	// url: https://api.unsplash.com
	apiUrl := "https://api.unsplash.com"                       //c.APIUrl                     //u.config.Get("api.url")
	query := imgConf.Category                                  //u.config.Get("api.query")
	maxImage := "3"                                            //imgConf.TotalDownloadImage                     //u.config.Get("config.max_image")
	accessKey := "Nw5jS2P4zr_oO_qbFt_39zyj7QTIMI49vYx5lCzxujY" //:= c.AccessKey               //u.config.Get("api.access_key")

	home, _ := os.UserHomeDir()
	fp := fmt.Sprintf("%s/.picasa/", home)
	imagePath := fp
	url := fmt.Sprintf("%s/photos/random?client_id=%s&count=%s&orientation=landscape&query=%s", apiUrl, accessKey, maxImage, query)
	fmt.Println(url)

	result, err := getImage(url)
	if err != nil {
		return err
	}
	var wg sync.WaitGroup
	for key, v := range result {
		wg.Add(1)
		go download(v.Urls.Full, key, &wg, imagePath)
	}
	wg.Wait()
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
