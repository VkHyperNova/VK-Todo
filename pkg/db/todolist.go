package db

import (
	"encoding/json"
	"vk-todo/pkg/util"
)
type Todolist struct {
	ID       int    `json:"id"`
	NAME     string `json:"name"`
	TASK     string `json:"task"`
	COMPLETE bool   `json:"complete"`
	DATE     string `json:"date"`
}

func SaveToDo(DB []Todolist) {
	DatabaseAsByte := util.InterfaceToByte(DB)
	util.WriteDataToFile("./database/todolist.json", DatabaseAsByte)
}

func LoadToDo() []Todolist {
	file := util.ReadFile("./database/todolist.json")

	Todolist := []Todolist{}

	err := json.Unmarshal(file, &Todolist)
	util.HandleError(err)

	return Todolist
}

func NewTask(name string, task string, DB []Todolist) Todolist {

	return Todolist{
		ID:       FindUniqueID(DB),
		NAME:     name,
		TASK:     task,
		COMPLETE: false,
		DATE:     util.GetFormattedDate(),
	}
}

func FindUniqueID(DB []Todolist) int {

	if len(DB) == 0 {
		return 1
	}

	return DB[len(DB)-1].ID + 1
}

func SearchIndexByID(id int, DB []Todolist) int {

	index := -1

	for key, value := range DB {
		if id == value.ID {
			index = key
		}
	}

	return index
}

func ValidateRequiredFiles() {
	if !util.DoesDirectoryExist("./database/todolist.json") {
		util.CreateDirectory("database")
		util.WriteDataToFile("./database/todolist.json", []byte("[]"))
	} 
}