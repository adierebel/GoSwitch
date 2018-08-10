package main

import (
    "fmt"
    "os"
	"path"
	"runtime"
	"io/ioutil"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	// Define variable
	env_name	= kingpin.Flag("name", "Name of new GOPATH workdir.").Short('n').Default("GOPATH").String()
	target_path	= kingpin.Flag("path", "Target path for script files.").Short('p').Default("scripts").String()
	filemode	= kingpin.Flag("filemode", "Filemode for creating file and directory.").Short('m').Default("0644").Int()
)

func main() {
	// Parse arguments
	kingpin.Version("1.0.0")
	kingpin.Parse()
	apply()
}

func apply() {
	// Get current dir
	go_path, err := os.Getwd()
    check(err)
	
	// Print info
	fmt.Println(" * Create '" +*target_path+ "' directory")

	// Create new directory
	os.Mkdir(path.Join(go_path, *target_path), os.FileMode(*filemode))
	
	// Check OS version
	if runtime.GOOS == "windows" {
		// File script files
		write_script_windows(go_path)

		// Print information
		fmt.Println(" * Create script files for Windows")
		fmt.Println(" * How to use:")
		fmt.Println("   - Activate -> call " +*target_path+ "/activate")
		fmt.Println("   - Deactivate -> deactivate")
		fmt.Println(" * Done!")
	} else if runtime.GOOS == "linux" {
		
	} else if runtime.GOOS == "darwin" {
		
	} else {
		// Unknown OS platform, exit(1)
		fmt.Println("Unknown operating system platform")
		os.Exit(1)
	}
}

func check(e error) {
	// Chek error
    if e != nil {
        panic(e)
    }
}

func file_writer(filepath string, body string) {
	// Write file
	data := []byte(body)
    err := ioutil.WriteFile(filepath, data, os.FileMode(*filemode))
    check(err)
}

func write_script_windows(go_path string) {
	// Script activate for windows
	text_activate := `@echo off
if not defined PROMPT (
	set "PROMPT=$P$G"
)
if defined _OLD_CMD_PROMPT (
	set "PROMPT=%_OLD_CMD_PROMPT%"
)
set "_OLD_CMD_PROMPT=%PROMPT%"
if defined _OLD_CMD_PATH (
	set "PATH=%_OLD_CMD_PATH%"
) else (
	set "_OLD_CMD_PATH=%PATH%"
)
if defined _OLD_CMD_GOPATH (
	set "GOPATH=%_OLD_CMD_GOPATH%"
) else (
	set "_OLD_CMD_GOPATH=%GOPATH%"
)
set "PROMPT=(` +*env_name+ `) %PROMPT%"
set "GOPATH=` +go_path+ `"
set "PATH=%GOPATH%;%PATH%"
set "PATH=%GOPATH%\scripts;%PATH%"
:END`

	// Script deactivate for windows
	text_deactivate := `@echo off
if defined _OLD_CMD_PROMPT (
	set "PROMPT=%_OLD_CMD_PROMPT%"
)
set _OLD_CMD_PROMPT=
if defined _OLD_CMD_PATH (
	set "PATH=%_OLD_CMD_PATH%"
)
set _OLD_CMD_PATH=
if defined _OLD_CMD_GOPATH (
	set "GOPATH=%_OLD_CMD_GOPATH%"
)
set "_OLD_CMD_GOPATH=
:END`

	// Write activate.bat on windows
	file_writer(path.Join(go_path, *target_path, "activate.bat"), text_activate)
	// Write deactivate.bat on windows
	file_writer(path.Join(go_path, *target_path, "deactivate.bat"), text_deactivate)
}