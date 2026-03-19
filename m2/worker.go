package main

type Job struct {
 Value int
}

type Result struct {
 Value int
}

func worker(jobs <-chan Job, results chan<- Result) {
 for j := range jobs {
  results <- Result{Value: j.Value * j.Value}
 }
}