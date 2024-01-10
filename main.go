package main

import (
	"vk-todo/pkg/cmd"
	"vk-todo/pkg/util"
	"vk-todo/pkg/db"
)

func main() {
	util.ClearScreen()
	db.ValidateRequiredFiles()
	cmd.Cmd()
}