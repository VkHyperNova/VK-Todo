package cmd

import (
	"fmt"
	"os"
	"strconv"
	"vk-todo/pkg/database"
	"vk-todo/pkg/global"
	"vk-todo/pkg/print"
)

func Cmd() {

	global.DB = database.LoadToDoDatabase()
	database.LoadTaskNames()
	database.CountCompletedTasks()

	print.PrintTodo(global.Version)

	print.PrintGoals()
	
	print.PrintTasks()

	print.PrintCompletedTasksCount()
	print.PrintTasksCount()

	print.PrintCommands()
	print.PrintGreen("\ndatabase")
	print.AddBrackets(strconv.Itoa(len(global.DB)))

	print.PrintCyan("=> ")

	var cmd string = ""
	var id int = 0

	fmt.Scanln(&cmd, &id)
	for {
		switch cmd {
		case "add":
			database.Create()
			Cmd()
		case "complete":
			database.Complete(id)
			Cmd()
		case "update":
			database.Update(id)
			Cmd()
		case "delete":
			database.Delete(id)
			Cmd()
		case "q":
			print.ClearScreen()
			os.Exit(0)
		default:
			print.ClearScreen()
			Cmd()
		}
	}
}
