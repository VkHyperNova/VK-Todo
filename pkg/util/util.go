package util

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
	"vk-todo/pkg/print"
)

func AddBrackets(name string) {
	print.PrintCyan("[")
	print.PrintYellow(name)
	print.PrintCyan("] ")
}

func GetInput(inputName string) string {
	print.PrintCyan(inputName)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	NewTaskString := scanner.Text()
	NewTaskString = strings.TrimSpace(NewTaskString)

	return NewTaskString
}

func Prompt(Question string) string {

	print.PrintCyan(Question)

	var user_input_string string
	fmt.Scanln(&user_input_string)

	return user_input_string
}

func GetFormattedDate() string {
	return time.Now().Format("02.01.2006")
}

func InterfaceToByte(data interface{}) []byte {
	dataBytes, err := json.MarshalIndent(data, "", "  ")
	print.HandleError(err)

	return dataBytes
}

func Confirm() bool {

	user_input := Prompt("\n\nThis One?: ")

	if user_input == "n" || user_input == "no" {
		return false
	}
	return true
}

func Contains(arr []string, name string) bool {
	for _, n := range arr {
		if n == name {
			return true
		}
	}
	return false
}


