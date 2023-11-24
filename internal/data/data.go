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

type Parameters []Parameter

type Parameter struct {
	Name  string `json:"name"`
	Value string `json:"value"`
	Info  string `json:"info"`
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
	return fmt.Sprintf("%s: %s", p.Name, p.Value)
}

func (p *Parameters) Add(prm Parameter) {
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
