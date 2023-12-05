package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/cdalar/parampiper/internal/data"
	"github.com/cdalar/parampiper/pkg/common"
	tfjson "github.com/hashicorp/terraform-json"
	"github.com/spf13/cobra"
)

var (
	importFile string
	tfstate    tfjson.State
)

func init() {
	rootCmd.AddCommand(importCmd)
	// importCmd.Flags().StringVarP(&outputType, "tfshowjson", "", "", "Output type: tfvars, ")
	importCmd.Flags().StringVarP(&importFile, "tfshowjson", "f", "", "Import parameters from command: terraform show -json")
	err := importCmd.MarkFlagRequired("tfshowjson")
	if err != nil {
		log.Println(err)
	}

}

var importCmd = &cobra.Command{
	Use:   "import",
	Short: "Import Parameters",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("[DEBUG] Import Parameters")
		parameters, err := provider.Read()
		if err != nil {
			log.Println(err)
		}
		log.Println("[DEBUG] Parameter: ", param)
		log.Println("[DEBUG] Parameters: ", importFile)
		tfstate, err = read(importFile)
		if err != nil {
			log.Println(err)
		}
		for _, v := range tfstate.Values.RootModule.Resources {
			log.Println(v.Type, v.Name, v.AttributeValues["id"])
			param := data.Parameter{
				//  TODO: Support for same name resources
				Name:  v.Type + "__" + v.Name,
				Value: v.AttributeValues["id"].(string),
			}
			parameters.Add(param)
		}

		err = provider.Save(parameters)
		if err != nil {
			log.Println(err)
		}
	},
}

func read(filePath string) (tfjson.State, error) {
	if !common.FileExists(filePath) {
		return tfjson.State{}, fmt.Errorf("File does not exist")
	}
	jsonBlob, err := os.ReadFile(filePath)
	if err != nil {
		log.Println(filePath, "file does not exist")
	}
	err = json.Unmarshal(jsonBlob, &tfstate)
	if err != nil {
		log.Println(err)
	}
	return tfstate, nil

}
