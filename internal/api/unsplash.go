package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"

	log "github.com/sirupsen/logrus"
)

type UnleaseService struct {
	apikey string
	path   string
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
	Apikey             string
	Path               string
}

func NewUnsplashService(apikey string, path string) *UnleaseService {
	return &UnleaseService{
		apikey: apikey,
		path:   path,
	}
}

func (u *UnleaseService) GetImages(imgConf ImageConfig) error {
	apiUrl := "https://api.unsplash.com"
	query := imgConf.Category
	imgCount := imgConf.TotalDownloadImage
	maxImage := strconv.Itoa(imgCount)
	accessKey := u.apikey // "Nw5jS2P4zr_oO_qbFt_39zyj7QTIMI49vYx5lCzxujY"

	var imagePath string
	if strings.Contains(u.path, ".picasa") {
		home, _ := os.UserHomeDir()
		fp := fmt.Sprintf("%s/.picasa/images/", home)
		imagePath = fp
	}

	imagePath = u.path
	url := fmt.Sprintf("%s/photos/random?client_id=%s&count=%s&orientation=landscape&query=%s", apiUrl, accessKey, maxImage, query)
	fmt.Println("Download...")
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
