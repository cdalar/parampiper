package cmd

import (
	"html/template"
	"log"
	"os"
	"text/tabwriter"
	"time"

	"github.com/spf13/viper"
	"k8s.io/apimachinery/pkg/util/duration"
)

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func ReadConfig(configFilePath string) {
	if !FileExists(configFilePath) {
		log.Println("Configuration File: " + configFilePath + " does not exist")
		os.Exit(1)
	}
	if configFilePath != "" {
		log.Println("[DEBUG] Config Path: " + configFilePath)
		viper.SetConfigFile(configFilePath)
		err := viper.ReadInConfig() // Find and read the config file
		if err != nil {
			log.Println("Problem on ReadInConfig") // Handle errors reading the config file
			log.Println(err)
		}
		err = viper.MergeInConfig()
		if err != nil {
			log.Println(err)
		}
	}

	log.Println("[DEBUG]", viper.AllSettings())

	// dir, err := os.Getwd()
	// if err != nil {
	// 	log.Println(err)
	// }
	// viper.SetConfigName("parampiper")
	// viper.AddConfigPath(dir + "/.pp")
	// err = viper.ReadInConfig()
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// if _, err := os.Stat(dir + "/.pp/" + filename + ".yaml"); err == nil {
	// 	viper.SetConfigName(filename)
	// 	err = viper.MergeInConfig()
	// 	if err != nil {
	// 		log.Println(err)
	// 	}
	// }

	// log.Println("[DEBUG]", viper.AllSettings())
	// // onctlConfig = viper.AllSettings()
}

func durationFromCreatedAt(createdAt time.Time) string {
	return duration.HumanDuration(time.Since(createdAt))
}

func TabWriter(res interface{}, tmpl string) { //nolint
	// var funcs = template.FuncMap{"getNameFromTags": getNameFromTags}
	var funcs2 = template.FuncMap{"durationFromCreatedAt": durationFromCreatedAt}
	w := tabwriter.NewWriter(os.Stdout, 2, 2, 3, ' ', 0)
	// tmp, err := template.New("test").Funcs(funcs).Funcs(funcs2).Parse(tmpl)
	tmp, err := template.New("test").Funcs(funcs2).Parse(tmpl)
	if err != nil {
		log.Println(err)
	}
	err = tmp.Execute(w, res)
	if err != nil {
		log.Println(err)
	}
	w.Flush()
}

func Contains(slice []string, searchValue string) bool {
	for _, value := range slice {
		if value == searchValue {
			return true
		}
	}
	return false
}
