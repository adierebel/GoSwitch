# GoSwitch

Easily switch GOPATH working directory

## Install

### Build from source

- Install Go from https://golang.org/
- Setup GOPATH, tutorial: https://github.com/golang/go/wiki/SettingGOPATH
- Download GoSwitch source on Release tab.
- Get dependencies, `go get gopkg.in/alecthomas/kingpin.v2`
- Build, `go build goswitch.go`

### Add to PATH

Add goswitch executable to PATH variable, so you can access goswitch everywhere on cmd/terminal

#### Windows

- Copy `goswitch.exe` to `C:\Program Files\GoSwitch\goswitch.exe`
- Open `Control Panel` -> `System and Security` -> `System` -> `Advanced system settings` -> Click button `Environment Variables..`
- On tab User variables for %user%, edit `PATH` and add `;C:\Program Files\GoSwitch` on end of text
- Open new CMD, type `goswitch`