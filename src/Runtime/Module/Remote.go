package Module

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"net"
	"net/rpc"

	IChatanium "antegral.net/chatanium/src/Runtime/Interface"
	"antegral.net/chatanium/src/Runtime/Log"
)

func Connect(Port uint16) error {
	Client, err := rpc.Dial("tcp", fmt.Sprintf("127.0.0.1:%v", &Port))
	if err != nil {
		return err
	}

	defer Client.Close()

	IncomingResponse := new(IChatanium.RemoteModuleResponse)

	Client.Call("Module.OnInit", nil, IncomingResponse)

	if IncomingResponse.IsSuccess {
		return nil
	}
}

func DeployBackend(Name string, Backend IChatanium.Backend) (error, uint16) {
	var Random interface{} = rand.Intn(math.MaxUint16) + 1
	Port, ok := Random.(uint16)

	if ok != true {
		return errors.New("The automatically created port number does not fit the port range. Please try again."), 0
	}

	Log.Verbose.Printf("Deploying Backend: %s", Name)

	if err := Backend.Init(IChatanium.ModuleInfo{
		Name:        "",
		Description: "",
		Version:     "1.0.0",
		Tags:        nil,
		Commands:    nil,
	}); err != nil {
		return err, 0
	}

	if err := Backend.Connect(); err != nil {
		return err, 0
	}

	rpc.Register(Backend)

	l, err := net.Listen("tcp", fmt.Sprintf(":%v", &Port))
	if err != nil {
		return err, 0
	}
	defer l.Close()

	for {
		conn, _ := l.Accept()
		go rpc.ServeConn(conn)
	}

	return nil, Port
}
