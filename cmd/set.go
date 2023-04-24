/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"main/pkg"
	"main/pkg/services"
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
	Long:  `Set wallpaper .`,
	Run: func(cmd *cobra.Command, args []string) {
		config := services.NewConfigService()
		triggerAction(cmd, config)
	},
}

func triggerAction(cmd *cobra.Command, config *services.ConfigService) {
	format := "hm"
	interval := "5m"
	value := cmd.Flags().Lookup("interval").Value.String()
	if value != "" {
		interval = value
	}

	if len(interval) > 3 {
		log.Fatal("The maximum chars is 3, Use this formating - 5m, 30m, 1h, 24h, etc")
	}

	if !strings.ContainsAny(interval, format) {
		log.Fatal("Interval must be in format in minutes or hours, example: 5m, 1h")
	}

	config.Set("config.interval", interval)

	res := pkg.GetTimeToCrontabFormat(interval)
	if !pkg.HasCronjob(res) {
		pkg.SetCronTab(res)
	}

	p := config.Get("config.image_path")
	if p == "" {
		log.Fatal("Wallpaper: config is broken, please check your config file")
	}

	if hasImageDownloaded(p) {
		setWallpaper(p)
	}
}

func init() {
	rootCmd.AddCommand(setCmd)
	setCmd.Flags().StringP("interval", "i", "", "set interval time to change wallpaper, default is 5m")
}

func hasImageDownloaded(p string) bool {
	filePath := fmt.Sprintf("%s/0.jpg", p)
	if _, err := os.Stat(filePath); err != nil {
		return os.IsNotExist(err)
	}
	return true
}

func setWallpaper(path string) {
	dir := path + "/"
	f, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println(err.Error())
		log.Fatal(err)
	}

	random := rand.Intn(len(f))
	file := f[random].Name()
	filepath, _ := filepath.Abs(dir + "/" + file)

	cmd := exec.Command("osascript", "-e", fmt.Sprintf("tell application \"Finder\" to set desktop picture to POSIX file \"%s\"", filepath))

	_, err = cmd.Output()
	if err != nil {
		fmt.Println(err)
		return
	}
}
