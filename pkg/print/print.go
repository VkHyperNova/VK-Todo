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

func PrintTodo(db_size int, Version string) {
	PrintCyan("\n<_________________ PWM v" + Version + " __________________>\n\n")
	AddBrackets("add")
	AddBrackets("done")
	AddBrackets("delete")
	AddBrackets("q")
	PrintGreen("\n\ndatabase")
	AddBrackets(strconv.Itoa(db_size))
}

func AddBrackets(name string) {
	PrintCyan("[")
	PrintYellow(name)
	PrintCyan("] ")
}

func PrintTasks() {
	PrintGray("\n\n")
	for _, value := range global.DB {
		if !value.COMPLETE{
			PrintRed(strconv.Itoa(value.ID) + ". ")
			PrintRed(value.DATE + " ")
			PrintGreen(value.TASK + " ")
			PrintYellow(strconv.FormatBool(value.COMPLETE))
			PrintCyan("\n")
		}
	}
}

func PrintTask(index int) {
	PrintYellow(strconv.Itoa(global.DB[index].ID) + ". ")
	PrintYellow(global.DB[index].DATE + " ")
	PrintYellow(global.DB[index].TASK + " ")
	PrintYellow(strconv.FormatBool(global.DB[index].COMPLETE))
}
