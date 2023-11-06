package main

import (
	"vk-todo/pkg/print"
	"vk-todo/pkg/dir"
	"vk-todo/pkg/cmd"
)

func main() {
	print.ClearScreen()
	dir.ValidateRequiredFiles()
	cmd.Cmd()
}