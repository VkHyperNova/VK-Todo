package database

import (
	"vk-todo/pkg/global"
	"vk-todo/pkg/dir"
	"vk-todo/pkg/util"
)

func LoadDatabase() []global.Todolist {
	file := dir.ReadFile("./database/todolist.json")
	data := util.GetTaskArray(file)

	return data
}