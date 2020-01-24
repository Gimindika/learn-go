package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("template.gohtml"))
}

func main() {
	data := struct {
		A int
		B int
	}{
		5,
		3,
	}

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}

// using builtin functions in template
// ex :
// {{if gt .A .B}}
// A > B
// {{end}}
// gt - greater than is placed after if, followed by the 2 elements compared
// same rule goes for and/or etc, if <operator> .<element1> .<element2>

// complete ref : https://godoc.org/text/template#Functions
