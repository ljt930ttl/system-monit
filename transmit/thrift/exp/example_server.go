package main

import (
	"example"
	"fmt"

	"git.apache.org/thrift.git/lib/go/thrift"
)

type exampleHandler struct{}

func (h *exampleHandler) Echo(request *example.Example) (*example.Example, error) {
	fmt.Printf("Received message: %s\n", request.Message)
	response := &example.Example{
		Message: "Echo: " + request.Message,
	}
	return response, nil
}
func runServer() error {
	handler := &exampleHandler{}
	processor := example.NewExampleServiceProcessor(handler)
	transport, err := thrift.NewTServerSocket(":9090")
	if err != nil {
		return err
	}
	server := thrift.NewTSimpleServer2(processor, transport)
	fmt.Println("Starting the server...")
	return server.Serve()
}
func main() {
	if err := runServer(); err != nil {
		fmt.Println("Error running server:", err)
	}
}
