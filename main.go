package main

import (
	"log"
	"os"

	"github.com/cdalar/parampiper/cmd"

	"github.com/hashicorp/logutils"
)

func init() {
	filter := &logutils.LevelFilter{
		Levels:   []logutils.LogLevel{"DEBUG", "INFO", "WARN", "ERROR"},
		MinLevel: logutils.LogLevel("WARN"),
		Writer:   os.Stderr,
	}
	if os.Getenv("PP_LOG") != "" {
		filter.MinLevel = logutils.LogLevel(os.Getenv("PP_LOG"))
		log.SetFlags(log.Lshortfile)
	} else {
		log.SetFlags(0)
	}
	log.SetOutput(filter)
}

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Println(err)
	}

}
