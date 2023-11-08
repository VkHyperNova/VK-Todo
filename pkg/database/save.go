package database

import (
	"vk-todo/pkg/dir"
	"vk-todo/pkg/global"
	"vk-todo/pkg/util"
)

func SaveToDoDatabase() {
	DatabaseAsByte := util.InterfaceToByte(global.DB)
	dir.WriteDataToFile("./database/todolist.json", DatabaseAsByte)
}

func SaveToGoalsDatabase(goals global.Goals) {
	DatabaseAsByte := util.InterfaceToByte(goals)
	dir.WriteDataToFile("./database/goals.json", DatabaseAsByte)
}

