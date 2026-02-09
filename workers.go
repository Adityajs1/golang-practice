package main

import "fmt"

func worker(id int, jobs chan int, results chan int) {
    for job := range jobs {  // Keep taking jobs from channel
        fmt.Printf("Worker %d processing job %d\n", id, job)
        results <- job * 2  // Send result back
    }
}

func workers() {
    jobs := make(chan int, 5)
    results := make(chan int, 5)
    
    // Start 3 workers
    go worker(1, jobs, results)
    go worker(2, jobs, results)
    go worker(3, jobs, results)
    
    // Send 5 jobs
    for i := 1; i <= 5; i++ {
        jobs <- i
    }
    close(jobs)  
    

    for i := 1; i <= 5; i++ {
        result := <-results
        fmt.Println("Result:", result)
    }
}