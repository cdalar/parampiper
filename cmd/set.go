package cmd

import (
	"encoding/json"
	"log"

	"github.com/cdalar/parampiper/internal/data"
	"github.com/cdalar/parampiper/pkg/common"
	"github.com/spf13/cobra"
)

var (
	param data.Parameter
)

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&param.Name, "name", "n", "", "Name of the parameter")
	err := addCmd.MarkFlagRequired("name")
	if err != nil {
		log.Println(err)
	}
	addCmd.Flags().StringVarP(&param.Value, "value", "v", "", "Value of the parameter")
	addCmd.Flags().StringVarP(&attributes, "attr", "a", "", "Parameter Attributes")
	addCmd.Flags().StringVarP(&param.Info, "info", "i", "", "Info of the parameter")
}

var addCmd = &cobra.Command{
	Use:     "set",
	Aliases: []string{"put", "add"},
	Short:   "Add/Update Parameter",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("[DEBUG] Add/Update Parameters")
		unsupportChars := []rune{'-', '*', ',', '.', ':', ';', '<', '>', '?', '\\', '|', ' '}
		if common.ContainsAny(param.Name, unsupportChars) {
			log.Println("[ERROR] Parameter name cannot contain any of the following characters:", string(unsupportChars))
			return
		}
		readData, err := provider.Read()
		if err != nil {
			log.Println(err)
		}
		if attributes != "" {
			log.Println("[DEBUG] Attributes: ", attributes)
			param.Attributes = make(map[string]interface{})
			err = json.Unmarshal([]byte(attributes), &param.Attributes)
			if err != nil {
				log.Println(err)
			}
		}
		for _, parameter := range readData.Parameters {
			if parameter.Name == param.Name {
				if param.Value == "" {
					param.Value = parameter.Value
				}
				if param.Info == "" {
					param.Info = parameter.Info
				}
				if param.Type == "" {
					param.Type = parameter.Type
				}
				if param.Attributes == nil {
					param.Attributes = parameter.Attributes
				}
			}
		}

		param.Type = "basic"

		log.Println("[DEBUG] Parameter: ", param)
		readData.Parameters.Add(param)
		log.Println("[DEBUG] Parameters: ", readData.Parameters)
		err = provider.Save(readData)
		if err != nil {
			log.Println(err)
		}
	},
}
