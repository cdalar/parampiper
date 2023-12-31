package data

import (
	"encoding/json"
	"fmt"
	"log"
	"sort"
	"strings"

	"gopkg.in/yaml.v2"
)

type DataProviderInterface interface {
	Read() (Parameters, error)
	Save(Parameters) error
}
type Parameter struct {
	Name       string                 `json:"name"`
	Type       string                 `json:"type"`
	Value      string                 `json:"value"`
	Attributes map[string]interface{} `json:"attributes"`
	Info       string                 `json:"info"`
}

type Parameters []Parameter

func (p *Parameters) Add(prm Parameter) {
	if prm.Name == "" {
		log.Println("[DEBUG] Parameter name is empty")
		return
	}
	isExists, paramPos := p.IfExists(prm)
	if isExists {
		if prm.Value == "" {
			log.Println("[DEBUG] Parameter value is empty")
			log.Println("[DEBUG] Keep Parameter Value: ", (*p)[paramPos].Value)
			prm.Value = (*p)[paramPos].Value
		}
		if prm.Info == "" {
			log.Println("[DEBUG] Parameter info is empty")
			log.Println("[DEBUG] Keep Parameter Info: ", (*p)[paramPos].Info)
			prm.Info = (*p)[paramPos].Info
		}
		(*p)[paramPos] = prm
	} else {
		*p = append(*p, prm)
	}
	// Sort the parameters by Name
	sort.Slice(*p, func(i, j int) bool {
		return (*p)[i].Name < (*p)[j].Name
	})
}

func (p *Parameters) Remove(parameterList string) {
	parameters := *p
	prm := strings.Split(parameterList, ",")
	for i := len(parameters) - 1; i >= 0; i-- {
		for _, parameter := range prm {
			if (*p)[i].Name == parameter {
				*p = append((*p)[:i], (*p)[i+1:]...)
			}
		}
	}
}

func (p *Parameters) Filter(filterList string) Parameters {
	var filtered Parameters
	sliceList := strings.Split(filterList, ",")
	for _, param := range *p {
		for _, filter := range sliceList {
			if param.Name == filter {
				log.Println("[DEBUG] Found parameter: ", param.Name)
				filtered = append(filtered, param)
			}
		}
	}
	return filtered
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
