package database

import (
	"bufio"
	"os"
	"strings"
	"vk-todo/pkg/global"
	"vk-todo/pkg/print"
	"vk-todo/pkg/util"
)

func Create() {
	print.PrintCyan("Task: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	NewTaskString := scanner.Text()
	NewTaskString = strings.TrimSpace(NewTaskString)

	NewTask := util.CompileTask(NewTaskString)
	global.DB = append(global.DB, NewTask)
	SaveDatabase()
	print.ClearScreen()
}

func Complete(id int) {
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
		SaveDatabase()
		print.PrintGreen("Task done!\n\n")
	} else {
		print.PrintGreen("Returning../\n\n")
	}

	print.ClearScreen()
}

func Update(id int) {
	index := util.SearchIndexByID(id)
	confirm := false

	if index == -1 {
		print.PrintRed("\nIndex out of range!\n")
	} else {
		print.PrintTask(index)
		confirm = util.Confirm()
	}

	if confirm {
		print.PrintCyan("Task: ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		UpdatedTaskString := scanner.Text()
		UpdatedTaskString = strings.TrimSpace(UpdatedTaskString)
		global.DB[index].TASK = UpdatedTaskString
		SaveDatabase()
		print.PrintGreen("Task Updated!\n\n")
	} else {
		print.PrintGreen("Returning../\n\n")
	}

	print.ClearScreen()

}

func Delete(id int) {
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
		SaveDatabase()
		print.PrintGreen("Task deleted!\n\n")
	} else {
		print.PrintGreen("Returning../\n\n")
	}

	print.ClearScreen()
}
