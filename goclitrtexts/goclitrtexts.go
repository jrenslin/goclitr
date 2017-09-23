package goclitrtexts

import (
	"../goclitrjson"
	"../jbasefuncs"
	"fmt"
	color "github.com/fatih/color"
	"os/user"
	"time"
)

func PrintHelp() {
	fmt.Println("\nGoclitr v.0.1 \n ")

	fmt.Printf("%-20s %-20s %-20s \n", "Command", "", "Description")
	fmt.Printf("%-20s %-20s %-20s \n", "init", "", "Initialize")
	fmt.Printf("%-20s %-20s %-20s \n", "current", "", "Lists currently active (=not completed) issues")
	fmt.Printf("%-20s %-20s %-20s \n", "listall", "", "Lists all projects you've worked on")

}

func ListIssues() {
	tasks := goclitrjson.DecodeTask(".goclitr/pending.json")

	headline := color.New(color.Underline)
	headline.Printf("%-6s %-8s %-30s", "ID", "Age", "Description")
	unequal := color.New(color.BgYellow, color.FgBlack)

	for i, p := range tasks {
		age := time.Now().Unix() - p.Entry
		if i%2 == 1 {
			fmt.Printf("\n%-6s %-8s %-30s \n", fmt.Sprint(i), jbasefuncs.ReadableTime(age, true), p.Description)
		} else {
			unequal.Printf("\n%-6s %-8s %-30s", fmt.Sprint(i), jbasefuncs.ReadableTime(age, true), p.Description)
		}
	}
}

func ListProjects() {
	user, _ := user.Current()
	folders := goclitrjson.DecodeFolderList(user.HomeDir + "/.config/goclitr/dirs.json")

	fmt.Println("\nYou've worked on the following " + fmt.Sprint(len(folders)) + " projects\n ")
	for _, p := range folders {
		fmt.Println(" - " + p)
	}
	fmt.Println("")
}
