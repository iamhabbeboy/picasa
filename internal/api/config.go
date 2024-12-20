package api

import (
	"encoding/json"
	"fmt"

	"github.com/akrylysov/pogreb"
)

type ConfigService struct {
	DB *pogreb.DB
}

type ConfigStorer struct {
	MaxImage  int
	Interval  string
	Query     string
	APIUrl    string
	ImagePath string
	AccessKey string
	SecretKey string
}

var LOCAL_DB_KEY = "picasa"

// func LoadConfig() *viper.Viper {
// v := viper.New()
// h, err := user.Current()
// if err != nil {
// 	log.Fatal(err)
// }
// appName := internal.APP_NAME
// configPath := fmt.Sprintf("%s/.%s", h.HomeDir, appName)

// v.SetConfigName("config")
// v.AddConfigPath(configPath)
// v.OnConfigChange(func(_ fsnotify.Event) {})

// v.SetDefault("config.max_image", 10)
// v.SetDefault("config.interval", "5m")
// v.SetDefault("api.query", "nature")
// v.SetDefault("api.url", "https://api.unsplash.com/")
// v.SetDefault("config.image_path", fmt.Sprintf("%s/images", configPath))
// v.SetDefault("api.access_key", "Nw5jS2P4zr_oO_qbFt_39zyj7QTIMI49vYx5lCzxujY")
// v.SetDefault("api.secret_key", "pseMeAYqR4G1I8cx8vbwkm4HTs1o56NzW6ZiKGHCMNs")

// // Write the configuration options to a YAML file
// if err := v.WriteConfigAs(fmt.Sprintf("%s/config.yaml", configPath)); err != nil {
// 	log.Printf("Error writing configuration file: %s", err)
// }

// // Read the configuration options from the YAML file
// if err := v.ReadInConfig(); err != nil {
// 	log.Printf("Error reading configuration file: %s", err)
// }
// return v
// }

// func NewConfig() *ConfigService {
// 	LoadDefaultConfig(&ConfigService{})
// 	return &ConfigService{
// 		// Db:      config,
// 		localDB: internal.DBConfig,
// 	}
// }

// func (c *ConfigService) Get(key string) string {
// 	return c.Db.GetString(key)
// }

// func (c *ConfigService) Set(key string, value string) error {
// 	c.Db.Set(key, value)
// 	err := c.Db.WriteConfig()
// 	if err != nil {
// 		log.Fatal("Error writing config file:", err)
// 	}
// 	return nil
// }

func (c *ConfigService) SetItem(key string, value interface{}) error {
	j, err := json.Marshal(value)
	if err != nil {
		return err
	}
	err = c.DB.Put([]byte(key), []byte(string(j)))
	if err != nil {
		return err
	}
	return nil
}

func (c *ConfigService) GetItem(key string) (ConfigStorer, error) {
	val, err := c.DB.Get([]byte(key))
	if err != nil {
		return ConfigStorer{}, err
	}
	if !json.Valid(val) {
		return ConfigStorer{}, fmt.Errorf("invalid json response")
	}
	var config ConfigStorer
	err = json.Unmarshal(val, &config)
	if err != nil {
		return ConfigStorer{}, err
	}
	return config, nil
}
