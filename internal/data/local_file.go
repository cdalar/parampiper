package data

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type LocalFile struct {
	FilePath string
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func (p LocalFile) Read() (Parameters, error) {
	log.Println("[DEBUG] Reading from LocalFile")
	parameters := Parameters{}
	if fileExists(p.FilePath) {
		jsonBlob, err := os.ReadFile(p.FilePath)
		if err != nil {
			log.Println(err)
		}
		err = json.Unmarshal(jsonBlob, &parameters)
		if err != nil {
			log.Println(err)
		}
	} else {
		fmt.Println(p.FilePath, "file does not exist")
		os.Exit(1)
	}
	log.Println("[DEBUG]", parameters)
	return parameters, nil
}

func (p LocalFile) Save(params Parameters) error {
	JsonData, err := json.MarshalIndent(params, "", "    ")
	if err != nil {
		log.Println(err)
	}

	err = os.WriteFile(p.FilePath, JsonData, 0644)
	if err != nil {
		fmt.Println("error:", err)
	}

	return nil
}
