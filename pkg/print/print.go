package print

import (
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
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

func PrintProgramName(Version string) {
	PrintCyan("\n<< VK-TODO " + Version + " >>\n")
}

func PrintCommands(commands []string) {
	PrintCyan("\n\n<< ")
	for _, value := range commands {
		PrintCyan("[")
		PrintYellow(value)
		PrintCyan("] ")
	}
	PrintCyan(" >>\n")
}

func AddBrackets(name string) {
	PrintCyan("[")
	PrintYellow(name)
	PrintCyan("] ")
}

func PrintTasks() {
	PrintCyan("\n\n================= Tasks ====================\n")

	for name, count := range global.AllTaskNamesMap {

		PrintGreen("\n" + strings.ToUpper(name) + " (" + strconv.Itoa(count) + ")\n")

		for _, value := range global.DB {
			if name == value.NAME && !value.COMPLETE {
				PrintCyan(" [")
				PrintYellow(strconv.Itoa(value.ID))
				PrintCyan("] ")
				PrintCyan(value.TASK + "\n")
			}
		}

		PrintCyan("--------------------------------------------\n")
	}
}

func PrintTask(index int) {
	PrintYellow(strconv.Itoa(global.DB[index].ID) + ". ")
	PrintYellow(global.DB[index].TASK + " ")
}

func PrintGoals() {
	PrintCyan("\n\n================= Goals ====================\n")
	
	PrintCyan("Day -> ")
	PrintPurple(global.DayGoal + "\n")
	PrintCyan("--------------------------------------------\n")

	PrintCyan("Week -> ")
	PrintPurple(global.WeekGoal + "\n")
	PrintCyan("--------------------------------------------\n")

	PrintCyan("Month -> ")
	PrintPurple(global.MonthGoal + "\n")
	PrintCyan("--------------------------------------------\n")

	PrintCyan("Year -> ")
	PrintPurple(global.YearGoal + "\n")
	PrintCyan("--------------------------------------------\n")

	PrintCyan("Lifetime -> ")
	PrintPurple(global.LifeTimeGoal + "\n")
}
