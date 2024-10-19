package main

import (
	"fmt"
	"goSched/schedular"
	"time"
)

func main() {
	// Initialize a scheduler with max 3 concurrent tasks
	s := schedular.NewSchedular(3)
	s.Start()

	// Simulate tasks with different execution times
	for i := 0; i <= 10; i++ {
		taskID := i
		s.Schedule(func() {
			fmt.Printf("Task %d started\n", taskID)
			time.Sleep(time.Duration(taskID) * time.Second) // Simulate task
			fmt.Printf("Task %d completed\n", taskID)
		})
	}
	
	// Wait for tasks to complete
	time.Sleep(30 * time.Second)
}
