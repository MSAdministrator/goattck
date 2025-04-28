package models

import (
	"encoding/json"
	"fmt"
	"os"
)

// jsonToMap converts a json string to a map
func jsonToMap(jsonStr string) map[string]interface{} {
	result := make(map[string]interface{})
	json.Unmarshal([]byte(jsonStr), &result)
	return result
}

// loadTestJSON loads a json file from the tests/data directory
func loadTestJSON(fileName string) map[string]interface{} {
	content, err := os.ReadFile(fmt.Sprintf("../../tests/data/%s", fileName))
	if err != nil {
		fmt.Println(err)
	}
	return jsonToMap(string(content))
}
