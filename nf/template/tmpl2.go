package main

import (
	"fmt"
	"os"
	"text/template"
)

type Friend struct {
	Fname string
}

type Person struct {
	UserName string
	Emails   []string
	Friends  []*Friend
}

func main() {
	//
	// f1 := Friend{Fname: "minux.ma"}
	// f2 := Friend{Fname: "xushiwei"}
	// p := Person{UserName: "Astaxie",
	// 	Emails:  []string{"astaxie@beego.me", "astaxie@gmail.com"},
	// 	Friends: []*Friend{&f1, &f2}}
	//
	s1, _ := template.ParseFiles("header.tmpl", "content.tmpl", "footer.tmpl")
	s1.ExecuteTemplate(os.Stdout, "header", nil)
	fmt.Println()
	s1.ExecuteTemplate(os.Stdout, "content", Person{UserName: "username", Emails: []string{"astaxie@beego.me", "astaxie@gmail.com"}, Friends: []*Friend{&Friend{"001"}}})
	fmt.Println()
	s1.ExecuteTemplate(os.Stdout, "footer", nil)
	fmt.Println()
	s1.Execute(os.Stdout, nil)
}
