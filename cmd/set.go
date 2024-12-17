/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"main/internal"
	"main/internal/api"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set random wallpaper",
	Long:  `Set wallpaper`,
	Run: func(cmd *cobra.Command, args []string) {
		config := &api.ConfigService{}
		triggerAction(cmd, config)
	},
}

func triggerAction(cmd *cobra.Command, config *api.ConfigService) {
	format := "hm"
	interval := "5m"
	c, _ := config.GetItem("picasa")
	value := cmd.Flags().Lookup("interval").Value.String()
	if value != "" {
		intervalValueFromConfig := c.Interval //config.Get("config.interval")
		if intervalValueFromConfig != "" {
			interval = intervalValueFromConfig
		} else {
			interval = value
		}
	}

	if len(interval) > 3 {
		log.Fatal("The maximum chars is 3, Use this formating - 5m, 30m, 1h, 24h, etc")
	}

	if !strings.ContainsAny(interval, format) {
		log.Fatal("Interval must be in format in minutes or hours, example: 5m, 1h")
	}

	// config.Set("config.interval", interval)
	config.SetItem("picasa", api.ConfigStorer{Interval: interval})

	res := internal.GetTimeToCrontabFormat(interval)

	if !internal.HasCronjob(res) {
		r := fmt.Sprintf("/usr/local/bin/%s set", internal.APP_NAME)
		internal.SetCronTab(res, r)
	}

	p := c.ImagePath //config.Get("config.image_path")
	if p == "" {
		log.Fatal("Picasa: config is broken, please check your config file")
	}
	fmt.Println(c)

	// if !hasImageDownloaded(p) {
	// 	HandleDownloadProcess()
	// }
	// setWallpaper(p)
}

func init() {
	rootCmd.AddCommand(setCmd)
	setCmd.Flags().StringP("interval", "i", "", "set interval time to change wallpaper, default is 5m")
	setCmd.Flags().StringP("query", "q", "", "set a query filter for image type, default is nature")
}

func hasImageDownloaded(p string) bool {
	filePath := fmt.Sprintf("%s/0.jpg", p)
	if _, err := os.Stat(filePath); err != nil {
		return false
	}
	return true
}

func setWallpaper(path string) {
	dir := path + "/"
	f, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	random := rand.Intn(len(f))
	file := f[random].Name()
	filepath, _ := filepath.Abs(dir + "/" + file)

	cmd := exec.Command("osascript", "-e", fmt.Sprintf("tell application \"Finder\" to set desktop picture to POSIX file \"%s\"", filepath))

	_, err = cmd.Output()
	if err != nil {
		log.Fatal(err)
		return
	}
}
