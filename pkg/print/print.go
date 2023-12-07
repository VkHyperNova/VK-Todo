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

func PrintTaskCommands() {
	PrintCyan("\n\n<< ")
	AddBrackets("add")
	AddBrackets("complete")
	AddBrackets("update")
	AddBrackets("delete")
	AddBrackets("q")
	PrintCyan(" >>\n")
}

func PrintGoalsCommands() {
	PrintCyan("\n\n<< ")
	AddBrackets("daygoal")
	AddBrackets("weekgoal")
	AddBrackets("monthgoal")
	AddBrackets("yeargoal")
	AddBrackets("lifetimegoal")
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
	// PrintYellow(global.DB[index].DATE + " ")
	PrintYellow(global.DB[index].TASK + " ")
	// PrintYellow(strconv.FormatBool(global.DB[index].COMPLETE))
}

func PrintGoals() {
	PrintCyan("Day Goal: ")
	PrintGreen(global.DayGoal + "\n")

	PrintCyan("Week Goal: ")
	PrintGreen(global.WeekGoal + "\n")

	PrintCyan("Month Goal: ")
	PrintGreen(global.MonthGoal + "\n")

	PrintCyan("Year Goal: ")
	PrintGreen(global.YearGoal + "\n")

	PrintCyan("Lifetime Goal: ")
	PrintGreen(global.LifeTimeGoal + "\n")
}
