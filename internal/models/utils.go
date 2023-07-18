package models

import (
	"os"
	"github.com/jedib0t/go-pretty/v6/table"
)

func ConvertInterfaceArrayToStringArray(aInterface []interface{}) []string {
	aString := make([]string, len(aInterface))
	for i, v := range aInterface {
		aString[i] = v.(string)
	}
	return aString
}

func InteractivePrompt() {
	// https://github.com/c-bata/go-prompt
}

func DisplayOutput(headers []string, row []string) {
	t := table.NewWriter()
    t.SetOutputMirror(os.Stdout)
    t.AppendHeader(table.Row{headers})
    t.AppendRow(table.Row{row})
    t.Render()
}

func ObjectAssign(target interface{}, object interface{}) {
	// object atributes values in target atributes values
	// using pattern matching (https://golang.org/pkg/reflect/#Value.FieldByName)
	// https://stackoverflow.com/questions/35590190/how-to-use-the-spread-operator-in-golang
	t := reflect.ValueOf(target).Elem()
	o := reflect.ValueOf(object).Elem()
	for i := 0; i < o.NumField(); i++ {
		for j := 0; j < t.NumField(); j++ {
			if t.Field(j) == o.Field(i) {
				fmt.Printf("Field %s is equal to %s\n", t.Field(j), o.Field(i))
				t.Field(j).Set(o.Field(i))
			}
		}
	}
}
