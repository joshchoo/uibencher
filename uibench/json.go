package uibench

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func loadConfig(filePath string) Config {
	file, err := os.Open(filePath)
	if err != nil {
		abort("Error opening " + filePath + ".")
	}
	defer file.Close()

	// Get the byte value of the JSON file
	byteFile, _ := ioutil.ReadAll(file)

	var cfg Config
	err = json.Unmarshal(byteFile, &cfg)
	if err != nil {
		abort("Error loading " + filePath + ". Please check if the JSON format is valid.")
	}

	return cfg
}
