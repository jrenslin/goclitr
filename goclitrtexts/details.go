// Printing details of a task

package goclitrtexts

import (
	"../goclitrjson"
	"fmt"
	color "github.com/fatih/color"
	"time"
)

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
