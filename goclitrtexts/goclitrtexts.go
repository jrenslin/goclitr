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

	headline := color.New(color.Underline)
	headline.Printf("%-12s %-12s %-20s \n", "Command", "", "Description")
	fmt.Printf("%-12s %-12s %-20s \n", "init", "", "Initialize")
	fmt.Printf("%-12s %-12s %-20s \n", "current", "", "Lists currently active (=not completed) issues")
	fmt.Printf("%-12s %-12s %-20s \n", "listall", "", "Lists all projects you've worked on")

}

// Function listing the tasks / issues within this project / folder.
func ListIssues() {
	tasks := goclitrjson.DecodeTask(".goclitr/pending.json")

	headline := color.New(color.Underline)
	headline.Printf("%-6s %-8s %-30s %-10s", "ID", "Age", "Description", "User")
	unequal := color.New(color.BgYellow, color.FgBlack)

	for i, p := range tasks {
		age := time.Now().Unix() - p.Entry
		if i%2 == 1 {
			fmt.Printf("\n%-6s %-8s %-30s %-10s", fmt.Sprint(i), jbasefuncs.ReadableTime(age, true), p.Description, p.User)
		} else {
			unequal.Printf("\n%-6s %-8s %-30s %-10s", fmt.Sprint(i), jbasefuncs.ReadableTime(age, true), p.Description, p.User)
		}
	}
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
