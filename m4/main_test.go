package main

import "testing"

func TestCalculate(t *testing.T) {
 in := Input{Numbers: []int{1, 2, 3}}
 out := calculate(in)
 if out.Sum != 14 {
  t.Fail()
 }
}