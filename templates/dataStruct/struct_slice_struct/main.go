package main

import (
	"github.com/golang/go/src/pkg/log"
	"os"
	"text/template"
)

var tpl *template.Template

type sage struct {
	Name  string
	Motto string
}

type car struct {
	Manufacter string
	Model      string
	Doors      int
}

type itens struct {
	Wisdom   []sage
	Trasport []car
}

func init() {
	tpl = template.Must(template.ParseFiles("templates/dataStruct/struct_slice_struct/struct_slice_struct.gohtml"))
}

func main() {

	buddha := sage{
		Name:  "Buddha",
		Motto: "The bilief of no beliefs.",
	}

	gandhi := sage{
		Name:  "Gandhi",
		Motto: "Be the change.",
	}

	mlk := sage{
		Name:  "MLK",
		Motto: "Hatred never ceases with hatred but with love alone is healed.",
	}

	jesus := sage{
		Name:  "Jesus",
		Motto: "Love all.",
	}

	muhammad := sage{
		Name:  "Muhammad",
		Motto: "To overcome evil with good is good, to resist evil by evil is evil.",
	}

	ford := car{
		Manufacter:"Ford",
		Model:"F150",
		Doors:2,
	}

	toyota := car{
		Manufacter:"Toyota",
		Model:"Corolla",
		Doors:4,
	}

	sages := []sage{buddha, gandhi, mlk, jesus, muhammad}
	cars := []car{ford, toyota}

	//data := itens{
	//	Wisdom:sages,
	//	Trasport:cars,
	//}

	//Anonimous type
	data := struct {
		Wisdom   []sage
		Trasport []car
	}{
		Wisdom:sages,
		Trasport:cars,
	}

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}

}
