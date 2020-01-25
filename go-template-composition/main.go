package main

import (
	"html/template"
	"log"
	"os"
)

type person struct {
	Name string
	Age  int
}

type police struct {
	person
	License bool
}

func (p police) Shoot() string {
	return "BANG!BANG!!"
}

func (p police) Fine() int {
	return 100
}

func (p police) DoubleFine(fine int) int {
	return fine * 2
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("template.html"))
}

func main() {
	John := police{
		person{
			"John",
			28,
		},
		true,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "template.html", John)
	if err != nil {
		log.Fatalln(err)
	}
}
