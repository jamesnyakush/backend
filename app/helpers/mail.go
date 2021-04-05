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

	t := template.New("template.html")

	var err error
	t, err = t.ParseFiles("template.html")
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
	//m.SetAddressHeader("Cc", "<RECIPIENT CC>", "<RECIPIENT CC NAME>")
	m.SetHeader("Subject", "golang test")
	m.SetBody("text/html", result)
	m.Attach("template.html") // attach whatever you want

	d := gomail.NewDialer("smtp.mailtrap.io", 2525, "aa9386539a6a4a", "14626f11d10aa5")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
