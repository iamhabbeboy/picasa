/*
Copyright © 2023 Abiodun Azeez iamhabbeboy@gmail.com
*/
package main

import (
	"main/cmd"
	"os"
)

const IMAGE_DIR = "./.wallpaper/images"

func main() {
	go createImageDir()
	cmd.Execute()
}

func createImageDir() bool {
	if _, err := os.Stat(IMAGE_DIR); os.IsNotExist(err) {
		err := os.MkdirAll(IMAGE_DIR, 0755)
		return err == nil
	}
	return true
}
