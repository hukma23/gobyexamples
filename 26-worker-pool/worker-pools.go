package main

import (
	"fmt"
	"time"
)

// in this example we'll look at how to implement a worker pool using goroutines and channels.
// Here's the worker, of which we'll run several concurrent instance.
// These workers will receive work on the jobs channels and send the corresponding result on results.
// We'll sleep a second perj job to simulate an expansive task.
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second * 4)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {
	// In order to use our pool of workers we need to send them work and collect their results.
	// We make 2 channels for this.
	const numJobs = 9
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// This start up 3 worker, initially blocked because there are no jobs yet.
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// Here we send 5 jobs and then close that channel to indicate that's all the work we have.
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}

	close(jobs)

	// Finally we collect all the results of the work.
	// This also ensures that the worker goroutines have finished.
	// An laternative way to wait for multiple goroutines is to use a waitGroup.
	for a := 1; a <= numJobs; a++ {
		<-results
	}
}
