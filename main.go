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
	"sort"
	"strconv"
	"time"
)

// Function to always be run when goclitr is run. Sets up goclitr's folder in the user's .config directory.
// Creates two files there:
// - config.json is the config file. So far nothing can be configured there however. To be implemented.
// - dirs.json is where the different folders (here synonymous with projects) the user has used goclitr in are listed.
func ensure_working_environment() {
	user, _ := user.Current()
	jbasefuncs.EnsureDir(user.HomeDir + "/.config/goclitr")
	jbasefuncs.EnsureJson(user.HomeDir + "/.config/goclitr/config.json")
	jbasefuncs.EnsureJson(user.HomeDir + "/.config/goclitr/dirs.json")
}

func addDirtoList() {
	user, _ := user.Current()                                                    // Get current user information
	dir, _ := os.Getwd()                                                         // Get filepath of current folder
	goclitrjson.AppendFolderList(user.HomeDir+"/.config/goclitr/dirs.json", dir) // Write cur. dir to list
	fmt.Println("Initialized at " + dir)                                         // Notify the user.
}

// Function to initialize goclitr in+for this folder.
func initialize() {
	// Create the required files and folder at .goclitr.
	jbasefuncs.EnsureDir(".goclitr")
	jbasefuncs.EnsureJsonList(".goclitr/pending.json")
	jbasefuncs.EnsureJsonList(".goclitr/completed.json")
	addDirtoList()
}

// Function to add new tasks. Creates the task and passes the json-related work to AppendTask from goclitrjson.go.
func addTask(args []string) {
	description := jbasefuncs.JoinSlice(" ", args) // Join arguments to description string.
	username, _ := user.Current()                  // Get username to later store it.

	newtask := goclitrjson.Task{Description: description, User: username.Username, Entry: time.Now().Unix()}
	goclitrjson.AppendTask(".goclitr/pending.json", newtask)
}

// Function to delete a task by ID.
// More than one ID can be passed to delete more than one task with one command.
func removeTask(args []string) {

	// First, get all arguments and convert them to integers to be
	// passed on to deletion function from goclitrjson.go.
	ids := []int{}
	for _, argument := range args {
		arg, err := strconv.Atoi(argument)
		if err != nil {
			fmt.Println("One of the arguments you provided could not be converted to int.")
			os.Exit(1)
		}
		ids = append(ids, arg)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(ids))) // Sort in reverse order.

	for _, id := range jbasefuncs.ArrayIntUnique(ids) { // Remove duplicates for loop
		goclitrjson.RemoveTask(".goclitr/pending.json", id)
		fmt.Println("Removed task #" + fmt.Sprint(id))
	}
}

// -----------------
// Main: Handle command line inputs
// -----------------

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
	} else if jbasefuncs.HandleCmdInput(args, []string{"help"}) {
		goclitrtexts.PrintHelp()
	} else if jbasefuncs.HandleCmdInput(args, []string{"listall"}) {
		goclitrtexts.ListProjects()
	} else if jbasefuncs.HandleCmdInput(args, []string{"add"}) {
		addTask(args[1:])
	} else if jbasefuncs.HandleCmdInput(args, []string{"remove"}) ||
		jbasefuncs.HandleCmdInput(args, []string{"delete"}) {
		removeTask(args[1:])
	} else {
		fmt.Println("Unknown command: " + args[0])
	}

}
