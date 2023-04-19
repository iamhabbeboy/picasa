/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Run wallpaper as a cron job",
	Long:  `Wallpaper runner as a cron job and set wallpaper every 5 minutes`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("start called")
		cronTab("*/5 * * * *")
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}

func cronTab(interval string) {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	dir := currentDir + "/wallpaper > /dev/null 2>&1"
	newJob := fmt.Sprintf("*/5 * * * * %s", dir)
	cmd := exec.Command("crontab", "-l")
	stdout, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	newJobs := string(stdout) + newJob + "\n"
	cmd = exec.Command("crontab", "-")
	cmd.Stdin = strings.NewReader(newJobs)
	stdout, err = cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(stdout))
}
