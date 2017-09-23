// -----------------
// The JSON backend to goclitr can be found in this file.
// -----------------
package goclitrjson

import (
	"../jbasefuncs"
	"encoding/json"
	"strconv"
	"time"
)

// ------------------------------------------------
// Set structs and their functions for different types of JSON files.
// ------------------------------------------------

// -----------------
// General functions for handling JSON
// Thanks: https://www.chazzuka.com/2015/03/load-parse-json-file-golang/
// -----------------

func ToJson(p interface{}) string {
	bytes, err := json.Marshal(p)
	jbasefuncs.Check(err)
	return string(bytes)
}

// -----------------
// Task
// -----------------

type Task struct {
	Description string   `json:"description"`
	User        string   `json:"user"`
	Uuid        string   `json:"uuid"`
	Status      string   `json:"status"`
	Entry       int64    `json:"entry"`
	End         int64    `json:"end"`
	Due         int64    `json:"due"`
	Progress    int      `json:"progress"`
	Annotation  []string `json:"annotation"`
	Modified    []int64  `json:"modified"`
}

func (p Task) ToString() string {
	return ToJson(p)
}

// Function for decoding the task list.
func DecodeTask(filename string) []Task {
	file := jbasefuncs.File_get_contents_bytes(filename)

	var data []Task
	err := json.Unmarshal(file, &data)
	jbasefuncs.Check(err)

	return data
}

// Function for appending a task to the local task list.
func AppendTask(filename string, toappend Task) {
	data := DecodeTask(filename)
	data = append(data, toappend)
	jbasefuncs.File_put_contents(filename, ToJson(data))
}

func checkExistentTask(data []Task, key int) {
	// Check if there is a task with this ID.
	if key < 0 || key > len(data) {
		jbasefuncs.Die("No task with this ID existent.") //
	}
}

// Function for modifying a task in the local task list.
func ModifyTask(filename string, key int, toeditKey string, toedit string) bool {

	// Check for bad arguments having been passed
	allowedKeys := []string{"description", "progress"}
	if jbasefuncs.InArrayStr(toeditKey, allowedKeys) != true {
		jbasefuncs.Die("Bad value passed.")
	}

	data := DecodeTask(filename)
	checkExistentTask(data, key) // Check for invalid ID.

	// Make changes here
	if toeditKey == "description" {
		data[key].Description = toedit
	} else if toeditKey == "progress" {
		value, err := strconv.Atoi(toedit)
		jbasefuncs.Check(err)
		if value > 10 || value < 0 {
			jbasefuncs.Die("Progress cannot be larger than 10.")
		}
		if value == 10 {
			data[key].Status = "done"
		}
		data[key].Progress = value
	}

	data[key].Modified = append(data[key].Modified, time.Now().Unix())
	jbasefuncs.File_put_contents(filename, ToJson(data))

	return true
}

// Function for deleting a folder from the user's list.
func RemoveTask(filename string, key int) bool {
	data := DecodeTask(filename)
	checkExistentTask(data, key) // Check for invalid ID.
	data = append(data[:key], data[key+1:]...)
	jbasefuncs.File_put_contents(filename, ToJson(data))
	return true
}

// Function for deleting a folder from the user's list.
func MoveTask(filenameOrigin string, filenameTarget string, key int) bool {
	dataOrigin := DecodeTask(filenameOrigin)
	checkExistentTask(dataOrigin, key) // Check for invalid ID.
	task := dataOrigin[key]            // Get task to transfer.

	dataOrigin = append(dataOrigin[:key], dataOrigin[key+1:]...)
	jbasefuncs.File_put_contents(filenameOrigin, ToJson(dataOrigin))

	dataTarget := DecodeTask(filenameTarget)
	dataTarget = append(dataTarget, task)
	jbasefuncs.File_put_contents(filenameTarget, ToJson(dataTarget))

	return true
}

// -----------------
// Folder list
// The folder list is stored in ~/.config/goclitr/dirs.json as a slice of strings.
// -----------------

// Function for decoding the folder list.
func DecodeFolderList(filename string) []string {
	file := jbasefuncs.File_get_contents_bytes(filename)

	var data []string
	err := json.Unmarshal(file, &data)
	jbasefuncs.Check(err)

	return data
}

// Function for appending a folder to the user's list.
// Should be included in any successfully executed function that does change contents to the current dir.
func AppendFolderList(filename string, toappend string) {
	data := DecodeFolderList(filename)
	data = append(data, toappend)
	data = jbasefuncs.ArrayStringUnique(data)
	jbasefuncs.File_put_contents(filename, ToJson(data))
}

// Function for remove a folder from the user's list.
func DeleteFolderList(filename string, toRemove string) {
	data := DecodeFolderList(filename)
	// Find key
	key := -1
	for i, value := range data {
		if value == toRemove {
			key = i
		}
	}
	if key == -1 { // If key hasn't been changed, abort
		jbasefuncs.Die("Directory not found for removal.")
	}
	data = append(data[:key], data[key+1:]...)
	jbasefuncs.File_put_contents(filename, ToJson(data))
}
