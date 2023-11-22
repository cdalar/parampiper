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
		MinLevel: logutils.LogLevel("DEBUG"),
		Writer:   os.Stderr,
	}
	if os.Getenv("PP_LOG") != "" {
		filter.MinLevel = logutils.LogLevel(os.Getenv("PP_LOG"))
		log.SetFlags(log.Lshortfile)
	} else {
		filter.MinLevel = logutils.LogLevel("DEBUG")
		log.SetFlags(log.Lshortfile)
	}
	log.SetOutput(filter)
}

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Println(err)
	}

	// params.read()
	// jsonBlob, err := os.ReadFile("parampiper.json")
	// handleError(err)
	// err = json.Unmarshal(jsonBlob, &params)
	// if err != nil {
	// 	fmt.Println("error:", err)
	// }
	// fmt.Print(params)
	// Parameter{Name: "testName2233", Value: "testValue"}.add()
	// Parameter{Name: "testName"}.remove()
	// params.save()
}
