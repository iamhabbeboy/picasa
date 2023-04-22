/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io/ioutil"
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
	Short: "A command to set wallpaper",
	Long:  `Set wallpaper .`,
	Run: func(cmd *cobra.Command, args []string) {
		triggerAction(cmd)
	},
}

func triggerAction(cmd *cobra.Command) {
	format := "hm"
	interval := "5m"
	value := cmd.Flags().Lookup("interval").Value.String()
	if value != "" {
		interval = value
	}
	// timeFormat := interval[len(interval)-1:]

	if !strings.ContainsAny(interval, format) {
		log.Fatal("Interval must be in format in minutes or hours, example: 5m, 1h")
		os.Exit(1)
	}
	config := services.NewConfigService()
	config.Set("config.interval", interval)

	// -- Check if file exist in image directory
	// get image path from config
	path := config.Get("config.image_path")
	if path == "" {
		// call download image here
	}
	// set wallpaper
}

func init() {
	rootCmd.AddCommand(setCmd)
	setCmd.Flags().StringP("interval", "i", "", "set interval time to change wallpaper, default is 5m")
	// setCmd.PersistentFlags().String("foo", "", "A help for foo")
}

func setWallpaper() {
	dir := IMAGE_DIR + "/"
	f, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println(err.Error())
		log.Fatal(err)
	}

	random := rand.Intn(len(f))
	file := f[random].Name()
	filepath, _ := filepath.Abs(dir + "/" + file)
	fmt.Println(filepath)

	cmd := exec.Command("osascript", "-e", fmt.Sprintf("tell application \"Finder\" to set desktop picture to POSIX file \"%s\"", filepath))

	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(output))
}
