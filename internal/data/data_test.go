package data

import (
	"encoding/json"
	"testing"
)

func TestParameters_Add(t *testing.T) {
	params := Parameters{}
	param := Parameter{Name: "param1", Value: "value1", Info: "info1"}

	params.Add(param)

	if len(params) != 1 {
		t.Errorf("Expected length of parameters to be 1, got %d", len(params))
	}

	if params[0] != param {
		t.Errorf("Expected parameter to be %v, got %v", param, params[0])
	}
}

func TestParameters_Remove(t *testing.T) {
	param := Parameter{Name: "param1", Value: "value1", Info: "info1"}
	params := Parameters{param}

	params.Remove(param)

	if len(params) != 0 {
		t.Errorf("Expected length of parameters to be 0, got %d", len(params))
	}
}

func TestParameters_IfExists(t *testing.T) {
	param := Parameter{Name: "param1", Value: "value1", Info: "info1"}
	params := Parameters{param}

	exists, index := params.IfExists(param)

	if !exists {
		t.Errorf("Expected parameter to exist, got false")
	}

	if index != 0 {
		t.Errorf("Expected index to be 0, got %d", index)
	}
}

func TestParameter_ToJSON(t *testing.T) {
	param := Parameter{Name: "param1", Value: "value1", Info: "info1"}

	jsonStr := param.ToJSON()

	expectedJSON, err := json.MarshalIndent(param, "", "    ")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if jsonStr != string(expectedJSON) {
		t.Errorf("Expected JSON string to be %s, got %s", expectedJSON, jsonStr)
	}
}

func TestParameter_ToYAML(t *testing.T) {
	param := Parameter{Name: "param1", Value: "value1", Info: "info1"}

	yamlStr := param.ToYAML()

	expectedYAML := "name: param1\nvalue: value1\ninfo: info1\ntime: \"\"\n"
	if yamlStr != expectedYAML {
		t.Errorf("Expected YAML string to be %s, got %s", expectedYAML, yamlStr)
	}
}

func TestParameter_String(t *testing.T) {
	param := Parameter{Name: "param1", Value: "value1", Info: "info1"}

	str := param.String()

	expectedStr := "param1: value1 (info1)"
	if str != expectedStr {
		t.Errorf("Expected string representation to be %s, got %s", expectedStr, str)
	}
}
