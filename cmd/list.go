package cmd

import (
	"log"

	"github.com/cdalar/parampiper/pkg/common"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:     "ls",
	Aliases: []string{"list"},
	Short:   "List Parameters",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("[DEBUG] List Parameters")
		readData, err := provider.Read()
		if err != nil {
			log.Println(err)
		}
		log.Println("[DEBUG]", readData)
		tmpl := "NAME\tTYPE\tVALUE\tATTRIBUTES\tINFO\n{{range .}}{{.Name}}\t{{.Type}}\t{{.Value | shorter }}\t{{.Attributes | count}}\t{{.Info}}\n{{end}}"
		common.TabWriter(readData.Parameters, tmpl)
	},
}
