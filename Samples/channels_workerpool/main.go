package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//Job struct
type Job struct {
	id       int
	randomno int
}

//Result struct
type Result struct {
	job         Job
	sumofdigits int
	workerNo    int
}

var jobs = make(chan Job, 10)
var results = make(chan Result, 10)

func digits(number int) int {
	sum := 0
	no := number
	for no != 0 {
		digit := no % 10
		sum += digit
		no /= 10
	}
	time.Sleep(1 * time.Second)
	return sum
}

func worker(wg *sync.WaitGroup, workerNo int) {

	for job := range jobs {
		output := Result{job, digits(job.randomno), workerNo}
		results <- output
	}
	wg.Done()
}

func createWorkerpool(nofWorkers int) {

	var wg sync.WaitGroup
	for i := 0; i < nofWorkers; i++ {
		wg.Add(1)
		go worker(&wg, i)
	}
	wg.Wait()
	close(results)
}

func allocate(nofJobs int) {

	for i := 0; i < nofJobs; i++ {
		randomno := rand.Intn(999)
		job := Job{i, randomno}
		jobs <- job
	}
	close(jobs)
}

func result(done chan bool) {
	for result := range results {
		fmt.Printf("Worker No %d,Job ID %d, input random no %d, sum of digits %d\n", result.workerNo, result.job.id, result.job.randomno, result.sumofdigits)
	}
	done <- true
}

func main() {
	startTime := time.Now()
	nofJobs := 100
	go allocate(nofJobs)
	done := make(chan bool)
	go result(done)
	nofWorkers := 10
	createWorkerpool(nofWorkers)
	<-done
	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("total time taken ", diff.Seconds(), "seconds")
}
