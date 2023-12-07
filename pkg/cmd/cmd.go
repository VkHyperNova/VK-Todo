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
	print.ClearScreen()
	global.DB = database.LoadToDoDatabase()
	database.LoadGoals()
	database.LoadTaskNames()

	print.PrintTodo(global.Version)

	print.PrintGoals()
	print.PrintGoalsCommands()

	print.PrintTasks()

	print.PrintTaskCommands()
	print.PrintGreen("\ndatabase")
	print.AddBrackets(strconv.Itoa(len(global.DB)))

	print.PrintCyan("=> ")

	var cmd string = ""
	var id int = 0

	fmt.Scanln(&cmd, &id)
	for {
		switch cmd {
		case "add":
			database.CreateTask()
			Cmd()
		case "complete":
			database.CompleteTask(id)
			Cmd()
		case "update":
			database.UpdateTask(id)
			Cmd()
		case "delete":
			database.DeleteTask(id)
			Cmd()
		case "daygoal":
			database.CreateDayGoal()
			Cmd()
		case "weekgoal":
			database.CreateWeekGoal()
			Cmd()
		case "monthgoal":
			database.CreateMonthGoal()
			Cmd()
		case "yeargoal":
			database.CreateYearGoal()
			Cmd()
		case "lifetimegoal":
			database.CreateLifeTimeGoal()
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
