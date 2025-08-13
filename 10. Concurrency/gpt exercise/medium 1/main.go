package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <- chan int, results chan<- bool){
	for j := range jobs{
		fmt.Printf("Worker %d start %d job\n", id, j)
		time.Sleep(time.Second)
		fmt.Printf("Worker %d end %d job\n", id, j)
		results <- true
	}
}

func main(){
	numJobs := 5
	jobs := make(chan int, numJobs)
	results := make(chan bool, numJobs)

	// Deploy 2 workers
	for w := 1; w <= 2; w++{
		go worker(w, jobs, results)
	}

	// Sending 5 jobs to jobs channel
	for j := 1; j <= 5; j++{
		jobs <- j
	}

	close(jobs)

	for a := 1; a <= numJobs; a++{
		<- results
	}

	fmt.Println("All works has been done!")
}