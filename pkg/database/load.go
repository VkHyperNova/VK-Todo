package database

import (
	"fmt"
	"vk-todo/pkg/dir"
	"vk-todo/pkg/global"
	"vk-todo/pkg/util"
)

func LoadToDoDatabase() []global.Todolist {
	file := dir.ReadFile("./database/todolist.json")
	data := util.GetTaskArray(file)

	return data
}

func LoadTaskNames() {

	for _, value := range global.DB {
		if !util.Contains(global.AllTaskNames, value.NAME) {
			global.AllTaskNames = append(global.AllTaskNames, value.NAME)
		}
	}
}

func CountCompletedTasks() {

	myMap := make(map[string]int)
	// Count all completed tasks
	for _, name := range global.AllTaskNames {
		for _, value := range global.DB {
			if name == value.NAME && value.COMPLETE {
				myMap[value.NAME] += 1
			} 
		}
	}

	// Count not completed tasks
	for _, name := range global.AllTaskNames {
		for _, value := range global.DB {
			if name == value.NAME && !value.COMPLETE {
				myMap[value.NAME] += 0
			} 
		}
	}
	fmt.Println(myMap)
	global.AllTaskNamesMap = myMap

}

