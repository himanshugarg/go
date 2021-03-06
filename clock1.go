// Clock1 is a TCP server that periodically writes the time.
package main

import (
  "io"
  "log"
  "net"
  "os"
  "time"
)

func main() {
  port := "8000"
  if len(os.Args) == 2 {
    port = os.Args[1]
  }
  listener, err := net.Listen("tcp", "localhost:" + port)
  if err != nil {
    log.Fatal(err)
  }

  for {
    conn, err := listener.Accept()
    if err != nil {
      log.Print(err) // e.g., connection aborted
      continue
    }
    go handleConn(conn) // handle one connection at a time
  }
}

func handleConn(c net.Conn) {
  defer c.Close()
  for {
    _, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
    if err != nil {
      return
    }
    time.Sleep(1 * time.Second)
  }
}
