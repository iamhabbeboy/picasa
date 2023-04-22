package services

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type ConfigService struct {
	db *viper.Viper
}

func LoadConfig() (config interface{}) {
	v := viper.New()

	v.SetConfigName("config")
	v.AddConfigPath("./.wallpaper_config")

	v.SetDefault("config.max_image", 5)
	v.SetDefault("api.url", "http://localhost")
	v.SetDefault("api.access_token", 5432)
	v.SetDefault("api.secret_key", "postgres")
	v.SetDefault("config.image_dir", "/Users/test/Pictures/wallpaper")

	// Write the configuration options to a YAML file
	if err := v.WriteConfigAs("./.wallpaper_config/config.yaml"); err != nil {
		log.Fatalf("Error writing configuration file: %s", err)
	}

	// Read the configuration options from the YAML file
	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("Error reading configuration file: %s", err)
	}
	return
}

func NewConfigService() *ConfigService {
	LoadConfig()
	return &ConfigService{
		db: viper.New(),
	}
}

func (c *ConfigService) Set(data map[string]string) error {
	fmt.Println(data)
	// c.db.set
	return nil
}
