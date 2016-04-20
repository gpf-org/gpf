package util

import (
	"encoding/json"
	"os"
)

func PrintJSON(data []byte) {
	var token map[string]interface{}
	if err := json.Unmarshal(data, &token); err != nil {
		panic(err)
	}

	result, _ := json.MarshalIndent(token, "", "  ")
	os.Stdout.Write(result)
}

func PrintJSONList(data []byte) {
	var token []map[string]interface{}
	if err := json.Unmarshal(data, &token); err != nil {
		panic(err)
	}

	result, _ := json.MarshalIndent(token, "", "  ")
	os.Stdout.Write(result)
}
