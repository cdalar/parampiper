package data

import (
	"encoding/json"
	"testing"
)

func TestParameters_Add(t *testing.T) {
	testCases := []struct {
		name     string
		params   Parameters
		param    Parameter
		expected Parameters
	}{
		{
			name:     "Add parameter to empty list",
			params:   Parameters{},
			param:    Parameter{Name: "param1", Value: "value1", Info: "info1"},
			expected: Parameters{{Name: "param1", Value: "value1", Info: "info1"}},
		},
		{
			name:     "Add parameter to non-empty list",
			params:   Parameters{{Name: "param1", Value: "value1", Info: "info1"}},
			param:    Parameter{Name: "param2", Value: "value2", Info: "info2"},
			expected: Parameters{{Name: "param1", Value: "value1", Info: "info1"}, {Name: "param2", Value: "value2", Info: "info2"}},
		},
		{
			name:     "Add parameter without Name to the list",
			params:   Parameters{{Name: "param1", Value: "value1", Info: "info1"}},
			param:    Parameter{Name: "", Value: "value2", Info: "info2"},
			expected: Parameters{{Name: "param1", Value: "value1", Info: "info1"}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			params := tc.params
			param := tc.param

			params.Add(param)

			if len(params) != len(tc.expected) {
				t.Errorf("Expected length of parameters to be %d, got %d", len(tc.expected), len(params))
			}

			for i, expectedParam := range tc.expected {
				if params[i] != expectedParam {
					t.Errorf("Expected parameter at index %d to be %v, got %v", i, expectedParam, params[i])
				}
			}
		})
	}
}

func TestParameters_Remove(t *testing.T) {
	testCases := []struct {
		name     string
		params   Parameters
		param    Parameter
		expected Parameters
	}{
		{
			name:     "Remove parameter from list",
			params:   Parameters{{Name: "param1", Value: "value1", Info: "info1"}},
			param:    Parameter{Name: "param1", Value: "value1", Info: "info1"},
			expected: Parameters{},
		},
		{
			name:     "Remove non-existing parameter from list",
			params:   Parameters{{Name: "param1", Value: "value1", Info: "info1"}},
			param:    Parameter{Name: "param2", Value: "value2", Info: "info2"},
			expected: Parameters{{Name: "param1", Value: "value1", Info: "info1"}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			params := tc.params
			param := tc.param

			params.Remove(param)

			if len(params) != len(tc.expected) {
				t.Errorf("Expected length of parameters to be %d, got %d", len(tc.expected), len(params))
			}

			for i, expectedParam := range tc.expected {
				if params[i] != expectedParam {
					t.Errorf("Expected parameter at index %d to be %v, got %v", i, expectedParam, params[i])
				}
			}
		})
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

	expectedYAML := "name: param1\nvalue: value1\ninfo: info1\n"
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
