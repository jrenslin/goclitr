// This file provides functions that mainly contain strings for outputs in goclitr.
package goclitrtexts

import (
	"../goclitrjson"
	"../jbasefuncs"
	"fmt"
	color "github.com/fatih/color"
	"os/user"
	"time"
)

// Help text
func PrintHelp() {
	fmt.Println("Goclitr v.0.1 \n ")

	// Formatting
	table := "%-12s %-4s %-12s %-20s \n"
	headline := color.New(color.Underline)

	// Print Table
	headline.Printf(table, "Command", "", "", "Description")
	fmt.Printf(table, "init", "", "", "Initialize")
	fmt.Printf(table, "teardown", "", "", "Tear down")
	fmt.Printf(table, "current", "", "", "Lists currently active (=not completed) issues")
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
	fmt.Printf(table, "listall", "", "", "Lists all projects you've worked on")
	fmt.Printf(table, "project", "<Project ID>", "", "Return path of project X")

}

// Function listing the tasks / issues within this project / folder.
func ListIssues() {
	tasks := goclitrjson.DecodeTask(".goclitr/pending.json")

	// Formatting
	table := "%2s %-8s %-30s %-10s %8s %11s"
	headline := color.New(color.Underline)
	unequal := color.New(color.BgYellow, color.FgBlack)

	headline.Printf(table, "ID", "Age", "Description", "User", "Progress", "Annotations")
	for i, p := range tasks {
		age := time.Now().Unix() - p.Entry
		fmt.Println("") // Without this, the background color would fill the entire line.
		switch {
		case i%2 == 1:
			fmt.Printf(table, fmt.Sprint(i), jbasefuncs.ReadableTime(age, true), p.Description,
				p.User, fmt.Sprint(p.Progress), fmt.Sprint(len(p.Annotation)))
		default:
			unequal.Printf(table, fmt.Sprint(i), jbasefuncs.ReadableTime(age, true), p.Description,
				p.User, fmt.Sprint(p.Progress), fmt.Sprint(len(p.Annotation)))
		}
	}
	fmt.Println("")
}

// Function listing the tasks / issues within this project / folder.
func ListCompleted() {
	tasks := goclitrjson.DecodeTask(".goclitr/completed.json")

	// Formatting
	table := "%2s %-10s  %-10s %-30s %-10s"
	headline := color.New(color.Underline)
	headline.Printf(table, "ID", "Entry", "Duration", "Description", "Creator")
	unequal := color.New(color.BgYellow, color.FgBlack)

	for i, p := range tasks {
		age := p.Modified[len(p.Modified)-1] - p.Entry
		fmt.Println("") // Without this, the background color would fill the entire line.
		switch {
		case i%2 == 1:
			fmt.Printf(table, fmt.Sprint(i), time.Unix(p.Entry, 0).Format("2006-01-02"),
				jbasefuncs.ReadableTime(age, true), p.Description, p.User)
		default:
			unequal.Printf(table, fmt.Sprint(i), time.Unix(p.Entry, 0).Format("2006-01-02"),
				jbasefuncs.ReadableTime(age, true), p.Description, p.User)
		}
	}
	fmt.Println("")
}

// Function to list all projects (say, directories) the user has contributed to.
func ListProjects() {
	user, _ := user.Current()
	folders := goclitrjson.DecodeFolderList(user.HomeDir + "/.config/goclitr/dirs.json")

	fmt.Println("You've worked on the following " + fmt.Sprint(len(folders)) + " projects\n ")

	// Formatting
	headline := color.New(color.Underline)
	unequal := color.New(color.BgYellow, color.FgBlack)

	headline.Printf("%-6s %-40s", "ID", "Path")
	for i, p := range folders {
		fmt.Println("") // Without this, the background color would fill the entire line.
		switch {
		case i%2 == 1:
			fmt.Printf("%-6s %-40s", fmt.Sprint(i), p)
		default:
			unequal.Printf("%-6s %-40s", fmt.Sprint(i), p)
		}

	}
	fmt.Println("")
}

// Printing details of a task
func PrintTasks(task goclitrjson.Task) {
	fmt.Println("Details ...")

	// Formatting
	table := "%-20s %-20s"
	headline := color.New(color.Underline)
	unequal := color.New(color.BgYellow, color.FgBlack)

	// Print table
	headline.Printf(table, "Name", "Value")
	fmt.Printf("\n"+table+"\n", "Description", task.Description)
	unequal.Printf(table, "Creator", task.User)
	fmt.Printf("\n"+table, "Entry", time.Unix(task.Entry, 0).Format("2006-01-02"))

	for i, p := range task.Modified {
		fmt.Println("")
		switch {
		case i%2 == 1:
			fmt.Printf(table, "Modification #"+fmt.Sprint(i), time.Unix(p, 0).Format("2006-01-02 15:04"))
		default:
			unequal.Printf(table, "Modification #"+fmt.Sprint(i), time.Unix(p, 0).Format("2006-01-02 15:04"))
		}
	}

	// Generate progress bar and print it
	headline.Println("\n\nProgress")
	progressbar := ""
	for i := 0; i <= 10; i++ {
		switch {
		case i < task.Progress:
			progressbar += "==="
		case i == task.Progress:
			progressbar += "==>"
		case i > task.Progress:
			progressbar += "———"
		}
	}
	fmt.Printf("\n%-33s %2s / %2s\n\n", progressbar, fmt.Sprint(task.Progress), "10")

	// Print annotations
	headline.Println("Annotations")
	for _, p := range task.Annotation {
		fmt.Printf("\n %s wrote on %s: \n", p.User, time.Unix(p.Entry, 0).Format("2006-01-02 15:04"))
		fmt.Println(p.Text)
	}

}
