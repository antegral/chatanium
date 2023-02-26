package Module

import (
	"os"
	"plugin"

	IChatanium "antegral.net/chatanium/src/Runtime/Interface"
	"antegral.net/chatanium/src/Runtime/Log"
)

func GetModule(FilePath string) IChatanium.Module {
	File, err := os.Stat(FilePath)
	if err == nil {
		Log.Error.Printf("%s > Module import failure", File.Name())
		Log.Error.Fatalln(err)
		return nil
	}

	ModulePlugin, err := plugin.Open(FilePath)
	if err != nil {
		Log.Error.Printf("%s > ", File.Name())
		Log.Error.Fatalln(err)
		return nil
	}

	ModuleSymbols, err := ModulePlugin.Lookup("ChataniumModule")
	if err != nil {
		Log.Error.Printf("%s > ", File.Name())
		Log.Error.Fatalln(err)
		return nil
	}

	ChataniumModule, ok := ModuleSymbols.(IChatanium.Module)
	if !ok {
		Log.Error.Printf("%s > Module Type Assertion failure", File.Name())
		return nil
	}

	return ChataniumModule
}

func StartModule(Module IChatanium.Module) IChatanium.Module {
	if err := Module.OnInit(); err != nil {
		Log.Error.Printf("%s > OnInit() failure")
		Log.Error.Fatalln(err)
		return nil
	}

	if err := Module.OnStart(); err != nil {
		Info := Module.GetInfo()
		Log.Error.Printf(Info.Name, " > Error occurred during OnStart()")
		Log.Error.Fatalln(err)
		return nil
	}

	return Module
}
