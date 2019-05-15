package main

import (
	"github.com/golang/go/src/pkg/log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("templates/dataStruct/map/map_2.gohtml"))
}

func main() {

	sages := map[string]string{
		"India":"Gandhi",
		"America":"MLK",
		"Meditate":"Buddha",
		"Love":"Jesus",
		"Prophet":"Muhammad",
	}

	err := tpl.Execute(os.Stdout, sages)
	if err != nil{
		log.Fatalln(err)
	}

}
