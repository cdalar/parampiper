package cmd

import (
	"log"

	"github.com/cdalar/parampiper/internal/data"
	"github.com/spf13/cobra"
)

var (
	param data.Parameter
)

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&param.Name, "name", "n", "", "Name of the parameter")
	err := addCmd.MarkFlagRequired("name")
	if err != nil {
		log.Println(err)
	}
	addCmd.Flags().StringVarP(&param.Value, "value", "v", "", "Value of the parameter")
	addCmd.Flags().StringVarP(&param.Info, "info", "i", "", "Info of the parameter")
}

var addCmd = &cobra.Command{
	Use:     "put",
	Aliases: []string{"add", "put", "set"},
	Short:   "Add/Update Parameter",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("[DEBUG] Add/Update Parameters")
		parameters, err := provider.Read()
		if err != nil {
			log.Println(err)
		}
		log.Println("[DEBUG] Parameter: ", param)
		if encryptionAlgo != "" {
			log.Println("[DEBUG] Encrypting parameters")

			for i, p := range parameters {
				decrypted, err := encrypter.Decrypt(p.Value)
				if err != nil {
					log.Println(err)
				}
				parameters[i].Value = string(decrypted)
			}
		}
		encryptedText, err := encrypter.Encrypt([]byte(param.Value))
		if err != nil {
			log.Println(err)
		}
		param.Value = string(encryptedText)
		parameters.Add(param)
		log.Println("[DEBUG] Parameters: ", parameters)
		err = provider.Save(parameters)
		if err != nil {
			log.Println(err)
		}
	},
}
