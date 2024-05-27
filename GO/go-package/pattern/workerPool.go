package pattern

import (
	"fmt"
	"sync"
)

// Worker represents a worker goroutine that processes tasks
func worker(id int, tasks <-chan Task, wg *sync.WaitGroup) {
	defer wg.Done() // Notify WaitGroup when worker has finished

	for task := range tasks {
		// Process task
		fmt.Printf("Worker %d processing task %d\n", id, task.ID)
	}
}

// Task represents a unit of work to be performed
type Task struct {
	ID int
}

func WoorkPool() {
	numWorkers := 3 // Number of worker goroutines
	numTasks := 10  // Number of tasks to be processed
	taskQueue := make(chan Task, numTasks)
	var wg sync.WaitGroup // WaitGroup to wait for all workers to finish

	// Start worker goroutines
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, taskQueue, &wg)
	}

	// Enqueue tasks
	for i := 1; i <= numTasks; i++ {
		taskQueue <- Task{ID: i}
	}
	close(taskQueue) // Close the task queue channel to signal completion

	// Wait for all workers to finish
	wg.Wait()
	fmt.Println("All tasks have been completed.")
}
