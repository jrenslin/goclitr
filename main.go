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

// Function for modifying tasks
func modifyTask(args []string, tomodify string) {
	// Convert the first left over argument to int
	id, err := strconv.Atoi(args[0])
	if err != nil {
		jbasefuncs.Die("You did not provide a valid ID.")
	}
	description := jbasefuncs.JoinSlice(" ", args[1:]) // Join arguments to description string.
	if goclitrjson.ModifyTask(".goclitr/pending.json", id, tomodify, description) {
		fmt.Println("Modified: " + tomodify + " >> " + args[1])
	}
}

// Function for modifying tasks
func finishTask(args []string) {
	// Convert the first left over argument to int
	id, err := strconv.Atoi(args[0])
	if err != nil {
		jbasefuncs.Die("You did not provide a valid ID.")
	}
	if goclitrjson.ModifyTask(".goclitr/pending.json", id, "progress", "10") {
		fmt.Println("Modified: progress >> 10")
	}
	goclitrjson.MoveTask(".goclitr/pending.json", ".goclitr/completed.json", id)
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
			jbasefuncs.Die("One of the arguments you provided could not be converted to int.")
		}
		ids = append(ids, arg)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(ids))) // Sort in reverse order.

	for _, id := range jbasefuncs.ArrayIntUnique(ids) { // Remove duplicates for loop
		goclitrjson.RemoveTask(".goclitr/pending.json", id)
		fmt.Println("Removed task #" + fmt.Sprint(id))
	}
}

// Function to tear down local tasklist.
func tearDown() {
	if jbasefuncs.FileExists(".goclitr") == false {
		jbasefuncs.Die("There is nothing to tear down here.")
	}
	os.RemoveAll(".goclitr/")
	fmt.Println("Removed .goclitr directory.")

	user, _ := user.Current() // Get username to later store it.
	dir, _ := os.Getwd()      // Get filepath of current folder
	goclitrjson.DeleteFolderList(user.HomeDir+"/.config/goclitr/dirs.json", dir)
	fmt.Println("Removed local directory from your project list.")
}

// Function to just output path of a project specified by ID
// Unfortunately making the terminal go to the directory from Go seems to be impossible.
// By returning the path, this function makes it easy to jump projects with a little bash.
// E.g.:
// cd `~/Sync/Programming/Golang/goclitr/goclitr goto 1`

func showProjectDir(args []string) {
	user, _ := user.Current()
	folders := goclitrjson.DecodeFolderList(user.HomeDir + "/.config/goclitr/dirs.json")
	key, _ := strconv.Atoi(args[0])
	targetDir := folders[key]
	fmt.Println(targetDir)
}

// Function to select a task to be displayed in detail.

func showIssue(args []string) {
	key, err := strconv.Atoi(args[0])
	if err != nil {
		jbasefuncs.Die("Invalid ID")
	}
	task := goclitrjson.DecodeTask(".goclitr/pending.json")
	goclitrjson.CheckExistentTask(task, key)
	goclitrtexts.PrintTasks(task[key])
}

// -----------------
// Main: Handle command line inputs
// -----------------

func main() {
	ensure_working_environment()

	args := os.Args[1:]

	switch {
	case len(args) == 0 && jbasefuncs.FileExists(".goclitr") == false:
		goclitrtexts.PrintHelp()
		fmt.Printf("\n--------------------------\n\n")
		goclitrtexts.ListProjects()
	case len(args) == 0:
		goclitrtexts.ListIssues()
	case jbasefuncs.HandleCmdInput(args, []string{"list"}):
		goclitrtexts.ListIssues()
	case jbasefuncs.HandleCmdInput(args, []string{"completed"}):
		goclitrtexts.ListCompleted()
	case jbasefuncs.HandleCmdInput(args, []string{"init"}):
		initialize()
	case jbasefuncs.HandleCmdInput(args, []string{"teardown"}):
		tearDown()
	case jbasefuncs.HandleCmdInput(args, []string{"help"}):
		goclitrtexts.PrintHelp()
	case jbasefuncs.HandleCmdInput(args, []string{"listall"}):
		goclitrtexts.ListProjects()
	case jbasefuncs.HandleCmdInput(args, []string{"add"}) ||
		jbasefuncs.HandleCmdInput(args, []string{"new"}):
		addTask(args[1:])
	case jbasefuncs.HandleCmdInput(args, []string{"modify"}):
		modifyTask(args[1:], "description")
	case len(args) == 2 && args[0] == "project":
		showProjectDir(args[1:])
	case len(args) == 2 && args[0] == "show":
		showIssue(args[1:])
	case (len(args) == 3 && args[0] == "progress" && args[2] == "10") ||
		(len(args) == 2 && args[0] == "done") ||
		(len(args) == 2 && args[0] == "complete"):
		finishTask(args[1:])
	case len(args) == 3 && args[0] == "progress":
		modifyTask(args[1:], "progress")
	case jbasefuncs.HandleCmdInput(args, []string{"remove"}) ||
		jbasefuncs.HandleCmdInput(args, []string{"delete"}):
		removeTask(args[1:])
	default:
		fmt.Println("Unknown command: " + args[0])
	}
}
