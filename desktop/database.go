package main

import (
	"fmt"
	"log"
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

func (a *AppConfig) Init() {
	viper.SetDefault("app.name", "Picasa Desktop")
	viper.SetDefault("app.version", "0.1.0")
	viper.SetDefault("app.default_path", "~/.picasa/")

	viper.SetDefault("image.selected_abs_path", "")
	viper.SetDefault("image.last_downloads", "")

	viper.SetDefault("api.download_limit", 10)
	viper.SetDefault("api.image_category", "country")
	viper.SetDefault("api.source_apis", []string{"unsplash"})

	viper.SetConfigName("picasa")
	viper.AddConfigPath(".")
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
	/*var config AppConfig
	err := viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("Error unmarshalling config: %v", err)

		return AppConfig{}, errors.New(fmt.Sprintf("Error unmarshalling config: %v", err))
	}
	return config, nil
	*/
	return r, nil
}

func (a *AppConfig) Set(key string, value interface{}) {
	viper.Set(key, value)
	err := viper.WriteConfig()
	if err != nil {
		log.Fatalf("Error writing to config file: %v", err)
	}
	fmt.Printf("Set '%s' to '%v'\n", key, value)
}

func (a *AppConfig) delete(key string) {
	viper.Set(key, nil)
	err := viper.WriteConfig()
	if err != nil {
		log.Fatalf("Error deleting from config file: %v", err)
	}
	fmt.Printf("Deleted key '%s'\n", key)
}
