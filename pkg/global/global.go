package global

var DB []Todolist

const Version = "1"

type Todolist struct {
	ID       int    `json:"id"`
	TASK     string `json:"task"`
	COMPLETE bool   `json:"complete"`
	DATE     string `json:"date"`
}
