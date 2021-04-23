package helpers

import (
	"bytes"
	gomail "gopkg.in/gomail.v2"
	"html/template"
	"log"
)

type Info struct {
	Name string
}

func (i Info) SendEMail() {

	//t := templates.New("templates")
	//
	//var err error
	t, err := template.ParseFiles("templates/trial.html")
	println(t)
	if err != nil {
		log.Println(err)
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, i); err != nil {
		log.Println(err)
	}

	result := tpl.String()
	m := gomail.NewMessage()
	m.SetHeader("From", "jnyakush99@gmail.com")
	m.SetHeader("To", "welcome@gmail.com")
	m.SetHeader("Subject", "golang test")
	m.SetBody("text/html", result)
	m.Attach("templates/trial.html") // attach whatever you want

	d := gomail.NewDialer("smtp.mailtrap.io", 2525, "aa9386539a6a4a", "14626f11d10aa5")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
