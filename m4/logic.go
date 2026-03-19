package main

type Input struct {
 Numbers []int `json:"numbers"`
}

type Output struct {
 Sum int `json:"sum"`
}

func calculate(input Input) Output {
 sum := 0
 for _, n := range input.Numbers {
  sum += n * n
 }
 return Output{Sum: sum}
}