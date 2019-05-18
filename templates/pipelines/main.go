package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

var tpl *template.Template

var fm = template.FuncMap{
	"fDateMDY": monthDayYear,
	"fDateMYD": monthYearDay,
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("templates/pipelines/pipeline.gohtml"))
}

func monthDayYear(t time.Time) string {
	return t.Format(time.Kitchen)
}

func monthYearDay(t time.Time) string {
	return t.Format(time.Stamp)
}

func main() {

	err := tpl.ExecuteTemplate(os.Stdout, "pipeline.gohtml", time.Now())
	if err != nil {
		log.Fatalln(err)
	}
	
}
