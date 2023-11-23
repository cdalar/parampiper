package cmd

import (
	"log"

	"github.com/cdalar/parampiper/internal/data"
	"github.com/spf13/cobra"
)

var (
	param data.Parameter
)

func init() {
	addCmd.Flags().StringVarP(&param.Name, "name", "n", "", "Name of the parameter")
	addCmd.Flags().StringVarP(&param.Value, "value", "v", "", "Value of the parameter")
	addCmd.Flags().StringVarP(&param.Info, "info", "i", "", "Info of the parameter")

}

var addCmd = &cobra.Command{
	Use:     "put",
	Aliases: []string{"add", "put"},
	Short:   "Add/Update Parameter",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("[DEBUG] Add/Update Parameters")

		parameters, err := provider.Read()
		if err != nil {
			log.Println(err)
		}
		p := data.Parameter{
			Name:  param.Name,
			Value: param.Value,
			Info:  param.Info,
		}
		log.Println("[DEBUG] Parameter: ", p)
		parameters.Add(p)

		log.Println("[DEBUG] Parameters: ", parameters)
		provider.Save(parameters)
	},
}
