package main

import (
	"fmt"
	"sync"
	"time"
)

// Singleton instance
var (
	once      sync.Once
	scheduler *Scheduler
)

// Scheduler struct
type Scheduler struct {
	ticker *time.Ticker
	quit   chan struct{}
}

func NewScheduler() *Scheduler {
	once.Do(func() { // Ensures only one instance is created
		scheduler = &Scheduler{
			ticker: time.NewTicker(1 * time.Second),
			quit:   make(chan struct{}),
		}
		go scheduler.run() // Start the scheduler
	})
	return scheduler
}

// Scheduler logic
func (s *Scheduler) run() {
	fmt.Println("Scheduler started...")
	for {
		select {
		case <-s.ticker.C:
			fmt.Println("Running scheduled task...")
		case <-s.quit:
			fmt.Println("Scheduler stopped.")
			return
		}
	}
}

func (s *Scheduler) Stop() {
	close(s.quit)
	s.ticker.Stop()
}
