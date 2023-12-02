package cmd

import (
	"log"

	"github.com/cdalar/parampiper/pkg/common"
	"github.com/spf13/cobra"
)

var parameterList string

func init() {
	outCmd.Flags().StringVarP(&outputType, "output", "o", "", "Output type: tfvars, ")
	outCmd.Flags().StringVarP(&parameterList, "list", "l", "", "List of parameters to output (comma separated)")
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
		log.Println("[DEBUG] Parameter List: ", parameterList)
		if parameterList != "" {
			parameters = parameters.Filter(parameterList)
		}

		log.Println("[DEBUG]", parameters)
		switch outputType {
		case "tfvars":
			tmpl = "{{range .}}{{.Name}} = \"{{.Value}}\"\n{{end}}"
		case "export":
			tmpl = "{{range .}}export {{.Name | upperCase}}=\"{{.Value}}\"\n{{end}}"
		default:
			tmpl = "NAME\tVALUE\tINFO\n{{range .}}{{.Name}}\t{{.Value}}\t{{.Info}}\n{{end}}"
		}

		common.TabWriter(parameters, tmpl)
	},
}
