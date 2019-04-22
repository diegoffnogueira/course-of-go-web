package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

func main() {

	//########## One ##########
	tpl, err := template.ParseFiles("templates/hello.gohtml")
	if err != nil{
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, nil)
	if err != nil{
		log.Fatalln(err)
	}

	fmt.Println("")
	fmt.Println("-----------two-------------")
	fmt.Println("")

	//########## two ##########
	nf, err := os.Create("index.html")
	if err != nil{
		log.Println("error creating file", err)
	}
	defer nf.Close()

	err = tpl.Execute(nf, nil)
	if err != nil{
		log.Fatalln(err)
	}

	fmt.Println("")
	fmt.Println("------------three------------")
	fmt.Println("")

	//########## three ##########
	tplFiles, err := template.ParseFiles("files/one.gohtml")
	if err != nil{
		log.Fatalln(err)
	}

	err = tplFiles.Execute(os.Stdout, nil)
	if err != nil{
		log.Fatalln("Passo 1 - ", err)
	}

	tplFiles, err = template.ParseFiles("files/three.gohtml", "files/two.gohtml", "files/one.gohtml")
	if err != nil{
		log.Fatalln("Passo 2 - ", err)
	}

	err = tplFiles.ExecuteTemplate(os.Stdout,"three.gohtml", nil)
	if err != nil{
		log.Fatalln("Passo 3 - ", err)
	}

	err = tplFiles.ExecuteTemplate(os.Stdout,"two.gohtml", nil)
	if err != nil{
		log.Fatalln("Passo 4 - ", err)
	}

	err = tplFiles.ExecuteTemplate(os.Stdout,"one.gohtml", nil)
	if err != nil{
		log.Fatalln("Passo 5 - ", err)
	}

	err = tplFiles.Execute(os.Stdout, nil)
	if err != nil{
		log.Fatalln(err)
	}

	fmt.Println("")
	fmt.Println("------------four------------")
	fmt.Println("")

	//########## four ##########
	tplFilesGlob, err := template.ParseGlob("files/*")
	if err != nil{
		log.Fatalln(err)
	}

	err = tplFilesGlob.Execute(os.Stdout, nil)
	if err != nil{
		log.Fatalln("Passo 1 - ", err)
	}

	err = tplFilesGlob.ExecuteTemplate(os.Stdout,"three.gohtml", nil)
	if err != nil{
		log.Fatalln("Passo 3 - ", err)
	}

	err = tplFilesGlob.ExecuteTemplate(os.Stdout,"two.gohtml", nil)
	if err != nil{
		log.Fatalln("Passo 4 - ", err)
	}

	err = tplFilesGlob.ExecuteTemplate(os.Stdout,"one.gohtml", nil)
	if err != nil{
		log.Fatalln("Passo 5 - ", err)
	}

}

