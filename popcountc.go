package main

import (
  "fmt"
  "popcount"
)

func main() {
  var i uint64
  for ; i < 1000000; i++ {
    fmt.Println(popcount.PopCount(i));
  }
}
