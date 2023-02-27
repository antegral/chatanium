package Module

import (
	"os"
	"plugin"

	IChatanium "antegral.net/chatanium/src/Runtime/Interface"
	"antegral.net/chatanium/src/Runtime/Log"
)

func Get(FilePath string) IChatanium.Module {
	Log.Verbose.Printf("Getting module: %s", FilePath)

	File, err := os.Stat(FilePath)
	if err == nil {
		Log.Error.Printf("%s > Module import failure", File.Name())
		Log.Error.Fatalln(err)
		return nil
	}

	Log.Verbose.Printf("Valid file. Opening plugin...")

	ModulePlugin, err := plugin.Open(FilePath)
	if err != nil {
		Log.Error.Printf("%s > ", File.Name())
		Log.Error.Fatalln(err)
		return nil
	}

	Log.Verbose.Printf("Looking up module symbols...")

	ModuleSymbol, err := ModulePlugin.Lookup("ChataniumModule")
	if err != nil {
		Log.Error.Printf("%s > ", File.Name())
		Log.Error.Fatalln(err)
		return nil
	}

	Log.Verbose.Printf("Type-Casting Module...")

	ChataniumModule, ok := ModuleSymbol.(IChatanium.Module)
	if !ok {
		Log.Error.Fatalf("%s > Module Type-Casting failure", File.Name())
		return nil
	}

	Log.Verbose.Printf("Successfully converted to module.")

	return ChataniumModule
}

func Start(Module IChatanium.Module) IChatanium.Module {
	Log.Verbose.Printf("Starting module: %s", Module.GetInfo().Name)
	if err := Module.OnInit(); err != nil {
		Log.Error.Printf("%s > OnInit() failure")
		Log.Error.Fatalln(err)
		return nil
	}

	if err := Module.OnStart(); err != nil {
		Info := Module.GetInfo()
		Log.Error.Printf(Info.Name, " > OnStart() failure")
		Log.Error.Fatalln(err)
		return nil
	}

	return Module
}
