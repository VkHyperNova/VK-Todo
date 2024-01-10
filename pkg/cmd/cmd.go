package cmd

import (
	"fmt"
	"os"
	"vk-todo/pkg/db"
	"vk-todo/pkg/util"
)

const Version = "1.1"

func Cmd() {
	util.ClearScreen()
	DATABASE := db.LoadToDo()

	PrintVKTODO(Version, DATABASE)

	var cmd string = ""
	var id int = 0

	fmt.Scanln(&cmd, &id)
	for {
		switch cmd {
		case "add", "a":
			Add(DATABASE)
			Cmd()
		case "complete", "c":
			Complete(id, DATABASE)
			Cmd()
		case "update", "u":
			Update(id, DATABASE)
			Cmd()
		case "delete", "d":
			Delete(id, DATABASE)
			Cmd()
		case "history", "h":
			PrintHistory(DATABASE)
			Cmd()
		case "q":
			util.ClearScreen()
			os.Exit(0)
		default:
			util.ClearScreen()
			Cmd()
		}
	}
}
