package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/cdalar/parampiper/internal/data"
	"github.com/cdalar/parampiper/internal/secure"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	rootCmd = &cobra.Command{
		Use:   "parampiper",
		Short: "parampiper is a tool to manage parameters cross different pipelines",
		Run: func(cmd *cobra.Command, args []string) {
			log.Println("[DEBUG] Root Command", encryptionAlgo)
			switch encryptionAlgo {
			case "none":
				log.Println("[DEBUG] Using Encryption None")
				encrypter = &secure.None{}
			case "aes":
				log.Println("[DEBUG] Using Encryption AES")
				encrypter = &secure.AES{}
			case "base64":
				log.Println("[DEBUG] Using Encryption Base64")
				encrypter = &secure.Base64{}
			default:
				log.Println("[DEBUG] Using Default Encryption AES")
				encrypter = &secure.AES{}
			}

		},
	}
	provider       data.DataProviderInterface
	providerList   = []string{"local_file", "azure_blob"}
	encryptionAlgo string
	encrypter      secure.Encrypter
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&encryptionAlgo, "enc", "e", "", "Encryption Algorithm to use: aes, rsa")
}

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
	fmt.Println("Using: " + dataProvider)
	if dataProvider != "" {
		if !Contains(providerList, dataProvider) {
			log.Println("Provider (" + dataProvider + ") is not Supported\nPlease use one of the following: " + strings.Join(providerList, ","))
			os.Exit(1)
		}
	}
}
