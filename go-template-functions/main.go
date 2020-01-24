package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

type fruit struct {
	Color string
	Name  string
}

//functions passed to template is passed in map form - template.FuncMap
var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

var tpl *template.Template

//local function to be passed to template
func firstThree(s string) string {
	s = strings.TrimSpace(s)
	if len(s) >= 3 {
		s = s[:3]
	}
	return s
}

//initialize template and pass the FuncMap - fm
//Funcs(fm) should be passed before parsing(ParseFiles) the template
//parsing before passing functions will resulting in undefined functions at the template
// ex :
// tpl = template.Must(template.ParseFiles("templatefunctions.gohtml"))
// tpl = tpl.Funcs(fm)
// code above doesnt work - bcs the functions doesnt declared yet when parsing the template
func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("templatefunctions.gohtml"))
}

func main() {
	apple := fruit{
		Name:  "Apple",
		Color: "Red",
	}

	banana := fruit{
		Name:  "Banana",
		Color: "Yellow",
	}

	fruits := []fruit{apple, banana}

	data := struct {
		Fruits []fruit
	}{
		fruits,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "templatefunctions.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}
}
