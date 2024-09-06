package main

import (
	_ "embed"
	"io/fs"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"text/template"

	"github.com/iancoleman/strcase"
)

func main() {
	apiDoc, err := LoadHackerRankApiDoc()
	if err != nil {
		panic(err)
	}

	funcMap := template.FuncMap{
		"TrimRef": func(s string) string {
			return strings.TrimPrefix(s, "#/definitions/")
		},
		"Contains": func(e string, s []string) bool {
			return slices.Contains(s, e)
		},
		"ToUpper":  strings.ToUpper,
		"ToPascal": strcase.ToCamel,
	}

	tmpl, err := template.New("api").Funcs(funcMap).ParseFiles(getLanguageTemplates("go")...)
	if err != nil {
		panic(err)
	}

	err = tmpl.ExecuteTemplate(os.Stdout, "api.tmpl", apiDoc)
	if err != nil {
		panic(err)
	}
}

func getLanguageTemplates(language string) []string {
	var filenames []string

	filepath.WalkDir("templates/"+language, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}

		if filepath.Ext(path) != ".tmpl" {
			return nil
		}

		filenames = append(filenames, path)
		return nil
	})

	return filenames
}
