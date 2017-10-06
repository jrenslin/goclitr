// Prints help information to command line.
package goclitrtexts

import (
	"fmt"
	color "github.com/fatih/color"
)

// Help text
func PrintHelp() {
	fmt.Println("Goclitr v.0.1 \n ")

	// Formatting
	table := "%-12s %-12s %-12s %-20s \n"
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
	fmt.Printf(table, "new", "", "", "Add a task (same as add)")
	fmt.Printf(table, "delete", "<ID>", "", "Delete task with the given ID")
	fmt.Printf(table, "remove", "<ID>", "", "Delete task (same as delete)")
	fmt.Printf(table, "modify", "<ID>", "<text>", "Modify the task's text")
	fmt.Printf(table, "progress", "<ID>", "<int: 0-10>", "Edit progress of the task's text")
	fmt.Printf(table, "annotate", "<ID>", "<text>", "Annotate a task")
	fmt.Printf(table, "done", "<ID>", "", "Finish task")
	fmt.Printf(table, "complete", "<ID>", "", "Finish task (same as done)")
	fmt.Printf(table, "listall", "", "", "Lists all projects you've worked on")
	fmt.Printf(table, "project", "<Project ID>", "", "Return path of project X")

}
