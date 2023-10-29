package database

import (
	"vk-todo/pkg/global"
	"vk-todo/pkg/util"
	"vk-todo/pkg/dir"
)


func Create() {
	NewTaskString := util.Prompt("Task: ")
	NewTask := util.CompileTask(NewTaskString)
	global.DB = append(global.DB, NewTask)
	DatabaseAsByte := util.InterfaceToByte(global.DB) 
	dir.WriteDataToFile("./database/todolist.json", DatabaseAsByte)	
}

func Update(id int) {

}

func Delete(id int) {

}