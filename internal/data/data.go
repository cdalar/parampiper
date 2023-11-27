package data

import (
	"encoding/json"
	"fmt"
	"log"

	"gopkg.in/yaml.v2"
)

type DataProviderInterface interface {
	Read() (Parameters, error)
	Save(Parameters) error
}
type Parameter struct {
	Name  string `json:"name"`
	Value string `json:"value"`
	Info  string `json:"info"`
	// Time  string `json:"time"`
}

type Parameters []Parameter

func (p *Parameters) Add(prm Parameter) {
	if prm.Name == "" {
		log.Println("[DEBUG] Parameter name is empty")
		return
	}
	isExists, paramPos := p.IfExists(prm)
	if isExists {
		(*p)[paramPos] = prm
		return
	} else {
		*p = append(*p, prm)
	}
}

func (p *Parameters) Remove(prm Parameter) {
	for i, param := range *p {
		if param.Name == prm.Name {
			*p = append((*p)[:i], (*p)[i+1:]...)
			break
		}
	}
}

func (p *Parameters) IfExists(prm Parameter) (bool, int) {
	for i, param := range *p {
		if param.Name == prm.Name {
			return true, i
		}
	}
	return false, -1
}

func (p *Parameter) ToJSON() string {
	jsonData, err := json.MarshalIndent(p, "", "    ")
	if err != nil {
		log.Println(err)
	}
	return string(jsonData)
}

func (p *Parameter) ToYAML() string {
	yamlData, err := yaml.Marshal(p)
	if err != nil {
		log.Println(err)
	}
	return string(yamlData)
}

func (p *Parameter) String() string {
	return fmt.Sprintf("%s: %s (%s)", p.Name, p.Value, p.Info)
}
