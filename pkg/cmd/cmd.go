package cmd

import (
	"fmt"
	"os"
	"strconv"
	"vk-todo/pkg/database"
	"vk-todo/pkg/global"
	"vk-todo/pkg/print"
	"vk-todo/pkg/util"
)

func Cmd() {
	print.ClearScreen()
	global.DB = database.LoadToDo()
	database.LoadGoals()
	GetTaskNames()

	print.PrintGray("============================================\n")
	print.PrintGray("============== VK TODO v1.1 ================\n")
	print.PrintGray("============================================\n")

	print.PrintGoals()
	print.PrintTasks()
	print.PrintCommands([]string{"add", "complete", "update", "delete", "q"})

	print.PrintGreen("\ndb")
	print.AddBrackets(strconv.Itoa(len(global.DB)))
	print.PrintCyan("=> ")

	var cmd string = ""
	var id int = 0

	fmt.Scanln(&cmd, &id)
	for {
		switch cmd {
		case "add", "a":
			CreateTask()
			Cmd()
		case "complete", "c":
			CompleteTask(id)
			Cmd()
		case "update", "u":
			UpdateTask(id)
			Cmd()
		case "delete", "d":
			DeleteTask(id)
			Cmd()
		case "daygoal", "day":
			CreateDayGoal()
			Cmd()
		case "weekgoal", "week":
			CreateWeekGoal()
			Cmd()
		case "monthgoal", "month":
			CreateMonthGoal()
			Cmd()
		case "yeargoal", "year":
			CreateYearGoal()
			Cmd()
		case "lifetimegoal", "life":
			CreateLifeTimeGoal()
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

func CreateDayGoal() {
	global.DayGoal = util.GetInput("Day Goal: ")
	Goals := util.CompileGoals()
	database.SaveGoals(Goals)
}

func CreateWeekGoal() {
	global.WeekGoal = util.GetInput("Week Goal: ")
	Goals := util.CompileGoals()
	database.SaveGoals(Goals)
}

func CreateMonthGoal() {
	global.MonthGoal = util.GetInput("Month Goal: ")
	Goals := util.CompileGoals()
	database.SaveGoals(Goals)
}

func CreateYearGoal() {
	global.YearGoal = util.GetInput("Year Goal: ")
	Goals := util.CompileGoals()
	database.SaveGoals(Goals)
}

func CreateLifeTimeGoal() {
	global.LifeTimeGoal = util.GetInput("Life Time Goal: ")
	Goals := util.CompileGoals()
	database.SaveGoals(Goals)
}

func CreateTask() {
	Name := util.GetInput("Name: ")
	Task := util.GetInput("Task: ")
	NewTask := util.CompileTask(Name, Task)
	global.DB = append(global.DB, NewTask)
	database.SaveToDo()
}

func CompleteTask(id int) {
	index := util.SearchIndexByID(id)

	confirm := false

	if index == -1 {
		print.PrintRed("\nIndex out of range!\n")
	} else {
		print.PrintTask(index)
		confirm = util.Confirm()
	}

	if confirm {
		global.DB[index].COMPLETE = true
		database.SaveToDo()
		print.PrintGreen("Task done!\n\n")
	} else {
		print.PrintGreen("Returning../\n\n")
	}

}

func UpdateTask(id int) {
	index := util.SearchIndexByID(id)
	confirm := false

	if index == -1 {
		print.PrintRed("\nIndex out of range!\n")
	} else {
		print.PrintTask(index)
		confirm = util.Confirm()
	}

	if confirm {
		Task := util.GetInput("Updated Task: ")
		global.DB[index].TASK = Task
		database.SaveToDo()
		print.PrintGreen("Task Updated!\n\n")
	} else {
		print.PrintGreen("Returning../\n\n")
	}

}

func DeleteTask(id int) {
	confirm := false

	index := util.SearchIndexByID(id)

	if index == -1 {
		print.PrintRed("\nIndex out of range!\n")
	} else {
		print.PrintTask(index)
		confirm = util.Confirm()
	}

	if confirm {
		global.DB = append(global.DB[:index], global.DB[index+1:]...)
		database.SaveToDo()
		print.PrintGreen("Task deleted!\n\n")
	} else {
		print.PrintGreen("Returning../\n\n")
	}

}

func GetTaskNames() {
	// Get names
	for _, value := range global.DB {
		if !util.Contains(global.AllTaskNames, value.NAME) {
			global.AllTaskNames = append(global.AllTaskNames, value.NAME)
		}
	}

	// Count
	myMap := make(map[string]int)
	for _, name := range global.AllTaskNames {
		for _, value := range global.DB {
			if name == value.NAME && value.COMPLETE {
				myMap[value.NAME] += 1
			}
			if name == value.NAME && !value.COMPLETE {
				myMap[value.NAME] += 0
			}
		}
	}

	global.AllTaskNamesMap = myMap
}
