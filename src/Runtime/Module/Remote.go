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

	if !IncomingResponse.IsSuccess {
		return errors.New("the remote module failed to initialize")
	}

	return nil
}

func DeployBackend(RemoteModuleInfo IChatanium.ModuleInfo, Backend IChatanium.Backend) (uint16, error) {
	Log.Verbose.Printf("Deploying Backend: %s", RemoteModuleInfo.Name)

	var Random interface{} = rand.Intn(math.MaxUint16) + 1
	Port, ok := Random.(uint16)
	if !ok {
		return 0, errors.New("the automatically created port number does not fit the port range. Please try again")
	}

	Log.Verbose.Printf("port: %v", Port)

	if err := Backend.Init(RemoteModuleInfo); err != nil {
		return 0, err
	}

	if err := Backend.Connect(); err != nil {
		return 0, err
	}

	if err := rpc.Register(Backend); err != nil {
		return 0, err
	}

	l, err := net.Listen("tcp", fmt.Sprintf(":%v", &Port))
	if err != nil {
		return 0, err
	}
	defer l.Close()

	Log.Verbose.Printf("backend server listening on port %v", Port)

	for {
		conn, _ := l.Accept()
		go rpc.ServeConn(conn)
	}

	return Port, nil
}
