package main

import (
	"io"
	"log"
	"net"
	"os"
	"strings"
	"fmt"
)

func main() {
	for _, arg := range os.Args[1:] {
		fmt.Println(strings.Split(arg, "="))
		loc, srv := strings.Split(arg, "=")[0], strings.Split(arg, "=")[1]
		conn, err := net.Dial("tcp", srv)
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()
		var prefixWriter PrefixWriter = loc.(string)
		mustCopy(prefixWriter, conn)
	}
}

type PrefixWriter string

func (prefix *PrefixWriter) Write(p []byte) (int, error) {
	return fmt.Fprintf(os.Stdout, "%s: %s", prefix, string(p))
}



func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
