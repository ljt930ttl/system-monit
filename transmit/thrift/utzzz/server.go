package utzzz

import (
	"context"
	"fmt"

	"github.com/apache/thrift/lib/go/thrift"
)

var MethodMap = map[string]interface{}{
	"SM_SendHeartBeat": HeartBeat,
}

type utzzzHandler struct {
}

func (h *utzzzHandler) RequestFunc(ctx context.Context, pmaMsg *PMAMsg) error {
	fmt.Printf("Received message: %s\n", pmaMsg)
	// response := &PMAMsg{
	// 	Head:    pmaMsg.Head,
	// 	Src:     "TestDemo",
	// 	Targets: pmaMsg.Targets,
	// 	Content: pmaMsg.Content,
	// }
	return nil
}
func RunServer(addr string) error {
	handler := &utzzzHandler{}
	processor := NewPMAServiceProcessor(handler)
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
