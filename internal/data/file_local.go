package data

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type LocalFile struct {
	FilePath   string
	Parameters Parameters
}

func (p LocalFile) Read() (Parameters, error) {
	log.Println("[DEBUG] Reading from LocalFile")
	jsonBlob, err := os.ReadFile(p.FilePath)
	if err != nil {
		log.Println(err)
	}

	err = json.Unmarshal(jsonBlob, &p.Parameters)
	if err != nil {
		log.Println(err)
	}
	log.Println("[DEBUG]", p.Parameters)
	return p.Parameters, nil
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
