package cmd

import (
	"log"

	"github.com/cdalar/parampiper/internal/data"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:     "ls",
	Aliases: []string{"list"},
	Short:   "List Parameters",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("[DEBUG] List Parameters")

		data.AzureStorageAccount{
			StorageAccountName: "stparampiper",
			ContainerName:      "sqlite",
		}.Read()

		log.Println("[DEBUG]", data.Params)
		tmpl := "NAME\tVALUE\tINFO\n{{range .}}{{.Name}}\t{{.Value}}\t{{.Info}}\n{{end}}"
		TabWriter(data.Params, tmpl)
	},
}
