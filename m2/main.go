package main

func process(numbers []int) []int {
 jobs := make(chan Job, len(numbers))
 results := make(chan Result, len(numbers))

 go worker(jobs, results)

 for _, n := range numbers {
  jobs <- Job{Value: n}
 }
 close(jobs)

 out := make([]int, 0, len(numbers))
 for range numbers {
  r := <-results
  out = append(out, r.Value)
 }
 return out
}