package util

import (
	"encoding/json"
	"fmt"
	"time"
	"vk-todo/pkg/global"
	"vk-todo/pkg/print"
)

func Prompt(Question string) string {

	print.PrintCyan(Question)

	var user_input_string string
	fmt.Scanln(&user_input_string)

	return user_input_string
}

func GetFormattedDate() string {
	return time.Now().Format("15:04 (02.01.2006)")
}

func InterfaceToByte(data interface{}) []byte {
	dataBytes, err := json.MarshalIndent(data, "", "  ")
	print.HandleError(err)

	return dataBytes
}

func FindUniqueID() int {

	if len(global.DB) == 0 {
		return 1
	}

	return global.DB[len(global.DB)-1].ID + 1
}

func CompileTask(task string) global.Todolist {

	return global.Todolist{
		ID:       FindUniqueID(),
		TASK:     task,
		COMPLETE: false,
		DATE:     GetFormattedDate(),
	}
}

func GetTaskArray(body []byte) []global.Todolist {

	TodolistStruct := []global.Todolist{}

	err := json.Unmarshal(body, &TodolistStruct)
	print.HandleError(err)

	return TodolistStruct
}