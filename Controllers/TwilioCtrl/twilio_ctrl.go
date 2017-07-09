package TwilioCtrl

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/sfreiberg/gotwilio"
	globals "github.com/tknott95/MasterGo/Globals"
)

func TwilioTest(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	var msgToSend string
	// msgToSend = ps.ByName("msg")

	msgToSend = req.FormValue("msg-to-trev")

	accountSid := globals.TwilioSID
	authToken := globals.TwilioAuthToken
	twilio := gotwilio.NewTwilioClient(accountSid, authToken)

	from := "+19707142241"
	to := "+19705813161"
	message := "Message From Ace Admin Board: " + msgToSend
	twilio.SendSMS(from, to, message, "", "")

	println("Sending txt to TK (admin):", msgToSend)

	http.Redirect(w, req, "/", 301)
}
