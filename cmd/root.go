package cmd

import (
	"log"
	"os"
	"strings"

	"github.com/hashicorp/logutils"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "pp",
		Short: "parampiper is a tool to manage parameter cross different pipelines",
	}
)

func init() {
	filter := &logutils.LevelFilter{
		Levels:   []logutils.LogLevel{"DEBUG", "WARN", "ERROR"},
		MinLevel: logutils.LogLevel("WARN"),
		Writer:   os.Stderr,
	}
	if os.Getenv("ONCTL_LOG") != "" {
		filter.MinLevel = logutils.LogLevel(os.Getenv("ONCTL_LOG"))
		log.SetFlags(log.Lshortfile)
	}
	log.SetOutput(filter)

}

// Execute executes the root command.
func Execute() error {
	log.Println("[DEBUG] Args: " + strings.Join(os.Args, ","))
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
