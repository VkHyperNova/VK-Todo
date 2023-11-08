package database

import (
	"vk-todo/pkg/global"
	"vk-todo/pkg/print"
	"vk-todo/pkg/util"
)

func CreateTask() {
	Name := util.GetInput("Name: ")
	Task := util.GetInput("Task: ")
	NewTask := util.CompileTask(Name, Task)
	global.DB = append(global.DB, NewTask)
	SaveToDoDatabase()
	print.ClearScreen()
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
		SaveToDoDatabase()
		print.PrintGreen("Task done!\n\n")
	} else {
		print.PrintGreen("Returning../\n\n")
	}

	print.ClearScreen()
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
		SaveToDoDatabase()
		print.PrintGreen("Task Updated!\n\n")
	} else {
		print.PrintGreen("Returning../\n\n")
	}

	print.ClearScreen()

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
		SaveToDoDatabase()
		print.PrintGreen("Task deleted!\n\n")
	} else {
		print.PrintGreen("Returning../\n\n")
	}

	print.ClearScreen()
}
