package data

import (
	"log"
	"os"
	"testing"
)

func TestLocalFile_Save(t *testing.T) {
	// Create a temporary file for testing
	tempFile, err := os.CreateTemp("", "testfile")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	// Create a LocalFile instance with the temporary file path
	localFile := LocalFile{FilePath: tempFile.Name()}

	// Define the test parameters
	testParams := Parameters{
		{Name: "param1", Value: "value1", Info: "info1"},
		{Name: "param2", Value: "value2", Info: "info2"},
	}

	// Call the Save method
	err = localFile.Save(ParampiperData{Parameters: testParams})
	if err != nil {
		t.Fatalf("Failed to save parameters to file: %v", err)
	}

	// Read the saved data from the file
	savedParams, err := localFile.Read()
	if err != nil {
		t.Fatalf("Failed to read parameters from file: %v", err)
	}

	// Verify the saved data
	if len(savedParams.Parameters) != len(testParams) {
		t.Errorf("Expected length of saved parameters to be %d, got %d", len(testParams), len(savedParams.Parameters))
	}
	for i, expectedParam := range testParams {
		if savedParams.Parameters[i].Value != expectedParam.Value {
			t.Errorf("Expected saved parameter at index %d to be %v, got %v", i, expectedParam, savedParams.Parameters[i])
		}
	}
}
func TestLocalFile_Read(t *testing.T) {
	// Create a temporary file for testing
	tempFile, err := os.CreateTemp("", "testfile")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	// Write test data to the temporary file
	ppData := ParampiperData{
		Version: DATA_FORMAT_VERSION,
		Parameters: Parameters{
			{Name: "param1", Value: "value1", Info: "info1"},
			{Name: "param2", Value: "value2", Info: "info2"},
		},
	}
	testData := ppData.ToJSON()
	err = os.WriteFile(tempFile.Name(), []byte(testData), 0644)
	if err != nil {
		t.Fatalf("Failed to write test data to file: %v", err)
	}

	// Create a LocalFile instance with the temporary file path
	localFile := LocalFile{FilePath: tempFile.Name()}

	// Call the Read method
	parameters, err := localFile.Read()
	if err != nil {
		t.Fatalf("Failed to read parameters from file: %v", err)
	}

	// Verify the result
	expectedParameters := Parameters{{Name: "param1", Value: "value1", Info: "info1"}, {Name: "param2", Value: "value2", Info: "info2"}}
	if len(parameters.Parameters) != len(expectedParameters) {
		t.Errorf("Expected length of parameters to be %d, got %d", len(expectedParameters), len(parameters.Parameters))
	}
	for i, expectedParam := range expectedParameters {
		if parameters.Parameters[i].Value != expectedParam.Value {
			t.Errorf("Expected parameter at index %d to be %v, got %v", i, expectedParam, parameters.Parameters[i])
		}
	}
}

func TestLocalFile_ReadNoFile(t *testing.T) {
	// Create a LocalFile instance with the temporary file path
	localFile := LocalFile{FilePath: "nonexistentfile"}

	// Call the Read method
	parameters, err := localFile.Read()
	if err != nil {
		t.Fatalf("Failed to read parameters from file: %v", err)
	}

	// Verify the result
	expectedParameters := Parameters{}
	if len(parameters.Parameters) != len(expectedParameters) {
		t.Errorf("Expected length of parameters to be %d, got %d", len(expectedParameters), len(parameters.Parameters))
	}
	for i, expectedParam := range expectedParameters {
		if parameters.Parameters[i].Value != expectedParam.Value {
			t.Errorf("Expected parameter at index %d to be %v, got %v", i, expectedParam, parameters.Parameters[i])
		}
	}
	err = os.Remove("nonexistentfile")
	if err != nil {
		log.Println(err)
	}
}

func TestLocalFile_ReadBadJson(t *testing.T) {
	// Create a LocalFile instance with the temporary file path
	localFile := LocalFile{FilePath: "nonexistentfile"}

	// Call the Read method
	_, err := localFile.Read()
	if err != nil {
		log.Println(err)
	}
}
