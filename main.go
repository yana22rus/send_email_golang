package main

import (
	"net/http"
	"net/smtp"
	"os"
	"text/template"
)

func main() {

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8000", nil)

}

func send_email(input_email string) {

	from := os.Getenv("email")
	pswd := os.Getenv("pswd")
	to := []string{input_email}
	host := "smtp.gmail.com"
	port := "587"
	address := host + ":" + port
	subject := "Golang\n"
	body := "INPUT"
	message := []byte(subject + body)
	auth := smtp.PlainAuth("", from, pswd, host)
	smtp.SendMail(address, auth, from, to, message)

}

func handler(w http.ResponseWriter, r *http.Request) {

	type input_email struct {
		Email string
	}

	data := input_email{
		Email: r.FormValue("email"),
	}

	tpl := template.Must(template.ParseGlob("templates/*.gohtml"))

	tpl.Execute(w, data)

	if r.Method == "POST" {

		input := r.FormValue("email")

		send_email(input)

	}

}
