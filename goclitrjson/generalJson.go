// -----------------
// JSON-based backend for goclitr.
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
