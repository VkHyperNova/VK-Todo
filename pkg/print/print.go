package print

import (
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"time"
	"vk-todo/pkg/global"
)

func ClearScreen() {

	if runtime.GOOS == "linux" {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func PrintTodo(Version string) {
	PrintCyan("\n<< VK-TODO " + Version + " >>\n")	
}

func PrintCommands() {
	PrintCyan("\n\n<< ")
	AddBrackets("add")
	AddBrackets("complete")
	AddBrackets("update")
	AddBrackets("delete")
	AddBrackets("q")
	PrintCyan(" >>\n")
}

func AddBrackets(name string) {
	PrintCyan("[")
	PrintYellow(name)
	PrintCyan("] ")
}

func PrintTasks() {
	PrintGray("\n")
	PrintGray("Tasks\n")
	for _, value := range global.DB {
		if !value.COMPLETE{
			PrintCyan(strconv.Itoa(value.ID) + ". ")
			// PrintRed(value.DATE + " ")
			PrintYellow(value.TASK + " ")
			// PrintYellow(strconv.FormatBool(value.COMPLETE))
			PrintCyan("\n")
			global.TaskCount += 1
		} 
	}
}

func PrintCompletedTasks() {
	cMonth := time.Now().Month()
	
	PrintGray("\n")
	PrintGray("Completed Tasks")
	PrintGray("\n")
	for _, value := range global.DB {
		dbMonth, err := time.Parse("02.01.2006", value.DATE)
   		HandleError(err)
		if value.COMPLETE && cMonth == dbMonth.Month() {
			PrintCyan(strconv.Itoa(value.ID) + ". ")
			PrintGreen(value.TASK + " ")
			PrintCyan("\n")
			global.CompletedTasksCount += 1
		} 
	}
}

func PrintTask(index int) {
	PrintYellow(strconv.Itoa(global.DB[index].ID) + ". ")
	PrintYellow(global.DB[index].DATE + " ")
	PrintYellow(global.DB[index].TASK + " ")
	PrintYellow(strconv.FormatBool(global.DB[index].COMPLETE))
}

func PrintCompletedTasksCount() {
	PrintCyan("\n\nCompleted: ")
	PrintGreen(strconv.Itoa(global.CompletedTasksCount))
}

func PrintTasksCount() {
	PrintCyan("\nTasks: ")
	PrintYellow(strconv.Itoa(global.TaskCount))
}

func PrintGoals() {
	PrintCyan("Day Goal: ")
	PrintCyan("Week Goal: ")
	PrintCyan("Month Goal: ")
	PrintCyan("Year Goal: ")
	PrintCyan("Lifetime Goal: ")
}




