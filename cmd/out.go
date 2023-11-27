package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

func init() {
	outCmd.Flags().StringVarP(&outputType, "output", "o", "", "Output type: raw, json, yaml, table")
	rootCmd.AddCommand(outCmd)
}

var outCmd = &cobra.Command{
	Use:     "out",
	Aliases: []string{"out"},
	Short:   "Output Parameters",
	Run: func(cmd *cobra.Command, args []string) {
		var tmpl string
		log.Println("[DEBUG] Output Parameters")

		parameters, err := provider.Read()
		if err != nil {
			log.Println(err)
		}

		log.Println("[DEBUG]", parameters)
		switch outputType {
		case "tfvars":
			tmpl = "{{range .}}{{.Name}} = \"{{.Value}}\"\n{{end}}"
		default:
			tmpl = "NAME\tVALUE\tINFO\n{{range .}}{{.Name}}\t{{.Value}}\t{{.Info}}\n{{end}}"
		}

		TabWriter(parameters, tmpl)
	},
}
