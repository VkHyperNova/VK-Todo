package main

import (
	"vk-todo/pkg/print"
	"vk-todo/pkg/dir"
	"vk-todo/pkg/cmd"
)



/*
1. Write test first
2. Pass the test
3. Refactor

->	Database to save todos
-> task: Add, Done, Update, Delete
*/

func main() {
	print.ClearScreen()
	dir.ValidateRequiredFiles()
	cmd.Cmd()
}