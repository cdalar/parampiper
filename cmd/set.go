package cmd

import (
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
		parameters, err := provider.Read()
		if err != nil {
			log.Println(err)
		}
		log.Println("[DEBUG] Parameter: ", param)
		parameters.Add(param)
		log.Println("[DEBUG] Parameters: ", parameters)
		err = provider.Save(parameters)
		if err != nil {
			log.Println(err)
		}
	},
}
