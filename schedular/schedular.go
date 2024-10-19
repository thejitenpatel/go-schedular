package schedular

import (
	"sync"
)

// Task represents the type of tasks that can be scheduled
type Task func()

// Scheduler struct to manage tasks and concurrency
type Schedular struct {
	maxConcurrent int // maximum number of concurrent tasks
	queue 				chan Task // task queue
	activeCount 	int // number of active tasks
	mu 						sync.Mutex // to synchronize activeCount updates
}

// NewScheduler initializes the scheduler with a maximum number of concurrent tasks
func NewSchedular(n int) *Schedular {
	return &Schedular{
		maxConcurrent: n,
		queue: make(chan Task, 100), // buffered channel for pending tasks
	}
}

// Start launches worker goroutines to handle tasks
func (s *Schedular) Start() {
	for i := 0; i < s.maxConcurrent; i++ {
		go s.worker()
	}
}

// Schedule adds a new task to the queue
func (s *Schedular) Schedule(task Task) {
	s.queue <- task
}

// worker function to process tasks
func (s *Schedular) worker() {
	for task := range s.queue {
		s.mu.Lock()
		if s.activeCount >= s.maxConcurrent {
			s.mu.Unlock()
			continue
		}

		s.activeCount++
		s.mu.Unlock()

		// Run the task in a separate goroutine so the worker remains free for new tasks
		go func (task Task)  {
			defer s.taskCompleted() // Ensure task count is decremented after the task completes
			task() // Execute the task
		}(task)
	}
}

// taskCompleted decreases the count of active tasks
func (s *Schedular) taskCompleted() {
	s.mu.Lock()
	s.activeCount--
	if s.activeCount <= s.maxConcurrent-2 {
		for i := 0; i < 2; i++ {
			if len(s.queue) > 0 {
				s.activeCount++
				go func(task Task)  {
					defer s.taskCompleted()
					task()
				}(<-s.queue)
			}
		}
	}
	s.mu.Unlock()
}
