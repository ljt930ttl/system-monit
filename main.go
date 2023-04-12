package main

import (
	"flag"
	"fmt"
	"os"

	"system-monit/transmit/thrift/exp"
)

func Usage() {
	fmt.Fprint(os.Stderr, "Usage of ", os.Args[0], ":\n")
	flag.PrintDefaults()
	fmt.Fprint(os.Stderr, "\n")
}

func main() {
	flag.Usage = Usage
	server := flag.Bool("server", false, "Run server")
	addr := flag.String("addr", "localhost:9090", "Address to listen to")

	flag.Parse()

	if *server {
		if err := exp.RunServer(*addr); err != nil {
			fmt.Println("error running server:", err)
		}
	} else {
		if err := exp.RunClient(*addr); err != nil {
			fmt.Println("error running client:", err)
		}
	}
}
