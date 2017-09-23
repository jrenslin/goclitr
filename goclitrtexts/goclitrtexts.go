package goclitrtexts

import (
	"../goclitrjson"
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
	headline.Printf("%-20s %-20s %-20s \n", "Name", "Age", "")

	for _, p := range tasks {
		age, _ := time.Parse()
		fmt.Printf("%-20s %-20s %-20s \n", p.Description, "", "")
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
