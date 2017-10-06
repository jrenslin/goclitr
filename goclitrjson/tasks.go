// -----------------
// Tasks
// -----------------
package goclitrjson

import (
	"../jbasefuncs"
	"encoding/json"
	"strconv"
	"time"
)

type Task struct {
	Description string       `json:"description"`
	User        string       `json:"user"`
	Uuid        string       `json:"uuid"`
	Status      string       `json:"status"`
	Entry       int64        `json:"entry"`
	End         int64        `json:"end"`
	Due         int64        `json:"due"`
	Progress    int          `json:"progress"`
	Annotation  []Annotation `json:"annotation"`
	Modified    []int64      `json:"modified"`
}

func (p Task) ToString() string {
	return ToJson(p)
}

// Annotations

type Annotation struct {
	Text  string
	User  string
	Entry int64
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

// Function for appending a task to the local task list.
func AddAnnotation(filename string, key int, annotation Annotation) {
	data := DecodeTask(filename)
	CheckExistentTask(data, key) // Check for invalid ID.
	data[key].Annotation = append(data[key].Annotation, annotation)
	jbasefuncs.File_put_contents(filename, ToJson(data))
}

func CheckExistentTask(data []Task, key int) {
	// Check if there is a task with this ID.
	if key < 0 || key >= len(data) {
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
	CheckExistentTask(data, key) // Check for invalid ID.

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
	CheckExistentTask(data, key) // Check for invalid ID.
	data = append(data[:key], data[key+1:]...)
	jbasefuncs.File_put_contents(filename, ToJson(data))
	return true
}

// Function for deleting a folder from the user's list.
func MoveTask(filenameOrigin string, filenameTarget string, key int) bool {
	dataOrigin := DecodeTask(filenameOrigin)
	CheckExistentTask(dataOrigin, key) // Check for invalid ID.
	task := dataOrigin[key]            // Get task to transfer.

	dataOrigin = append(dataOrigin[:key], dataOrigin[key+1:]...)
	jbasefuncs.File_put_contents(filenameOrigin, ToJson(dataOrigin))

	dataTarget := DecodeTask(filenameTarget)
	dataTarget = append(dataTarget, task)
	jbasefuncs.File_put_contents(filenameTarget, ToJson(dataTarget))

	return true
}
