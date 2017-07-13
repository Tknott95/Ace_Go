package sGrid_Ctrl

import (
	"fmt"
	"log"

	"net/http"

	"github.com/julienschmidt/httprouter"
	sendgrid "github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	globals "github.com/tknott95/Ace_Go/Globals"
)

func SendEmail(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	var fromWho string
	var subjOfMail string
	var mailToSend string
	var nameOfSender string

	fromWho = req.FormValue("mail-from")
	subjOfMail = req.FormValue("mail-subj")
	mailToSend = req.FormValue("mail-to-trev")
	nameOfSender = req.FormValue("mail-author")

	if subjOfMail == "" {
		subjOfMail = "Sent From AceAdmin via. Sendgrid"
	}

	if nameOfSender == "" {
		nameOfSender = "Hitler(ANON)"
	}

	if fromWho == "" {
		fromWho = "Anonymous@trevorknott.io"
	}

	from := mail.NewEmail("TK - From AceAdmin", fromWho)
	subject := "From: " + nameOfSender + " Subj: " + subjOfMail
	to := mail.NewEmail(fromWho+" - "+nameOfSender, "tk@trevorknott.io")
	plainTextContent := mailToSend
	htmlContent := "<strong>..." + mailToSend + "</strong>"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(globals.SGridApi)
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		println("Response StatusCode")
		fmt.Println(response.StatusCode)
		println("Response Body")
		fmt.Println(response.Body)
		println("Response Headers")
		fmt.Println(response.Headers)
	}

	http.Redirect(w, req, "/", 301)
}
