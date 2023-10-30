package database

import (
	"vk-todo/pkg/dir"
	"vk-todo/pkg/global"
	"vk-todo/pkg/util"
)

func SaveDatabase() {
	DatabaseAsByte := util.InterfaceToByte(global.DB)
	dir.WriteDataToFile("./database/todolist.json", DatabaseAsByte)
}