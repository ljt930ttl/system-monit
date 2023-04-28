package main

import (
	"flag"
	"fmt"
	"os"

	"system-monit/transmit/thrift/utzzz"
)

func Usage() {
	fmt.Fprint(os.Stderr, "Usage of ", os.Args[0], ":\n")
	flag.PrintDefaults()
	fmt.Fprint(os.Stderr, "\n")
}

func main() {
	flag.Usage = Usage
	// server := flag.Bool("server", false, "Run server")
	addr := flag.String("addr", "10.8.4.208:6100", "Address to listen to")

	flag.Parse()
	if err := utzzz.RunServer(*addr); err != nil {
		fmt.Println("error running server:", err)
	}

	// if *server {
	// 	if err := utzzz.RunServer(*addr); err != nil {
	// 		fmt.Println("error running server:", err)
	// 	}
	// } else {
	// 	if err := utzzz.RunClient(*addr); err != nil {
	// 		fmt.Println("error running client:", err)
	// 	}
	// }
}
