package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

var tpl *template.Template

var fm = template.FuncMap{
	"fdateDMY": formatDate,
}

func formatDate(t time.Time) string {
	return t.Format("02-01-2006")
	// 01 for Month
	// 02 for Date
	// 2006 for Year
	// see golang docs on time sections for complete reference
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("template.gohtml"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "template.gohtml", time.Now())
	if err != nil {
		log.Fatalln(err)
	}
}
