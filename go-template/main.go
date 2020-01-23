package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	//execute the 1st template parsed
	err := tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	//execute template by name
	err = tpl.ExecuteTemplate(os.Stdout, "three.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "two.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	//execute template and pass a data
	err = tpl.ExecuteTemplate(os.Stdout, "withdata.gohtml", "i am the data passed to the template")

	//execute template and pass a data with variable
	err = tpl.ExecuteTemplate(os.Stdout, "withvariable.gohtml", "who's that po*emon ? it's variable")
}
