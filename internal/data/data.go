package data

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

var Params Parameters

type DataInterface interface {
	read() error
	save() error
}

type Parameters []Parameter

type Parameter struct {
	Name  string `json:"name"`
	Value string `json:"value"`
	Info  string `json:"info"`
}

func (p Parameter) String() string {
	return fmt.Sprintf("%s: %s", p.Name, p.Value)
}

func (p Parameter) add() {
	jsonBlob, err := os.ReadFile("parampiper.json")
	if err != nil {
		log.Println(err)
	}

	err = json.Unmarshal(jsonBlob, &Params)
	if err != nil {
		log.Println(err)
	}
	if p.ifExists() {
		fmt.Println("Parameter already exists")
		return
	} else {
		Params = append(Params, p)
		JsonData, err := json.MarshalIndent(Params, "", "    ")
		if err != nil {
			log.Println(err)
		}

		err = os.WriteFile("parampiper.json", JsonData, 0644)
		if err != nil {
			fmt.Println("error:", err)
		}
	}

}

func (p Parameter) remove() {
	jsonBlob, err := os.ReadFile("parampiper.json")
	if err != nil {
		log.Println(err)
	}

	err = json.Unmarshal(jsonBlob, &Params)
	if err != nil {
		log.Println(err)
	}

	for i, param := range Params {
		if param.Name == p.Name {
			Params = append(Params[:i], Params[i+1:]...)
		}
	}
	JsonData, err := json.MarshalIndent(Params, "", "    ")
	if err != nil {
		log.Println(err)
	}

	err = os.WriteFile("parampiper.json", JsonData, 0644)
	if err != nil {
		fmt.Println("error:", err)
	}
}

func (p Parameter) ifExists() bool {
	jsonBlob, err := os.ReadFile("parampiper.json")
	if err != nil {
		log.Println(err)
	}

	err = json.Unmarshal(jsonBlob, &Params)
	if err != nil {
		log.Println(err)
	}

	for _, param := range Params {
		if param.Name == p.Name {
			return true
		}
	}
	return false
}
