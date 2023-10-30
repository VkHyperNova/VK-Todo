package cmd

import (
	"fmt"
	"os"
	"vk-todo/pkg/global"
	"vk-todo/pkg/print"
	"vk-todo/pkg/database"
)

func Cmd() {

	global.DB = database.LoadDatabase()

	print.PrintTasks()

	print.PrintTodo(len(global.DB), global.Version)

	print.PrintCyan("=> ")

	var cmd string = ""
	var id int = 0

	fmt.Scanln(&cmd, &id)
	for {
		switch cmd {
		case "add":
			database.Create()
			Cmd()
		case "done":
			database.Update(id)
			Cmd()
		case "delete":
			database.Delete(id)
			Cmd()
		case "q":
			os.Exit(0)
		default:
			Cmd()
		}
	}
}
