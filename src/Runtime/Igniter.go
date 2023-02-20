package runtime

import (
	"flag"

	"antegral.net/chatanium/src/Runtime/Log"
)

var LOGMODE int

func Ignite() {
	InitLog()
}

func InitLog() {
	flag.IntVar(&LOGMODE, "logging-mode", 3, "Logging mode")
	flag.Parse()

	Log.Init(LOGMODE)
	Log.Info.Println("Starting Chatanium...")
}
