@echo off
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
:END