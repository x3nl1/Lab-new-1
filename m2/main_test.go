package main

import "testing"

func TestProcess(t *testing.T) {
 res := process([]int{1, 2, 3})
 if len(res) != 3 {
  t.Fail()
 }
}