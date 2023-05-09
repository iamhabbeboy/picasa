package api

import (
	"fmt"
	"main/internal"
	"os/user"

	"github.com/fsnotify/fsnotify"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type ConfigService struct {
	Db *viper.Viper
}

func LoadConfig() *viper.Viper {
	v := viper.New()
	h, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	appName := internal.APP_NAME
	configPath := fmt.Sprintf("%s/.%s", h.HomeDir, appName)

	v.SetConfigName("config")
	v.AddConfigPath(configPath)
	v.WatchConfig()
	v.OnConfigChange(func(_ fsnotify.Event) {})

	v.SetDefault("config.max_image", 10)
	v.SetDefault("config.interval", "5m")
	v.SetDefault("api.query", "technology")
	v.SetDefault("api.url", "https://api.unsplash.com/")
	v.SetDefault("config.image_path", fmt.Sprintf("%s/images", configPath))
	v.SetDefault("api.access_key", "Nw5jS2P4zr_oO_qbFt_39zyj7QTIMI49vYx5lCzxujY")
	v.SetDefault("api.secret_key", "pseMeAYqR4G1I8cx8vbwkm4HTs1o56NzW6ZiKGHCMNs")

	// Write the configuration options to a YAML file
	if err := v.WriteConfigAs(fmt.Sprintf("%s/config.yaml", configPath)); err != nil {
		log.Fatalf("Error writing configuration file: %s", err)
	}

	// Read the configuration options from the YAML file
	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("Error reading configuration file: %s", err)
	}
	return v
}

func NewConfigService() *ConfigService {
	config := LoadConfig()
	return &ConfigService{
		Db: config,
	}
}

func (c *ConfigService) Get(key string) string {
	return c.Db.GetString(key)
}

func (c *ConfigService) Set(key string, value string) error {
	c.Db.Set(key, value)
	err := c.Db.WriteConfig()
	if err != nil {
		log.Fatal("Error writing config file:", err)
	}
	return nil
}
