package main

import (
	"flag"
	"log"
	"os"
	"text/template"

	"github.com/mewkiz/pkg/jsonutil"
	"github.com/pkg/errors"
)

func main() {
	flag.Parse()
	for _, jsonPath := range flag.Args() {
		data, err := parseData(jsonPath)
		if err != nil {
			log.Fatalf("%+v", err)
		}
		if err := generateSheet(data); err != nil {
			log.Fatalf("%+v", err)
		}
	}
}

func parseData(jsonPath string) (map[string]interface{}, error) {
	data := make(map[string]interface{})
	if err := jsonutil.ParseFile(jsonPath, &data); err != nil {
		return nil, errors.WithStack(err)
	}
	return data, nil
}

const tmplName = "sheet.tmpl"

func generateSheet(data map[string]interface{}) error {
	t, err := template.ParseFiles(tmplName)
	if err != nil {
		return errors.WithStack(err)
	}
	if err := t.Execute(os.Stdout, data); err != nil {
		return errors.WithStack(err)
	}
	return nil
}
