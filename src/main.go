package chatanium

import (
	"flag"

	"antegral.net/chatanium/src/Runtime/Log"
)

var LoggingMode int

func main() {
	Log.Init()
}

func InitLog() {
	CheckFlags()
	Log.Init(LoggingMode)
	Log.Info.Println("Starting Chatanium...")
}

func CheckFlags() {
	flag.IntVar(&LoggingMode, "logging-mode", 3, "Logging mode")
	flag.Parse()
}
