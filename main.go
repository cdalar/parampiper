package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

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

func isInputFromPipe() bool {
	fileInfo, _ := os.Stdin.Stat()
	return fileInfo.Mode()&os.ModeCharDevice == 0
}

func toUppercase(r io.Reader, w io.Writer) error {
	scanner := bufio.NewScanner(bufio.NewReader(r))
	for scanner.Scan() {
		_, e := fmt.Fprintln(
			w, strings.ToUpper(scanner.Text()))
		if e != nil {
			return e
		}
	}
	return nil
}
