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

func GetInput(inputType string) string {
	start:
	print.PrintCyan(inputType + ": ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	input = strings.TrimSpace(input)
	if input == "" {
		print.PrintRed("Please enter a " + inputType + "\n")
		goto start
	}

	return input
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

	print.PrintYellow("\n\nThis One?: ")

	var user_input string
	fmt.Scanln(&user_input)

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


