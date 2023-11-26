package cmd

import (
	"log"
	"os"
	"strings"

	"github.com/cdalar/parampiper/internal/data"
	"github.com/cdalar/parampiper/internal/secure"
	"github.com/cdalar/parampiper/pkg/common"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	rootCmd = &cobra.Command{
		Use:   "p8r",
		Short: "a tool to manage parameters cross different environments",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			log.Println("[DEBUG] Args: " + strings.Join(os.Args, ","))
			log.Println("[DEBUG] Configuration File: " + configFilePath)
			if len(os.Args) > 1 && os.Args[1] != "init" && os.Args[1] != "version" {
				common.ReadConfig(configFilePath)
			}
			log.Println("[DEBUG]", viper.AllSettings())
			dataProvider, isEnvExits := os.LookupEnv("PP_DATA")
			if isEnvExits {
				checkDataProvider(dataProvider)
			} else {
				dataProvider = "local_file"
			}

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

			switch dataProvider {
			case "local_file":
				provider = &data.LocalFile{
					FilePath: viper.GetString("local_file.filePath"),
				}
				log.Println("[DEBUG] Using LocalFile:", viper.GetString("local_file.filePath"))
			case "azure_blob":
				provider = &data.AzureStorageAccount{
					StorageAccountName: viper.GetString("azure_blob.StorageAccountName"),
					ContainerName:      viper.GetString("azure_blob.ContainerName"),
					BlobName:           viper.GetString("azure_blob.BlobName"),
				}
				log.Println("[DEBUG] Using AzureStorageAccount")
			default:
				log.Println("Provider (" + dataProvider + ") is not Supported\nPlease use one of the following: " + strings.Join(providerList, ","))
				os.Exit(1)
			}
		},
	}
	provider       data.DataProviderInterface
	providerList   = []string{"local_file", "azure_blob"}
	configFilePath string
	encryptionAlgo string
	encrypter      secure.Encrypter
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&configFilePath, "config", "c", ".p8r/parampiper.yaml", "Configuration file")
	rootCmd.PersistentFlags().StringVarP(&encryptionAlgo, "enc", "e", "", "Encryption Algorithm to use: aes, rsa, none")
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func checkDataProvider(dataProvider string) {
	log.Println("[DEBUG] Using: " + dataProvider)
	if dataProvider != "" {
		if !common.Contains(providerList, dataProvider) {
			log.Println("Provider (" + dataProvider + ") is not Supported\nPlease use one of the following: " + strings.Join(providerList, ","))
			os.Exit(1)
		}
	}
}
