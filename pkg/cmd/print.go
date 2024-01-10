package cmd

import (
	"fmt"
	"strconv"
	"vk-todo/pkg/db"
	"vk-todo/pkg/util"
)

func PrintVKTODO(Version string, DATABASE []db.Todolist) {
	util.PrintGray("============================================\n")
	util.PrintGray("============== VK TODO v1.1 ================\n")
	util.PrintGray("============================================\n")

	PrintTasks(DATABASE)

	util.PrintCyan("\n\n<< ")
	commands := [6]string{"add", "complete", "update", "delete", "history", "q"}
	for _, value := range commands {
		PrintBrackets(value)
	}
	util.PrintCyan(" >>\n")

	util.PrintGreen("\ndb")
	PrintBrackets(strconv.Itoa(len(DATABASE)))
	util.PrintCyan("=> ")
}

func PrintTasks(DB []db.Todolist) {
	util.PrintCyan("\n\nProjects ")
	util.PrintGray(strconv.Itoa(len(DB)))

	// Get Projects
	var Projects []string

	// Count
	CompletedCount := 0

	for _, value := range DB {
		if value.COMPLETE {
			CompletedCount++
		}
	}
	UnCompletedCount := len(DB) - CompletedCount
	util.PrintCyan("\tCompleted: ")
	util.PrintGreen(strconv.Itoa(CompletedCount))
	util.PrintCyan("\tUnFinished: ")
	util.PrintRed(strconv.Itoa(UnCompletedCount) + "\n\n")

	// Uncompleted projects
	for _, value := range DB {
		if !util.Contains(Projects, value.NAME) && !value.COMPLETE {
			Projects = append(Projects, value.NAME)
		}
	}

	// Getting all project names
	var AllProjects []string
	for _, value := range DB {
		if !util.Contains(AllProjects, value.NAME) {
			AllProjects = append(AllProjects, value.NAME)
		}
	}

	// Print all project names
	for _, value := range AllProjects {
		util.PrintGreen(value + "; ")
	}

	util.PrintCyan("\n\n=========== Unfinished Tasks ===============\n")
	for _, project := range Projects {

		util.PrintYellow("\n" + project + "\n")

		count := 0

		for _, value := range DB {
			if project == value.NAME && !value.COMPLETE {
				PrintBrackets(strconv.Itoa(value.ID))
				util.PrintCyan(value.TASK + "\n")
			}

			if project == value.NAME && value.COMPLETE {
				count += 1
			}
		}

		util.PrintCyan("--------------------(" + strconv.Itoa(count) + ")----------------------\n")
	}
}

func PrintOneTask(index int, DB []db.Todolist) {
	util.PrintCyan("\n\n================= Tasks ====================\n")
	util.PrintGreen("\n" + DB[index].NAME + "\n")
	PrintBrackets(strconv.Itoa(DB[index].ID))
	util.PrintCyan(DB[index].TASK + " ")
	util.PrintCyan("\n\n--------------------------------------------\n")
}

func PrintHistory(DB []db.Todolist) {
	util.PrintGreen("\n\n============ Completed Tasks ===============\n")

	// Get Names
	var Names []string
	for _, value := range DB {
		if !util.Contains(Names, value.NAME) {
			Names = append(Names, value.NAME)
		}
	}

	for _, name := range Names {

		util.PrintGreen("\n" + name + "\n")
		count := 0

		for _, value := range DB {
			if name == value.NAME && value.COMPLETE {
				util.PrintGreen(" [")
				util.PrintGray(strconv.Itoa(value.ID))
				util.PrintGreen("] ")
				util.PrintGreen(value.TASK + "\n")
			}

			if name == value.NAME && value.COMPLETE {
				count += 1
			}
		}

		util.PrintGreen("--------------------(" + strconv.Itoa(count) + ")----------------------\n")
	}
	util.PrintPurple("\n << Press ENTER to continue! >>")
	fmt.Scanln() // Press enter to continue
}

func PrintBrackets(name string) {
	util.PrintCyan("[")
	util.PrintYellow(name)
	util.PrintCyan("] ")
}
