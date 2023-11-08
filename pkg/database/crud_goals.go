package database

import (
	"vk-todo/pkg/global"
	"vk-todo/pkg/util"
)

func CreateDayGoal() {
	global.DayGoal = util.GetInput("Day Goal: ")
	Goals := util.CompileGoals()
	SaveToGoalsDatabase(Goals)
}

func CreateWeekGoal() {
	global.WeekGoal = util.GetInput("Week Goal: ")
	Goals := util.CompileGoals()
	SaveToGoalsDatabase(Goals)
}

func CreateMonthGoal() {
	global.MonthGoal = util.GetInput("Month Goal: ")
	Goals := util.CompileGoals()
	SaveToGoalsDatabase(Goals)
}

func CreateYearGoal() {
	global.YearGoal = util.GetInput("Year Goal: ")
	Goals := util.CompileGoals()
	SaveToGoalsDatabase(Goals)
}

func CreateLifeTimeGoal() {
	global.LifeTimeGoal = util.GetInput("Life Time Goal: ")
	Goals := util.CompileGoals()
	SaveToGoalsDatabase(Goals)
}
