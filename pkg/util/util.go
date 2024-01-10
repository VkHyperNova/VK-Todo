package util

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)



func GetInput(inputType string) string {
start:
	PrintCyan(inputType + ": ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	input = strings.TrimSpace(input)
	if input == "" {
		PrintRed("Please enter a " + inputType + "\n")
		goto start
	}

	return input
}

func GetFormattedDate() string {
	return time.Now().Format("02.01.2006")
}

func InterfaceToByte(data interface{}) []byte {
	dataBytes, err := json.MarshalIndent(data, "", "  ")
	HandleError(err)

	return dataBytes
}

func Confirm() bool {

	PrintYellow("\n\nThis One?: ")

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
