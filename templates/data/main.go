package main

import (
	"log"
	"os"
	"text/template"
)

var (
	tpl *template.Template
	tplVariable *template.Template
)

func init() {
	tpl = template.Must(template.ParseFiles("files/data.gohtml"))
	tplVariable = template.Must(template.ParseFiles("files/dataVariable.gohtml"))
}

func main() {

	err := tpl.ExecuteTemplate(os.Stdout, "data.gohtml", 31)
	if err != nil{
		log.Fatalln(err)
	}

	err = tplVariable.ExecuteTemplate(os.Stdout, "dataVariable.gohtml", "Essa é a idade da vida")
	if err != nil{
		log.Fatalln(err)
	}

}
