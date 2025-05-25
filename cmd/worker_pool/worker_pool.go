package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func workers(id int, jobs <-chan int, results chan<- [3]int, wg *sync.WaitGroup){
	defer wg.Done()
	for job := range jobs{
		fmt.Printf("Worker #%v starting job #%v\n", id, job)
		time.Sleep(time.Duration(rand.Intn(1000))*time.Millisecond) // Work simulation
		fmt.Printf("Worker #%v finished job #%v\n", id, job)
		results <- [3]int{id, job, job*2}
	}
}


func resLogger(results <-chan [3]int, rl *sync.WaitGroup){
	defer rl.Done()
	for res := range results{
		fmt.Printf("result recieved by worker #%v of job #%v, final result is %v\n", res[0], res[1], res[2])
	}
}

func main (){
	const numJobs = 8
	const numWorkers = 3

	var jobs chan int = make(chan int)
	var results chan [3]int = make(chan [3]int)
	var wg sync.WaitGroup
	var rl sync.WaitGroup

	for w := range numWorkers{
		wg.Add(1)
		go workers(w, jobs, results, &wg)
	}
	rl.Add(1)
	go resLogger(results, &rl)

	for j := range numJobs{
		jobs <- j
	}
	close(jobs)
	wg.Wait()
	close(results)
	rl.Wait()
	
}