package Module

import (
	"os"
	"path/filepath"
	"plugin"
	"strings"

	IChatanium "antegral.net/chatanium/src/Runtime/Interface"
	"antegral.net/chatanium/src/Runtime/Log"
)

/**
 * @brief      Gets a module.
 * @param      FilePath  The file path
 * @return     The module.
 */
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

/**
 * @brief      Starts a module.
 * @param      Module  The module
 * @return     The module.
 */
func Start(Module IChatanium.Module) IChatanium.Module {
	// Starting a module is a 2-step process:
	// 1. OnInit() - This is where the module is initialized. This is where the module should do any initialization work.
	// 2. OnStart() - This is where the module is started. This is where the module should start any background processes.

	// Get the name of the module
	Name := Module.GetInfo().Name

	Log.Verbose.Printf("Starting module: %s", Name)

	// Trigger: OnInit() of the Module
	if err := Module.OnInit(); err != nil {
		Log.Error.Printf("%s > OnInit() failure", Name)
		Log.Error.Fatalln(err)
		return nil
	}

	// Trigger: OnStart() of the Module
	if err := Module.OnStart(); err != nil {
		Info := Module.GetInfo()
		Log.Error.Printf(Info.Name, " > OnStart() failure")
		Log.Error.Fatalln(err)
		return nil
	}

	return Module
}

/**
 * @brief      Searches for modules.
 * @param      FolderPath  The folder path
 * @return     All modules in the folder.
 */
func Search(FolderPath string) []IChatanium.Module {
	Log.Verbose.Printf("Searching for modules in: %s", FolderPath)

	// all modules in the folder
	var Modules []IChatanium.Module

	// get all entries in the folder
	Entries, err := os.ReadDir("./")
	if err != nil {
		Log.Error.Printf("ChataniumRuntime > Module.Search() failure")
		Log.Error.Fatalln(err)
	}

	for _, e := range Entries {
		// check if the file is a DLL or SO
		IsDLL := strings.HasSuffix(e.Name(), ".dll")
		IsSO := strings.HasSuffix(e.Name(), ".so")

		if !IsDLL && !IsSO {
			// not a valid extension
			Log.Verbose.Printf("Search() > Not a valid extension. Ignored. (%s)", e.Name())
			continue
		}

		// get the module
		Module := Get(filepath.Join(FolderPath, e.Name()))

		// add to the list of modules
		Modules = append(Modules, Module)
	}

	return Modules
}

/**
 * @brief      Searches for all modules in the given folder and starts them.
 * @param      FolderPath  The folder path
 * @return     All modules.
 */
func StartAllModules(FolderPath string) []IChatanium.Module {
	// get all modules
	Modules := Search(FolderPath)

	for _, e := range Modules {
		// start the module
		Start(e)
	}

	return Modules
}
