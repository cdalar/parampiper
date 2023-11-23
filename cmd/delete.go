package cmd

import (
	"log"

	"github.com/cdalar/parampiper/internal/data"
	"github.com/spf13/cobra"
)

func init() {
	rmCmd.Flags().StringVarP(&param.Name, "name", "n", "", "Name of the parameter")
}

var rmCmd = &cobra.Command{
	Use:     "rm",
	Aliases: []string{"delete", "remove", "rm"},
	Short:   "Delete Parameter",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("[DEBUG] Delete Parameter")

		parameters, err := provider.Read()
		if err != nil {
			log.Println(err)
		}
		p := data.Parameter{
			Name: param.Name,
		}
		log.Println("[DEBUG] Parameter: ", p)
		parameters.Remove(p)

		log.Println("[DEBUG] Parameters: ", parameters)
		err = provider.Save(parameters)
		if err != nil {
			log.Println(err)
		}
	},
}
