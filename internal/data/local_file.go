package data

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/cdalar/parampiper/pkg/common"
)

type LocalFile struct {
	FilePath string
}

func (p LocalFile) Read() (ParampiperData, error) {
	log.Println("[DEBUG] Reading from LocalFile")
	ppData := ParampiperData{}
	if common.FileExists(p.FilePath) {
		jsonBlob, err := os.ReadFile(p.FilePath)
		if err != nil {
			log.Println(p.FilePath, "file does not exist")
			return ParampiperData{}, err
		}
		err = json.Unmarshal(jsonBlob, &ppData)
		if err != nil {
			return ParampiperData{}, err
		}
	} else {
		log.Println(p.FilePath, "file does not exist. Creating file")
		ppData = ParampiperData{
			Version:    DATA_FORMAT_VERSION,
			Parameters: Parameters{},
		}
		JsonData, err := json.MarshalIndent(ppData, "", "    ")
		if err != nil {
			log.Println(err)
		}

		err = os.WriteFile(p.FilePath, JsonData, 0644)
		if err != nil {
			log.Println(err)
		}
	}
	log.Println("[DEBUG]", ppData)
	return ppData, nil
}

func (p LocalFile) Save(ppData ParampiperData) error {
	ppData.Version = DATA_FORMAT_VERSION
	JsonData, err := json.MarshalIndent(ppData, "", "    ")
	if err != nil {
		log.Println(err)
	}

	err = os.WriteFile(p.FilePath, JsonData, 0644)
	if err != nil {
		fmt.Println("error:", err)
	}

	return nil
}

func (p LocalFile) String() string {
	s := ""
	data, err := p.Read()
	if err != nil {
		log.Println(err)
	}

	for _, parameter := range data.Parameters {
		s += parameter.String() + "\n"
	}
	return s
}
