package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

func init() {
	rmCmd.Flags().StringVarP(&param.Name, "name", "n", "", "Name of the parameter")
	err := rmCmd.MarkFlagRequired("name")
	if err != nil {
		log.Println(err)
	}
	rootCmd.AddCommand(rmCmd)
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
		parameters.Remove(param.Name)

		log.Println("[DEBUG] Parameters: ", parameters)
		err = provider.Save(parameters)
		if err != nil {
			log.Println(err)
		}

		log.Println("[DEBUG] Parameters: ", parameters)
		err = provider.Save(parameters)
		if err != nil {
			log.Println(err)
		}
	},
}
