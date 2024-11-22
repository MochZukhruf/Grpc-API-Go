package helpers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

func ReadJSONFile(filePath string, v interface{}) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, v)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
		return err
	}
	return nil
}
