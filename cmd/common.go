package cmd

import (
	"html/template"
	"log"
	"os"
	"text/tabwriter"
)

func TabWriter(res interface{}, tmpl string) { //nolint
	// var funcs = template.FuncMap{"getNameFromTags": getNameFromTags}
	// var funcs2 = template.FuncMap{"durationFromCreatedAt": durationFromCreatedAt}
	w := tabwriter.NewWriter(os.Stdout, 2, 2, 3, ' ', 0)
	// tmp, err := template.New("test").Funcs(funcs).Funcs(funcs2).Parse(tmpl)
	tmp, err := template.New("test").Parse(tmpl)
	if err != nil {
		log.Println(err)
	}
	err = tmp.Execute(w, res)
	if err != nil {
		log.Println(err)
	}
	w.Flush()
}
