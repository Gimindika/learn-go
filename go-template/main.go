package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

type fruit struct {
	Color string
	Name  string
}

type animal struct {
	Name string
	Food string
}

var tpl *template.Template

// initialize - parsing templates from "templates" folder
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

	//////////////////////////linebreak///////////////////////////////
	fmt.Println("")
	fmt.Println("*************************************************")
	fmt.Println("")
	/////////////////////////////////////////////////////////////////

	//execute template and pass slice as data and range through it
	fruits := []string{"Apple", "Orange", "Dragonfruit"}
	err = tpl.ExecuteTemplate(os.Stdout, "sliceormap.gohtml", fruits)
	if err != nil {
		log.Fatalln(err)
	}

	//////////////////////////linebreak///////////////////////////////
	fmt.Println("")
	fmt.Println("*************************************************")
	fmt.Println("")
	/////////////////////////////////////////////////////////////////

	//execute template and pass slice as data with variable
	err = tpl.ExecuteTemplate(os.Stdout, "slicevariable.gohtml", fruits)
	if err != nil {
		log.Fatalln(err)
	}

	//////////////////////////linebreak///////////////////////////////
	fmt.Println("")
	fmt.Println("*************************************************")
	fmt.Println("")
	/////////////////////////////////////////////////////////////////

	//execute template and pass map as data and range through it
	fm := map[string]string{
		"red":    "apple",
		"orange": "orange",
		"pink":   "dragonfruit",
	}
	err = tpl.ExecuteTemplate(os.Stdout, "sliceormap.gohtml", fm)
	if err != nil {
		log.Fatalln(err)
	}

	//////////////////////////linebreak///////////////////////////////
	fmt.Println("")
	fmt.Println("*************************************************")
	fmt.Println("")
	/////////////////////////////////////////////////////////////////

	//execute template and pass map as data and range through it with key and value
	err = tpl.ExecuteTemplate(os.Stdout, "mapkv.gohtml", fm)
	if err != nil {
		log.Fatalln(err)
	}

	//////////////////////////linebreak///////////////////////////////
	fmt.Println("")
	fmt.Println("*************************************************")
	fmt.Println("")
	/////////////////////////////////////////////////////////////////

	//execute template and pass struct as data
	apple := fruit{
		Color: "Red",
		Name:  "Apple",
	}
	err = tpl.ExecuteTemplate(os.Stdout, "structdata.gohtml", apple)
	if err != nil {
		log.Fatalln(err)
	}

	//////////////////////////linebreak///////////////////////////////
	fmt.Println("")
	fmt.Println("*************************************************")
	fmt.Println("")
	/////////////////////////////////////////////////////////////////

	//execute template and pass slice of structs as data
	orange := fruit{
		Color: "Orange",
		Name:  "Orange",
	}
	dragonfruit := fruit{
		Color: "Pink",
		Name:  "Dragonfruit",
	}
	fx := []fruit{apple, orange, dragonfruit}
	err = tpl.ExecuteTemplate(os.Stdout, "sliceofstruct.gohtml", fx)
	if err != nil {
		log.Fatalln(err)
	}

	//////////////////////////linebreak///////////////////////////////
	fmt.Println("")
	fmt.Println("*************************************************")
	fmt.Println("")
	/////////////////////////////////////////////////////////////////

	//execute template and pass struct with slices of structs as data
	lion := animal{
		Name: "Lion",
		Food: "Meat",
	}
	girrafe := animal{
		Name: "Girrafe",
		Food: "Leaf",
	}
	ax := []animal{lion, girrafe}

	//defining anonymous literal struct
	data := struct {
		Fruits  []fruit
		Animals []animal
	}{
		fx,
		ax,
	}

	err = tpl.ExecuteTemplate(os.Stdout, "structslicestruct.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}
}
