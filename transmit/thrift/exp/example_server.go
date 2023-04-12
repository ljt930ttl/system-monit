package exp

import (
	"context"
	"fmt"
	"system-monit/transmit/thrift/exp/example"

	"github.com/apache/thrift/lib/go/thrift"
)

type exampleHandler struct{}

func (h *exampleHandler) Echo(ctx context.Context, request *example.Example) (*example.Example, error) {
	fmt.Printf("Received message: %s\n", request.Message)
	response := &example.Example{
		Message: "Echo: " + request.Message,
	}
	return response, nil
}
func RunServer(addr string) error {
	handler := &exampleHandler{}
	processor := example.NewExampleServiceProcessor(handler)
	transport, err := thrift.NewTServerSocket(addr)
	if err != nil {
		return err
	}
	server := thrift.NewTSimpleServer2(processor, transport)
	fmt.Println("Starting the server...")
	return server.Serve()
}

// func main() {
// 	if err := runServer(); err != nil {
// 		fmt.Println("Error running server:", err)
// 	}
// }
