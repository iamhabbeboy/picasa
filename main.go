/*
Copyright © 2023 Abiodun Azeez iamhabbeboy@gmail.com
*/
package main

import (
	"fmt"
	"log"
	"main/cmd"
	"main/internal"
	"os"
)

func main() {
	if !createImageDir() {
		log.Fatal("Error occured while creating directory")
	}

	cmd.Execute()
}

func createImageDir() bool {
	app := internal.APP_NAME
	imgDir := fmt.Sprintf("./.%s/images", app)
	if _, err := os.Stat(imgDir); os.IsNotExist(err) {
		err := os.MkdirAll(imgDir, 0755)
		return err == nil
	}
	return true
}
