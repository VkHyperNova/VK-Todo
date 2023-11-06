package global

var DB []Todolist

const Version = "1"

var CompletedTasksCount = 0
var TaskCount = 0
var AllTaskNames []string
var AllTaskNamesMap map[string]int

type Todolist struct {
	ID       int    `json:"id"`
	NAME     string `json:"name"`
	TASK     string `json:"task"`
	COMPLETE bool   `json:"complete"`
	DATE     string `json:"date"`
}
