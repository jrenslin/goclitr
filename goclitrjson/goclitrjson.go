// -----------------
// The JSON backend to goclitr can be found in this file.
// -----------------
package goclitrjson

import (
	"../jbasefuncs"
	"encoding/json"
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
	Description string `json:"description"`
	User        string `json:"user"`
	Uuid        string `json:"uuid"`
	Status      string `json:"status"`
	Entry       int64  `json:"entry"`
	End         int64  `json:"end"`
	Due         int64  `json:"due"`
	Progress    string `json:"progress"`
	Annotation  string `json:"annotation"`
	Modified    []int  `json:"modified"`
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

// Function for deleting a folder from the user's list.
func RemoveTask(filename string, key int) {
	data := DecodeTask(filename)
	data = append(data[:key], data[key+1:]...)
	jbasefuncs.File_put_contents(filename, ToJson(data))
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
