package main

import (
  "tempconv"
  "fmt"
)

func main() {
  fmt.Println(tempconv.CToK(tempconv.FreezingC)) // "212°F"
  fmt.Printf("ZBrrrr! %v\n", tempconv.AbsoluteZeroC) // "Brrrr! -273.15°C"
  fmt.Println(tempconv.CToF(tempconv.BoilingC)) // "212°F"
}
