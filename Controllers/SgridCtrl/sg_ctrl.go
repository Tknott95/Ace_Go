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
	var fromWho string
	var subjOfMail string
	var mailToSend string

	fromWho = req.FormValue("mail-from")
	subjOfMail = req.FormValue("mail-subj")
	mailToSend = req.FormValue("mail-to-trev")

	// if subjOfMail == "" {
	// 	subjOfMail = "Sent From AceAdmin via. Sendgrid"
	// }

	// if fromWho == "" {
	// 	fromWho = "Anonymous@trevorknott.io"
	// }

	from := mail.NewEmail("TK - From AceAdmin", fromWho)
	subject := subjOfMail
	to := mail.NewEmail("TK - From AceAdmin", "tk@trevorknott.io")
	plainTextContent := mailToSend
	htmlContent := "<strong>..." + mailToSend + "</strong>"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		println("Response StatusCode")
		fmt.Println(response.StatusCode)
		println("\n" + "Response Body")
		fmt.Println(response.Body + "\n")
		println("Response Headers")
		fmt.Println(response.Headers)
	}

	http.Redirect(w, req, "/", 301)
}
