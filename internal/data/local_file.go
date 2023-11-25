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

func (p LocalFile) Read() (Parameters, error) {
	log.Println("[DEBUG] Reading from LocalFile")
	jsonBlob, err := os.ReadFile(p.FilePath)
	if err != nil {
		log.Println(err)
	}
	parameters := Parameters{}
	err = json.Unmarshal(jsonBlob, &parameters)
	if err != nil {
		log.Println(err)
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
