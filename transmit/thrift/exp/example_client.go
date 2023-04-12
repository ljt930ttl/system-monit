package exp

import (
	"context"
	"fmt"
	"system-monit/transmit/thrift/exp/example"

	"github.com/apache/thrift/lib/go/thrift"
)

var defaultCtx = context.Background()

func RunClient(addr string) error {
	transport, err := thrift.NewTSocket(addr)
	if err != nil {
		return err
	}
	if err := transport.Open(); err != nil {
		return err
	}
	defer transport.Close()
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	client := example.NewExampleServiceClientFactory(transport, protocolFactory)
	request := &example.Example{
		Message: "Hello, world!",
	}
	response, err := client.Echo(defaultCtx, request)
	if err != nil {
		return err
	}
	fmt.Println("Server response:", response.Message)
	return nil
}

// func main() {
// 	if err := runClient(); err != nil {
// 		fmt.Println("Error running client:", err)
// 	}
// }
