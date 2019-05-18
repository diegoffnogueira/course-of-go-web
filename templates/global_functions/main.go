package main

import (
	"github.com/golang/go/src/pkg/log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("templates/global_functions/global_functions.gohtml"))
}

func main() {

	sages := []string{"Gandhi","MLK","Buddha","Jesus","Muhammad"}

	data := struct {
		Words []string
		LName string
	}{
		sages,
		"Diego",
	}

	err := tpl.Execute(os.Stdout, data)
	if err != nil{
		log.Fatalln(err)
	}

}
