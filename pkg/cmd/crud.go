package cmd

import (
	"strings"
	"vk-todo/pkg/db"
	"vk-todo/pkg/util"
)

func Add(DB []db.Todolist) {

	Name := strings.ToUpper(util.GetInput("Name"))
	if Name == "Q" {
		Cmd()
	}

	Task := util.GetInput("Task")
	if Task == "q" {
		Cmd()
	}

	NewTask := db.NewTask(Name, Task, DB)
	DB = append(DB, NewTask)
	db.SaveToDo(DB)
}

func Complete(id int, DB []db.Todolist) {

	index := db.SearchIndexByID(id, DB)

	if index == -1 {
		util.PrintRed("\nIndex out of range!\n")
	} else {
		DB[index].COMPLETE = true
		db.SaveToDo(DB)
	}
}

func Update(id int, DB []db.Todolist) {
	index := db.SearchIndexByID(id, DB)

	if index == -1 {
		util.PrintRed("\nIndex out of range!\n")
	} else {
		PrintOneTask(index, DB)
		Task := util.GetInput("Update: ")
		DB[index].TASK = Task
		db.SaveToDo(DB)
		util.PrintGreen("Task Updated!\n\n")
	}
}

func Delete(id int, DB []db.Todolist) {

	index := db.SearchIndexByID(id, DB)

	if index == -1 {
		util.PrintRed("\nIndex out of range!\n")
	} else {
		DB = append(DB[:index], DB[index+1:]...)
		db.SaveToDo(DB)
	}
}