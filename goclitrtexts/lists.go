// This file provides functions that mainly contain strings for outputs in goclitr.
package goclitrtexts

import (
	"../goclitrjson"
	jbasefuncs "github.com/jrenslin/jbasefuncs"
	"fmt"
	color "github.com/fatih/color"
	"os/user"
	"time"
)

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
