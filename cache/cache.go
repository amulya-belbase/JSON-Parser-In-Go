package cache

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// define the strucutre of proceses inside json files
type Process struct {
	ProcessId  string      `json:"processId"`
	First      interface{} `json:"first"`
	Second     interface{} `json:"second"`
	MappedFunc string      `json:"mappedFunc"`
	Predicate  string      `json:"predicate,omitempty"`
}

// individual json file structure
type ProcessAction struct {
	Action    string    `json:"action"`
	Processes []Process `json:"processes"`
}

var AllActions = make(map[string]ProcessAction)

func init() {
	dirName := "configs"

	files, err := os.ReadDir(dirName)
	if err != nil {
		log.Fatalf("failed to read directory: %v", err)
	}

	for _, file := range files {
		fileName := fmt.Sprintf("%s/%s", dirName, file.Name())

		jsonFile, err := os.Open(fileName)
		if err != nil {
			log.Printf("failed to open json file %s: %v", fileName, err)
			continue
		}

		var jsonData ProcessAction
		reader := bufio.NewReader(jsonFile)
		err = json.NewDecoder(reader).Decode(&jsonData)
		jsonFile.Close()

		if err != nil {
			log.Printf("failed to decode json file %s: %v", fileName, err)
			continue
		}

		AllActions[jsonData.Action] = jsonData
	}
}
