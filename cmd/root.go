package cmd

import (
	"log"
	"os"
	"strings"

	"github.com/cdalar/parampiper/internal/data"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "parampiper",
		Short: "parampiper is a tool to manage parameter cross different pipelines",
	}
	provider data.DataProviderInterface
	// dataProvider string
)

// func init() {
// 	rootCmd.Flags().StringVarP(&dataProvider, "data", "d", "local_file", "Data Provider to use")

// }

// Execute executes the root command.
func Execute() error {
	log.Println("[DEBUG] Args: " + strings.Join(os.Args, ","))
	dataProvider := os.Getenv("PP_DATA")
	if dataProvider == "" {
		dataProvider = "local_file"
	}
	switch dataProvider {
	case "local_file":
		provider = &data.LocalFile{
			FilePath: "parampiper.json",
		}
		log.Println("[DEBUG] Using LocalFile")
	case "AzureStorageAccount":
		provider = &data.AzureStorageAccount{
			StorageAccountName: "stparampiper",
			ContainerName:      "sqlite",
		}
		log.Println("[DEBUG] Using AzureStorageAccount")
	}

	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(rmCmd)
}
