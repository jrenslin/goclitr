package goclitrtexts

import (
	"../goclitrjson"
	"../jbasefuncs"
	"fmt"
	color "github.com/fatih/color"
	"os/user"
	"time"
)

//
func PrintHelp() {
	fmt.Println("Goclitr v.0.1 \n ")

	table := "%-12s %-4s %-12s %-20s \n"
	headline := color.New(color.Underline)
	headline.Printf(table, "Command", "", "", "Description")
	fmt.Printf(table, "init", "", "", "Initialize")
	fmt.Printf(table, "teardown", "", "", "Tear down")
	fmt.Printf(table, "current", "", "", "Lists currently active (=not completed) issues")
	fmt.Printf(table, "listall", "", "", "Lists all projects you've worked on")
	fmt.Printf(table, "help", "", "", "Print this message")
	fmt.Printf(table, "list", "", "", "List current tasks")
	fmt.Printf(table, "completed", "", "", "List completed")
	fmt.Printf(table, "add", "<text>", "", "Add a task")
	fmt.Printf(table, "new", "", "", "Same as add")
	fmt.Printf(table, "delete", "<ID>", "", "Delete task with the given ID")
	fmt.Printf(table, "remove", "<ID>", "", "Same as delete")
	fmt.Printf(table, "modify", "<ID>", "<text>", "Modify the task's text")
	fmt.Printf(table, "progress", "<ID>", "<int: 1-10>", "Edit progress of the task's text")
	fmt.Printf(table, "done", "<ID>", "", "Finish task")
	fmt.Printf(table, "completed", "<ID>", "", "Finish task")

}

// Function listing the tasks / issues within this project / folder.
func ListIssues() {
	tasks := goclitrjson.DecodeTask(".goclitr/pending.json")

	headline := color.New(color.Underline)
	headline.Printf("%2s %-8s %-30s %-10s %-8s", "ID", "Age", "Description", "User", "Progress")
	unequal := color.New(color.BgYellow, color.FgBlack)

	//	fmt.Println(tasks)
	for i, p := range tasks {
		age := time.Now().Unix() - p.Entry
		if i%2 == 1 {
			fmt.Println("")
			fmt.Printf("%2s %-8s %-30s %-10s %8s", fmt.Sprint(i), jbasefuncs.ReadableTime(age, true), p.Description, p.User, fmt.Sprint(p.Progress))
		} else {
			fmt.Println("")
			unequal.Printf("%2s %-8s %-30s %-10s %8s", fmt.Sprint(i), jbasefuncs.ReadableTime(age, true), p.Description, p.User, fmt.Sprint(p.Progress))
		}
	}
	fmt.Println("")
}

// Function listing the tasks / issues within this project / folder.
func ListCompleted() {
	tasks := goclitrjson.DecodeTask(".goclitr/completed.json")

	headline := color.New(color.Underline)
	headline.Printf("%2s %-10s %-30s %-10s %-8s", "ID", "Duration", "Description", "User", "Progress")
	unequal := color.New(color.BgYellow, color.FgBlack)

	for i, p := range tasks {
		age := p.Modified[len(p.Modified)-1] - p.Entry
		fmt.Println("")
		if i%2 == 1 {
			fmt.Printf("%2s %-10s %-30s %-10s %8s", fmt.Sprint(i), jbasefuncs.ReadableTime(age, true), p.Description, p.User, fmt.Sprint(p.Progress))
		} else {
			unequal.Printf("%2s %-10s %-30s %-10s %8s", fmt.Sprint(i), jbasefuncs.ReadableTime(age, true), p.Description, p.User, fmt.Sprint(p.Progress))
		}
	}
	fmt.Println("")
}

// Function to list all projects (say, directories) the user has contributed to.
func ListProjects() {
	user, _ := user.Current()
	folders := goclitrjson.DecodeFolderList(user.HomeDir + "/.config/goclitr/dirs.json")

	fmt.Println("You've worked on the following " + fmt.Sprint(len(folders)) + " projects\n ")

	headline := color.New(color.Underline)
	headline.Printf("%-6s %-40s", "ID", "Path")
	unequal := color.New(color.BgYellow, color.FgBlack)

	for i, p := range folders {
		if i%2 == 1 {
			fmt.Printf("\n%-6s %-40s", fmt.Sprint(i), p)
		} else {
			unequal.Printf("\n%-6s %-40s", fmt.Sprint(i), p)
		}

	}
	fmt.Println("")
}
