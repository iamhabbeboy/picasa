/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io/ioutil"
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
		setWallpaper()
	},
}

func init() {
	rootCmd.AddCommand(setCmd)
	setCmd.Flags().StringP("title", "t", "", "specify task title / heading")
	// setCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
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
