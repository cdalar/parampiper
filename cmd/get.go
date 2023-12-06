package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var outputType string

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.Flags().StringVarP(&param.Name, "name", "n", "", "Name of the parameter")
	// getCmd.MarkFlagRequired("name")
	getCmd.Flags().StringVarP(&outputType, "output", "o", "", "Output type: raw, json, yaml, table")
}

var getCmd = &cobra.Command{
	Use:     "get",
	Aliases: []string{"get"},
	Short:   "Get Parameter Value by Name",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("[DEBUG] Get Parameter by Name")
		if param.Name == "" {
			if len(args) == 0 {
				log.Println("[ERROR] Parameter name is required")
				return
			} else if len(args) == 1 {
				log.Println("[DEBUG] Parameter name: ", args[0])
				param.Name = args[0]
			}
		} else {
			log.Println("[DEBUG] Parameter name: ", param.Name)
		}

		parameters, err := provider.Read()
		if err != nil {
			log.Println(err)
		}
		log.Println("[DEBUG] Searching for parameter named: ", param.Name)
		for _, p := range parameters {
			if p.Name == param.Name {
				log.Println("[DEBUG] Found parameter: ", p)
				log.Println("[DEBUG] Value: ", p.Value)
				switch outputType {
				case "raw":
					fmt.Print(p.Value, "\n")
				case "json":
					fmt.Print(p.ToJSON())
				case "yaml":
					fmt.Print(p.ToYAML())
				default:
					fmt.Print(p.Value, "\n")
				}
			}
		}

		// log.Println("[DEBUG]", parameters)
		// tmpl := "NAME\tVALUE\tINFO\n{{range .}}{{.Name}}\t{{.Value}}\t{{.Info}}\n{{end}}"
		// TabWriter(parameters, tmpl)
	},
}
