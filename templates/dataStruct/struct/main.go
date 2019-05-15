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
	tpl = template.Must(template.ParseFiles("templates/dataStruct/struct/struct_2.gohtml"))
}

func main() {

	buddha := sage{
		Name:"Buddha",
		Motto:"The bilief of no beliefs",
	}

	err := tpl.Execute(os.Stdout, buddha)
	if err != nil{
		log.Fatalln(err)
	}

}
