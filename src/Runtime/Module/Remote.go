package Module

import (
	"flag"
	"fmt"
	"net"
	"net/rpc"

	IChatanium "antegral.net/chatanium/src/Runtime/Interface"
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

func Listen() {
	rpc.Register()

	Port := flag.Int("port", 27493, "port")
	flag.Parse()

	l, err := net.Listen("tcp", fmt.Sprintf(":%v", &Port))
	if err != nil {
	}
	defer l.Close()

	for {
		conn, _ := l.Accept()
		go rpc.ServeConn(conn)
	}
}
