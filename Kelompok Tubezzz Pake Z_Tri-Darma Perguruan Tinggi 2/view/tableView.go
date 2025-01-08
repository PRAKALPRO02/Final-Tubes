package view

import (
	"fmt"
	"os"
	"reflect"
	"strings"
	"text/template"
)

type Table struct {
	Headers []string
	Rows    [][]string
	Widths  []int
}

func printTable(data interface{}) {

	table := structToTable(data)

	const tableTemplate = `
{{- /* Print header row */ -}}
+{{range $i, $header := .Headers}}{{repeat "-" (index $.Widths $i)}}+{{end}}
|{{range $i, $header := .Headers}} {{printf "%-*s" (index $.Widths $i) $header}} |{{end}}
+{{range $i, $header := .Headers}}{{repeat "-" (index $.Widths $i)}}+{{end}}
{{- /* Print data rows */ -}}
{{- range .Rows}}
|{{range $i, $value := .}} {{printf "%-*s" (index $.Widths $i) $value}} |{{end}}
+{{range $i, $value := $.Headers}}{{repeat "-" (index $.Widths $i)}}+{{end}}
{{- end}}
`

	funcMap := template.FuncMap{
		"repeat": func(char string, count int) string {
			return strings.Repeat(char, count)
		},
	}

	tmpl, err := template.New("table").Funcs(funcMap).Parse(tableTemplate)
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(os.Stdout, table)
	if err != nil {
		panic(err)
	}
}

func structToTable(data interface{}) Table {
	val := reflect.ValueOf(data)
	if val.Kind() != reflect.Slice {
		panic("structToTable requires a slice of structs")
	}

	elemType := val.Type().Elem()
	if elemType.Kind() != reflect.Struct {
		panic("structToTable requires a slice of structs")
	}

	headers := []string{}
	for i := 0; i < elemType.NumField(); i++ {
		headers = append(headers, elemType.Field(i).Name)
	}

	rows := [][]string{}
	widths := make([]int, len(headers))

	for i := 0; i < val.Len(); i++ {
		row := []string{}
		structVal := val.Index(i)
		for j := 0; j < structVal.NumField(); j++ {
			fieldValue := fmt.Sprintf("%v", structVal.Field(j).Interface())
			row = append(row, fieldValue)

			if len(fieldValue) > widths[j] {
				widths[j] = len(fieldValue)
			}
		}
		rows = append(rows, row)
	}

	for i, header := range headers {
		if len(header) > widths[i] {
			widths[i] = len(header)
		}
	}

	return Table{Headers: headers, Rows: rows, Widths: widths}
}
