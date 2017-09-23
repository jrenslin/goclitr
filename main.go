// -----------------
// Main package of goclitr.
// Handles the command line interface.
// -----------------
package main

import (
	"./goclitrjson"
	"./goclitrtexts"
	"./jbasefuncs"
	"fmt"
	"os"
	"os/user"
	"time"
)

type Settings struct {
	Time string
	User string
	Dir  string
}

func ensure_working_environment() {
	user, _ := user.Current()
	jbasefuncs.EnsureDir(user.HomeDir + "/.config/goclitr")
	jbasefuncs.EnsureJson(user.HomeDir + "/.config/goclitr/config.json")
	jbasefuncs.EnsureJson(user.HomeDir + "/.config/goclitr/dirs.json")
}

func initialize() {
	jbasefuncs.EnsureDir(".goclitr")
	jbasefuncs.EnsureJsonList(".goclitr/pending.json")
	jbasefuncs.EnsureJsonList(".goclitr/completed.json")

	// Get user info
	user, _ := user.Current()
	dir, _ := os.Getwd()
	goclitrjson.AppendFolderList(user.HomeDir+"/.config/goclitr/dirs.json", dir)

	fmt.Println("Initialized at " + dir)
}

func addTask(args []string) {
	description := jbasefuncs.JoinSlice(" ", args)
	username, _ := user.Current()

	newtask := goclitrjson.Task{Description: description, User: username.Username, Entry: time.Now().Unix()}
	goclitrjson.AppendTask(".goclitr/pending.json", newtask)
}

// -----------------
// Main: Handle command line inputs
// -----------------

func Ausgelager() {
	pwd, err := os.Getwd()
	jbasefuncs.Check(err)
	Settings := Settings{User: "X", Time: time.Now().Format(time.RFC850), Dir: pwd}
	fmt.Println(Settings)
}

func main() {
	ensure_working_environment()

	args := os.Args[1:]

	if len(args) == 0 && jbasefuncs.FileExists(".goclitr") == false { // If
		goclitrtexts.PrintHelp()
		goclitrtexts.ListIssues()
	} else if len(args) == 0 {
		goclitrtexts.ListIssues()
	} else if jbasefuncs.HandleCmdInput(args, []string{"init"}) {
		initialize()
	} else if jbasefuncs.HandleCmdInput(args, []string{"listall"}) {
		goclitrtexts.ListProjects()
	} else if jbasefuncs.HandleCmdInput(args, []string{"add"}) {
		addTask(args[1:])
	} else {
		fmt.Println("Unknown command: " + args[0])
	}

	/*
		A := []string{"hi"}
		b := "hia"
		a = append(a, b)
	*/

}
