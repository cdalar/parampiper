package data

import (
	"fmt"
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

func (p *Parameter) String() string {
	return fmt.Sprintf("%s: %s", p.Name, p.Value)
}

func (p *Parameters) Add(prm Parameter) {
	*p = append(*p, prm)
}

func (p *Parameters) Remove(prm Parameter) {
	for i, param := range *p {
		if param.Name == prm.Name {
			*p = append((*p)[:i], (*p)[i+1:]...)
			break
		}
	}
}

// func (p Parameter) IfExists() bool {
// 	for _, param := range Params {
// 		if param.Name == p.Name {
// 			return true
// 		}
// 	}
// 	return false
// }
