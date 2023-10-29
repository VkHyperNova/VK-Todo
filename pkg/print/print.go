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
	AddBrackets("update")
	AddBrackets("delete")
	AddBrackets("q")
	PrintGray("\n\n")
	

	PrintGreen("\n\ndatabase")
	AddBrackets(strconv.Itoa(db_size))

}

func AddBrackets(name string) {
	PrintCyan("[")
	PrintYellow(name)
	PrintCyan("] ")
}

func PrintTasks() {
	for _, value := range global.DB {
		PrintCyan(strconv.Itoa(value.ID) + ". ")
		PrintCyan(value.DATE + " ")
		PrintCyan(value.TASK)
		
		
		PrintCyan("\n")

	}
}
