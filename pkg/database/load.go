package database

import (
	"vk-todo/pkg/dir"
	"vk-todo/pkg/global"
	"vk-todo/pkg/util"
)

func LoadToDoDatabase() []global.Todolist {
	file := dir.ReadFile("./database/todolist.json")
	data := util.GetTaskArray(file)

	return data
}

func LoadGoals() {
	file := dir.ReadFile("./database/goals.json")
	data := util.GetGoalsJson(file)
	global.DayGoal = data.DAYGOAL
	global.WeekGoal = data.WEEKGOAL
	global.MonthGoal = data.MONTHGOAL
	global.YearGoal = data.YEARGOAL
	global.LifeTimeGoal = data.LIFETIMEGOAL
}

func LoadTaskNames() {
	// Get names
	for _, value := range global.DB {
		if !util.Contains(global.AllTaskNames, value.NAME) {
			global.AllTaskNames = append(global.AllTaskNames, value.NAME)
		}
	}

	// Count
	myMap := make(map[string]int)
	for _, name := range global.AllTaskNames {
		for _, value := range global.DB {
			if name == value.NAME && value.COMPLETE {
				myMap[value.NAME] += 1
			} 
			if name == value.NAME && !value.COMPLETE {
				myMap[value.NAME] += 0
			} 
		}
	}

	global.AllTaskNamesMap = myMap
}


