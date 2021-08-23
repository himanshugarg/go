// Outputs the commandline arguments 

package main

import (
  "fmt"
  "os"
  "strings"
)

func main() {
  fmt.Println(strings.Join(os.Args, " "))
}
