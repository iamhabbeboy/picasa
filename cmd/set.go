/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"main/services"
	"math/rand"
	"os/exec"
	"path/filepath"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "A command to set wallpaper",
	Long:  `Set wallpaper .`,
	Run: func(cmd *cobra.Command, args []string) {
		dft := "5m"
		if len(args) > 0 {
			dft = args[0]
		}
		c := map[string]string{
			"interval": dft,
		}
		config := services.NewConfigService()
		if err := config.Set(c); err != nil {
			fmt.Println("viola")
		}
		// setWallpaper()
	},
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
