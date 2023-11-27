package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(importCmd)
}

var importCmd = &cobra.Command{
	Use:     "import",
	Aliases: []string{"import"},
	Short:   "Import Parameters",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("[DEBUG] Import Parameters")
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
