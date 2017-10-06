// -----------------
// Folder list
// The folder list is stored in ~/.config/goclitr/dirs.json as a slice of strings.
// -----------------
package goclitrjson

import (
	"../jbasefuncs"
	"encoding/json"
)

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
