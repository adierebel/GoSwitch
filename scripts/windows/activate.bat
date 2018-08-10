@echo off
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
set "PROMPT=(GOPATH) %PROMPT%"
set "GOPATH=C:\goo"
set "PATH=%GOPATH%;%PATH%"
set "PATH=%GOPATH%\scripts;%PATH%"
:END