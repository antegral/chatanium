package Module

import (
	"fmt"
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

func DeployBackend(Name string, Backend IChatanium.Backend) error {
	Port := rand.Intn(65535-1) + 1

	Log.Verbose.Printf("Deploying Backend: %s", Name)

	if err := Backend.Init(IChatanium.ModuleInfo{
		Name:        "",
		Description: "",
		Version:     "1.0.0",
		Tags:        nil,
		Commands:    nil,
	}); err != nil {
		return err
	}

	if err := Backend.Connect(); err != nil {
		return err
	}

	rpc.Register(Backend)

	l, err := net.Listen("tcp", fmt.Sprintf(":%v", &Port))
	if err != nil {
	}
	defer l.Close()

	for {
		conn, _ := l.Accept()
		go rpc.ServeConn(conn)
	}
}
