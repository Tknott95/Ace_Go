package sGrid_Ctrl

import (
	"encoding/json"
	"fmt"
	"log"

	"net/http"

	"github.com/julienschmidt/httprouter"
	sendgrid "github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	globals "github.com/tknott95/Ace_Go/Globals"
)

type Email struct {
	SenderEmail string `json:"mail-email"`
	EmailSubj   string `json:"mail-subj"`
	EmailMsg    string `json:"mail-to-trev"`
	EmailAuthor string `json:"mail-author"`
}

func SendEmail(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var emailRes Email
	decoder := json.NewDecoder(req.Body)
	decoder.Decode(&emailRes)

	from := mail.NewEmail("TK - From AceAdmin", emailRes.SenderEmail)
	subject := "From: " + emailRes.EmailAuthor + " Subj: " + emailRes.EmailSubj
	to := mail.NewEmail(emailRes.SenderEmail+" - "+emailRes.EmailAuthor, "tk@trevorknott.io")
	plainTextContent := emailRes.EmailMsg
	htmlContent := "<strong>..." + emailRes.EmailMsg + "</strong>"
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

	w.Header().Set("server", "Sucesfully sent to Trevor!")
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("OK"))

	w.WriteHeader(200)
	// http.Redirect(w, req, "http://trevorknott.io", 301)
}
