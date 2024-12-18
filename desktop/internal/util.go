package internal

import (
	"bytes"
	"desktop/internal/api"
	"fmt"
	"log"
	"os/exec"
	"os/user"
	"regexp"
	"strconv"
	"strings"

	"github.com/akrylysov/pogreb"
	"github.com/sirupsen/logrus"
)

const (
	APP_NAME    = "picasa"
	CRON_WEEKLY = "0 0 * * 0"
)

var DBConfig *api.ConfigService

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

func InitPogrebDB() {
	if DBConfig == nil {
		db, err := pogreb.Open("picasa.db", nil)
		if err != nil {
			logrus.Fatal(err.Error())
		}
		c := &api.ConfigService{
			DB: db,
		}
		LoadDefaultConfig(c)
		DBConfig = c
	}
}

func LoadDefaultConfig(c *api.ConfigService) {
	appName := APP_NAME
	h, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	configPath := fmt.Sprintf("%s/.%s", h.HomeDir, appName)
	conf := api.ConfigStorer{
		MaxImage:  10,
		Interval:  "5m",
		Query:     "cars",
		APIUrl:    "https://api.unsplash.com/",
		ImagePath: fmt.Sprintf("%s/images", configPath),
		AccessKey: "Nw5jS2P4zr_oO_qbFt_39zyj7QTIMI49vYx5lCzxujY",
		SecretKey: "pseMeAYqR4G1I8cx8vbwkm4HTs1o56NzW6ZiKGHCMNs",
	}
	err = c.SetItem("picasa", conf)
	if err != nil {
		log.Printf("error occured while storing config: %s", err.Error())
	}
}
