package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

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

	err = json.Unmarshal(jsonBlob, &params)
	if err != nil {
		log.Println(err)
	}
	if p.ifExists() {
		fmt.Println("Parameter already exists")
		return
	} else {
		params = append(params, p)
		JsonData, err := json.MarshalIndent(params, "", "    ")
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

	err = json.Unmarshal(jsonBlob, &params)
	if err != nil {
		log.Println(err)
	}

	for i, param := range params {
		if param.Name == p.Name {
			params = append(params[:i], params[i+1:]...)
		}
	}
	JsonData, err := json.MarshalIndent(params, "", "    ")
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

	err = json.Unmarshal(jsonBlob, &params)
	if err != nil {
		log.Println(err)
	}

	for _, param := range params {
		if param.Name == p.Name {
			return true
		}
	}
	return false
}
