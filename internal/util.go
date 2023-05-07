package internal

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

const (
	APP_NAME    = "rwallpaper"
	CRON_WEEKLY = "0 0 * * 0"
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

func SetCronTab(timing string, command string) {
	newJob := fmt.Sprintf("%s %s", timing, command)
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

func RemoveCronTab(cronjob string) bool {
	if HasCronjob(cronjob) {
		cli := exec.Command("crontab", "-l")
		var stdout, stderr bytes.Buffer
		cli.Stdout = &stdout
		cli.Stderr = &stderr
		err := cli.Run()
		if err != nil {
			return false
		}

		lines := strings.Split(stdout.String(), "\n")
		var newLines []string
		for _, line := range lines {
			if !strings.Contains(line, cronjob) {
				newLines = append(newLines, line)
			}
		}

		newCron := strings.Join(newLines, "\n")
		cmd1x := exec.Command("crontab", "-")
		cmd1x.Stdin = strings.NewReader(newCron)
		var stdot []byte
		stdot, err = cmd1x.Output()
		if err != nil {
			return false
		}
		fmt.Println(string(stdot))
		return true
	}
	return false
}

func HasLetters(arg string) bool {
	re := regexp.MustCompile("[a-zA-Z]+")
	num := re.FindString(arg)
	return num != ""
}
