package main

import (
	"github.com/golang/go/src/pkg/log"
	"os"
	"text/template"
)

var tpl *template.Template

type sage struct {
	Name string
	Motto string
}

func init() {
	tpl = template.Must(template.ParseFiles("templates/dataStruct/slice_struct/slice_struct.gohtml"))
}

func main() {

	buddha := sage{
		Name:"Buddha",
		Motto:"The bilief of no beliefs.",
	}

	gandhi := sage{
		Name:"Gandhi",
		Motto:"Be the change.",
	}

	mlk := sage{
		Name:"MLK",
		Motto:"Hatred never ceases with hatred but with love alone is healed.",
	}

	jesus := sage{
		Name:"Jesus",
		Motto:"Love all.",
	}

	muhammad := sage{
		Name:"Muhammad",
		Motto:"To overcome evil with good is good, to resist evil by evil is evil.",
	}

	sages := []sage{buddha, gandhi, mlk, jesus, muhammad}

	err := tpl.Execute(os.Stdout, sages)
	if err != nil{
		log.Fatalln(err)
	}

}
