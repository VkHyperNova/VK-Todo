package cmd

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"vk-todo/pkg/database"
	"vk-todo/pkg/print"
	"vk-todo/pkg/util"
)

func Cmd() {
	print.ClearScreen()
	DB := database.LoadToDo()
	database.LoadGoals()

	print.PrintGray("============================================\n")
	print.PrintGray("============== VK TODO v1.1 ================\n")
	print.PrintGray("============================================\n")

	PrintGoals()
	PrintTasks(DB)
	PrintCommands([]string{"add", "complete", "update", "delete", "history", "q"})

	print.PrintGreen("\ndb")
	util.AddBrackets(strconv.Itoa(len(DB)))
	print.PrintCyan("=> ")

	var cmd string = ""
	var id int = 0

	fmt.Scanln(&cmd, &id)
	for {
		switch cmd {
		case "add", "a":
			CreateTask(DB)
			Cmd()
		case "complete", "c":
			CompleteTask(id, DB)
			Cmd()
		case "update", "u":
			UpdateTask(id, DB)
			Cmd()
		case "delete", "d":
			DeleteTask(id, DB)
			Cmd()
		case "history", "h":
			ShowHistory(DB)
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
	database.DayGoal = util.GetInput("Day Goal: ")
	Goals := database.NewGoals()
	database.SaveGoals(Goals)
}

func CreateWeekGoal() {
	database.WeekGoal = util.GetInput("Week Goal: ")
	Goals := database.NewGoals()
	database.SaveGoals(Goals)
}

func CreateMonthGoal() {
	database.MonthGoal = util.GetInput("Month Goal: ")
	Goals := database.NewGoals()
	database.SaveGoals(Goals)
}

func CreateYearGoal() {
	database.YearGoal = util.GetInput("Year Goal: ")
	Goals := database.NewGoals()
	database.SaveGoals(Goals)
}

func CreateLifeTimeGoal() {
	database.LifeTimeGoal = util.GetInput("Life Time Goal: ")
	Goals := database.NewGoals()
	database.SaveGoals(Goals)
}

func CreateTask(DB []database.Todolist) {

	Name := strings.ToUpper(util.GetInput("Name"))
	if Name == "Q" {
		Cmd()
	}

	Task := util.GetInput("Task")
	if Task == "Q" {
		Cmd()
	}

	NewTask := database.NewTask(Name, Task, DB)
	DB = append(DB, NewTask)
	database.SaveToDo(DB)
}

func CompleteTask(id int, DB []database.Todolist) {

	index := database.SearchIndexByID(id, DB)

	confirm := false

	if index == -1 {
		print.PrintRed("\nIndex out of range!\n")
	} else {
		PrintOneTask(index, DB)
		confirm = util.Confirm()
	}

	if confirm {
		DB[index].COMPLETE = true
		database.SaveToDo(DB)
		print.PrintGreen("Task done!\n\n")
	} else {
		print.PrintGreen("Returning../\n\n")
	}
}

func UpdateTask(id int, DB []database.Todolist) {
	index := database.SearchIndexByID(id, DB)
	confirm := false

	if index == -1 {
		print.PrintRed("\nIndex out of range!\n")
	} else {
		PrintOneTask(index, DB)
		confirm = util.Confirm()
	}

	if confirm {
		Task := util.GetInput("Updated Task: ")
		DB[index].TASK = Task
		database.SaveToDo(DB)
		print.PrintGreen("Task Updated!\n\n")
	} else {
		print.PrintGreen("Returning../\n\n")
	}
}

func DeleteTask(id int, DB []database.Todolist) {
	confirm := false

	index := database.SearchIndexByID(id, DB)

	if index == -1 {
		print.PrintRed("\nIndex out of range!\n")
	} else {
		PrintOneTask(index, DB)
		confirm = util.Confirm()
	}

	if confirm {
		DB = append(DB[:index], DB[index+1:]...)
		database.SaveToDo(DB)
		print.PrintGreen("Task deleted!\n\n")
	} else {
		print.PrintGreen("Returning../\n\n")
	}
}

func PrintCommands(commands []string) {
	print.PrintCyan("\n\n<< ")
	for _, value := range commands {
		print.PrintCyan("[")
		print.PrintYellow(value)
		print.PrintCyan("] ")
	}
	print.PrintCyan(" >>\n")
}

func PrintTasks(DB []database.Todolist) {
	print.PrintCyan("\n\n================= Projects " + strconv.Itoa(len(DB)) + " =================\n")

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
	print.PrintGreen("Completed: ")
	print.PrintGreen(strconv.Itoa(CompletedCount))
	print.PrintCyan("/")
	print.PrintRed(strconv.Itoa(UnCompletedCount))
	print.PrintCyan("\n--------------------------------------------\n\n")


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
		print.PrintGreen(value + " ")
	}

	print.PrintCyan("\n\n=========== Unfinished Tasks ===============\n")
	for _, project := range Projects {

		print.PrintYellow("\n" + project + "\n")

		count := 0

		for _, value := range DB {
			if project == value.NAME && !value.COMPLETE {
				print.PrintCyan(" [")
				print.PrintYellow(strconv.Itoa(value.ID))
				print.PrintCyan("] ")
				print.PrintCyan(value.TASK + "\n")
			}

			if project == value.NAME && value.COMPLETE {
				count += 1
			}
		}

		print.PrintCyan("--------------------(" + strconv.Itoa(count) + ")----------------------\n")
	}
}

func PrintOneTask(index int, DB []database.Todolist) {
	print.PrintCyan("\n\n================= Tasks ====================\n")
	print.PrintGreen("\n" + DB[index].NAME + "\n")
	print.PrintCyan(" [")
	print.PrintYellow(strconv.Itoa(DB[index].ID))
	print.PrintCyan("] ")
	print.PrintCyan(DB[index].TASK + " ")
	print.PrintCyan("\n\n--------------------------------------------\n")
}

func PrintGoals() {
	print.PrintCyan("\n\n================= Goals ====================\n")

	print.PrintCyan("Day -> ")
	print.PrintPurple(database.DayGoal + "\n")
	print.PrintCyan("--------------------------------------------\n")

	print.PrintCyan("Week -> ")
	print.PrintPurple(database.WeekGoal + "\n")
	print.PrintCyan("--------------------------------------------\n")

	print.PrintCyan("Month -> ")
	print.PrintPurple(database.MonthGoal + "\n")
	print.PrintCyan("--------------------------------------------\n")

	print.PrintCyan("Year -> ")
	print.PrintPurple(database.YearGoal + "\n")
	print.PrintCyan("--------------------------------------------\n")

	print.PrintCyan("Lifetime -> ")
	print.PrintPurple(database.LifeTimeGoal + "\n")
}

func ShowHistory(DB []database.Todolist) {
	print.PrintGreen("\n\n============ Completed Tasks ===============\n")

	// Get Names
	var Names []string
	for _, value := range DB {
		if !util.Contains(Names, value.NAME) {
			Names = append(Names, value.NAME)
		}
	}

	for _, name := range Names {

		print.PrintGreen("\n" + name + "\n")
		count := 0

		for _, value := range DB {
			if name == value.NAME && value.COMPLETE {
				print.PrintGreen(" [")
				print.PrintGray(strconv.Itoa(value.ID))
				print.PrintGreen("] ")
				print.PrintGreen(value.TASK + "\n")
			}

			if name == value.NAME && value.COMPLETE {
				count += 1
			}
		}

		print.PrintGreen("--------------------(" + strconv.Itoa(count) + ")----------------------\n")
	}
	print.PrintPurple("\n << Press ENTER to continue! >>")
	fmt.Scanln() // Press enter to continue
}
