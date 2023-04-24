package pkg

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

func GetTimeToCrontabFormat(dur string) string {
	if dur == "" {
		return "*/5 * * * *"
	}

	re := regexp.MustCompile("[0-9]+")
	num := re.FindString(dur)
	n, _ := strconv.Atoi(num)

	ext := dur[len(dur)-1:]

	if ext == "h" && n >= 24 {
		return "* */24 * * *"
	}

	if ext == "h" {
		return fmt.Sprintf("* */%d * * *", n)
	}

	if ext == "m" && n < 60 {
		return fmt.Sprintf("*/%d * * * *", n)
	}

	return "*/5 * * * *"
}

func SetCronTab(timing string) {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	dir := currentDir + "/wallpaper set > /dev/null 2>&1"
	newJob := fmt.Sprintf("%s %s", timing, dir)
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

func HasCronjob(cronjob string) bool {
	cmd := exec.Command("crontab", "-l")
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		return false
	}

	return strings.Contains(stdout.String(), cronjob)
}
