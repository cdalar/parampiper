package cmd

import (
	"log"
	"os"
	"strings"

	"github.com/cdalar/parampiper/internal/data"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	rootCmd = &cobra.Command{
		Use:   "parampiper",
		Short: "parampiper is a tool to manage parameters cross different pipelines",
	}
	provider     data.DataProviderInterface
	providerList = []string{"local_file", "azure_blob"}
	// dataProvider string
)

// func init() {
// 	rootCmd.Flags().StringVarP(&dataProvider, "data", "d", "local_file", "Data Provider to use")

// }

// Execute executes the root command.
func Execute() error {
	ReadConfig("parampiper")
	log.Println("[DEBUG] Args: " + strings.Join(os.Args, ","))
	dataProvider := os.Getenv("PP_DATA")
	if dataProvider == "" {
		dataProvider = "local_file"
	} else {
		checkDataProvider()
	}
	switch dataProvider {
	case "local_file":
		provider = &data.LocalFile{
			FilePath: viper.GetString("local_file.filePath"),
		}
		log.Println("[DEBUG] Using LocalFile")
	case "azure_blob":
		provider = &data.AzureStorageAccount{
			StorageAccountName: viper.GetString("azure_blob.StorageAccountName"),
			ContainerName:      viper.GetString("azure_blob.ContainerName"),
			BlobName:           viper.GetString("azure_blob.BlobName"),
		}
		log.Println("[DEBUG] Using AzureStorageAccount")
	}

	return rootCmd.Execute()
}

func checkDataProvider() {
	dataProvider := os.Getenv("PP_DATA")
	log.Println("[DEBUG] Using: " + dataProvider)
	if dataProvider != "" {
		if !Contains(providerList, dataProvider) {
			log.Println("Provider (" + dataProvider + ") is not Supported\nPlease use one of the following: " + strings.Join(providerList, ","))
			os.Exit(1)
		}
	}
}
