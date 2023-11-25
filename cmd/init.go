package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/cdalar/parampiper/internal/files"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "create a .pp directory with the default configuration files",
	Run: func(cmd *cobra.Command, args []string) {
		if _, err := os.Stat(".pp"); os.IsNotExist(err) {
			if err := os.Mkdir(".pp", os.ModePerm); err != nil {
				log.Fatal(err)
			}
			embedDir, err := files.EmbededFiles.ReadDir("init")
			if err != nil {
				log.Fatal(err)
			}
			for _, configFile := range embedDir {
				log.Println("[DEBUG] initFile:" + configFile.Name())
				eFile, _ := files.EmbededFiles.ReadFile("init/" + configFile.Name())
				err = os.WriteFile(".pp/"+configFile.Name(), eFile, 0644)
				if err != nil {
					log.Fatal(err)
				}
			}
			fmt.Println("parampiper environment initialized")
		} else {
			fmt.Println("parampiper environment exists!")
		}
	},
}
