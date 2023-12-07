package global

var DB []Todolist



var AllTaskNames []string
var AllTaskNamesMap map[string]int

// Goals

var DayGoal string
var WeekGoal string
var MonthGoal string
var YearGoal string
var LifeTimeGoal string

type Todolist struct {
	ID       int    `json:"id"`
	NAME     string `json:"name"`
	TASK     string `json:"task"`
	COMPLETE bool   `json:"complete"`
	DATE     string `json:"date"`
}

type Goals struct {
	DAYGOAL      string `json:"daygoal"`
	WEEKGOAL     string `json:"weekgoal"`
	MONTHGOAL    string `json:"monthgoal"`
	YEARGOAL     string `json:"yeargoal"`
	LIFETIMEGOAL string `json:"lifetimegoal"`
}
