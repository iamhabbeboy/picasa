package pkg

import (
	"fmt"
	"regexp"
	"strconv"
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
