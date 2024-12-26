package internal

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
)

type AppConfig struct {
	App   AppConf `mapstructure:"app"`
	Image Image   `mapstructure:"image"`
	Api   Api     `mapstructure:"api"`
}

type AppConf struct {
	Name        string `mapstructure:"name"`
	Version     string `mapstructure:"version"`
	DefaultPath string `mapstructure:"default_path"`
	Interval    string `mapstructure:"interval"`
}

type Image struct {
	SelectedAbsPath string    `mapstructure:"selected_abs_path"`
	LastDownloads   time.Time `mapstructure:"last_downloads"`
}

type Api struct {
	DownloadLimit int      `mapstructure:"download_limit"`
	ImageCategory string   `mapstructure:"image_category"`
	SourceApis    []string `mapstructure:"source_apis"`
}

func (a *AppConfig) Init(sp string) {
	viper.SetDefault("app.name", "Picasa Desktop")
	viper.SetDefault("app.version", "0.1.0")
	// viper.SetDefault("app.default_path", "~/.picasa/")

	home, _ := os.UserHomeDir()
	fpt := fmt.Sprintf("%s/.picasa/images", home)

	viper.SetDefault("image.selected_abs_path", fpt)
	viper.SetDefault("image.last_downloads", "")
	viper.SetDefault("image.interval", "10m")

	viper.SetDefault("api.download_limit", 10)
	viper.SetDefault("api.download_interval", "1w")
	viper.SetDefault("api.unsplash_apikey", "Nw5jS2P4zr_oO_qbFt_39zyj7QTIMI49vYx5lCzxujY")
	viper.SetDefault("api.image_category", "country")
	viper.SetDefault("api.source_apis", []string{"unsplash"})

	var fp string = "."
	if sp != "" {
		fp = sp
	}

	viper.SetConfigName("picasa")
	viper.AddConfigPath(fp)
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig() // Read the config file
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("Config file not found, creating a new one")
			err = viper.SafeWriteConfig()
			if err != nil {
				log.Fatalf("Error creating config file: %v", err)
			}
		} else {
			log.Fatalf("Error reading config file: %v", err)
		}
	}
}

func (a *AppConfig) Get(key string) (interface{}, error) {
	if !viper.IsSet(key) {
		return nil, nil
	}
	r := viper.Get(key)
	return r, nil
}

func (a *AppConfig) Set(key string, value interface{}) {
	viper.Set(key, value)
	err := viper.WriteConfig()
	if err != nil {
		log.Fatalf("Error writing to config file: %v", err)
	}
}

func (a *AppConfig) delete(key string) {
	viper.Set(key, nil)
	err := viper.WriteConfig()
	if err != nil {
		log.Fatalf("Error deleting from config file: %v", err)
	}
	fmt.Printf("Deleted key '%s'\n", key)
}
