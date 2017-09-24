// -----------------
// Basic, useful functions as existent in other languages are reimplemented here
// -----------------

package jbasefuncs

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// -----------------
// Function to check if there is a
// -----------------

func FileExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}

// -----------------
// The following functions EnsureDir and EnsureJson are used to create an empty JSON file / an empty directy if there is none existent yet at the specified filepath.
// -----------------

func EnsureDir(folder string) {
	if _, err := os.Stat(folder); os.IsNotExist(err) {
		os.Mkdir(folder, 0755)
	}
}

func EnsureJson(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		File_put_contents(path, "{}")
	}
}

func EnsureJsonList(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		File_put_contents(path, "[]")
	}
}

// -----------------
// The following functions are basic ports of PHP's file_get_contents and file_put_contents for local usage. As e.g. parsing JSON requires a byte map, not string, a new function for reading the file contents to that without parsing it to a string has been added.
// -----------------

func File_get_contents_bytes(path string) []byte {
	file, e := ioutil.ReadFile(path)
	Check(e)
	return file
}

func File_get_contents(path string) string {
	file, e := ioutil.ReadFile(path)
	Check(e)
	return string(file)
}

func File_put_contents(path string, contents string) {
	d1 := []byte(contents)
	err := ioutil.WriteFile(path, d1, 0644)
	Check(err)
}

// -----------------
// Scandir (also following PHP's scandir) lists the contents of a folder
// -----------------

func Scandir(folder string) []string {
	// Ensure that the provided filepath *folder* ends with a string
	if strings.HasSuffix(folder, "/") == false {
		folder += "/"
	}
	files, _ := filepath.Glob(folder + "*")
	return files
}

func ScandirPlus(folder string, selector string) []string {
	// Ensure that the provided filepath *folder* ends with a string
	if strings.HasSuffix(folder, "/") == false {
		folder += "/"
	}
	files, _ := filepath.Glob(folder + selector)
	return files
}

// -----------------
// Check if two slices of strings equal each other
// -----------------

func TestEqSliceStrings(a, b []string) bool {

	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

// -----------------
// Check is
// -----------------

func Check(e error) {
	if e != nil {
		fmt.Println(e.Error())
		os.Exit(1)
	}
}

func Die(str string) {
	fmt.Println(str)
	os.Exit(1)
}

// -----------------
// Reimplement array_unique, but based on type of slice
// - Thanks: https://www.dotnetperls.com/duplicates-go
// -----------------

func ArrayIntUnique(elements []int) []int {

	// Use map to record duplicates as we find them.
	encountered := map[int]bool{}
	result := []int{}

	for v := range elements {
		if encountered[elements[v]] == true {
			// Do not add duplicate.
		} else {
			// Record this element as an encountered element.
			encountered[elements[v]] = true
			// Append to result slice.
			result = append(result, elements[v])
		}
	}
	// Return the new slice.
	return result
}

// The same for arrays of strings
func ArrayStringUnique(elements []string) []string {

	// Use map to record duplicates as we find them.
	encountered := map[string]bool{}
	result := []string{}

	for v := range elements {
		if encountered[elements[v]] == true {
			// Do not add duplicate.
		} else {
			// Record this element as an encountered element.
			encountered[elements[v]] = true
			// Append to result slice.
			result = append(result, elements[v])
		}
	}
	// Return the new slice.
	return result
}

// Reimplementation of PHP's in_array
func InArrayStr(needle string, haystack []string) bool {
	output := false
	for _, value := range haystack {
		if needle == value {
			output = true
		}
	}
	return output
}

// -----------------
// Function to join a slice of strings to a single string
// -----------------

func JoinSlice(joinwith string, list []string) string {
	output := ""
	for _, p := range list {
		output += p + joinwith
	}
	return output[:len(output)-len(joinwith)]
}

// -----------------
// Struct for describing categories of timespans, e.g. minutes and their relation to seconds.
// Used to create more understandable output than seconds. See function ReadableTime.
// -----------------

type timecorrespondence struct {
	duration   int64  // Duration in seconds
	descriptor string // Abbreviated form (e.g. h for hour)
}

// -----------------
// Function to convert a time difference (e.g. age) to a human-readable time.
// -----------------

func ReadableTime(seconds int64, roundto bool) string { // Parameter roundto is not used, but included for futher use later.
	correspondence := []timecorrespondence{
		timecorrespondence{duration: 78840000, descriptor: "y"}, // Counts the year as 365 days, ignoring leap years.
		timecorrespondence{duration: 1512000, descriptor: "w"},
		timecorrespondence{duration: 216000, descriptor: "d"},
		timecorrespondence{duration: 3600, descriptor: "h"},
		timecorrespondence{duration: 60, descriptor: "m"},
	}
	for _, timespan := range correspondence {
		if seconds > timespan.duration {
			return fmt.Sprint(seconds/timespan.duration) + timespan.descriptor
		}
	}
	return fmt.Sprint(seconds) + "s"
}

// -----------------
// Shortcut for handling command line inputs
// -----------------

func HandleCmdInput(args []string, condition []string) bool {
	if len(args) < len(condition) {
		return false
	} else if TestEqSliceStrings(args[:len(condition)], condition) {
		return true
	} else {
		return false
	}
}
