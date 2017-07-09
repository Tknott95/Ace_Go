package sGrid_Ctrl

import (
	"fmt"
	"log"
	"os"

	"net/http"

	"github.com/julienschmidt/httprouter"
	sendgrid "github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func SendEmail(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	var mailToSend string

	mailToSend = req.FormValue("mail-to-trev")

	from := mail.NewEmail("TK - From AceAdmin", "tknott95@hotmail.com")
	subject := "Sent From AceAdmin via. Sendgrid"
	to := mail.NewEmail("TK - From AceAdmin", "tk@trevorknott.io")
	plainTextContent := mailToSend
	htmlContent := "<strong>..." + mailToSend + "</strong>"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}
