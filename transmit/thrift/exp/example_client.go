package exp

import (
	"example"
	"fmt"

	"git.apache.org/thrift.git/lib/go/thrift"
)

func runClient() error {
	transport, err := thrift.NewTSocket("localhost:9090")
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
	response, err := client.Echo(request)
	if err != nil {
		return err
	}
	fmt.Println("Server response:", response.Message)
	return nil
}
func main() {
	if err := runClient(); err != nil {
		fmt.Println("Error running client:", err)
	}
}
