package main

import (
	"fmt"
	"time"
)

var (
	INTERVAL_SEC = 10
)

func main() {
	t := time.NewTicker(time.Duration(INTERVAL_SEC) * time.Second)
	for _ = range t.C {
		fmt.Println("PrintRoutine1")
	}
	/*for {
		fmt.Println("Main program is running...")
		time.Sleep(3 * time.Second)
	}*/
}
