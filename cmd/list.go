package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:     "ls",
	Aliases: []string{"list"},
	Short:   "List Parameters",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("[DEBUG] List Parameters")

		parameters, err := provider.Read()
		if err != nil {
			log.Println(err)
		}

		log.Println("[DEBUG]", parameters)
		tmpl := "NAME\tVALUE\tINFO\n{{range .}}{{.Name}}\t{{.Value}}\t{{.Info}}\n{{end}}"
		TabWriter(parameters, tmpl)
	},
}