package print

import (
	"os"
	"os/exec"
	"runtime"
	"strconv"
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
	PrintCyan("\n\n<< Tasks >>\n")

	for name, count := range global.AllTaskNamesMap {

		PrintGreen("\n" + name + " (" + strconv.Itoa(count) + ")\n")

		for _, value := range global.DB {
			if name == value.NAME && !value.COMPLETE {
				AddBrackets(strconv.Itoa(value.ID))
				PrintCyan(value.TASK + "\n")
			}
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
