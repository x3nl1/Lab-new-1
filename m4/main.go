package main

import (
 "encoding/json"
 "os"
)

func main() {
 var input Input
 json.NewDecoder(os.Stdin).Decode(&input)
 output := calculate(input)
 json.NewEncoder(os.Stdout).Encode(output)
}