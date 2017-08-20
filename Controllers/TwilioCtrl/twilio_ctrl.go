package TwilioCtrl

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/sfreiberg/gotwilio"
	globals "github.com/tknott95/Ace_Go/Globals"
)

func TwilioTest(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var msgToSend string = req.FormValue("msg-to-trev")
	var msgSenderName string = req.FormValue("msg-name")
	var msgSenderNum string = req.FormValue("msg-num")

	accountSid := globals.TwilioSID
	authToken := globals.TwilioAuthToken
	twilio := gotwilio.NewTwilioClient(accountSid, authToken)

	if msgSenderName == "" {
		msgSenderName = "Tesla(ANON)"
	}

	if msgSenderNum == "" {
		msgSenderNum = "(970) 581-3161"
	}

	from := "+19707142241"
	to := "+19705813161"
	message := "Ace - FROM:" + msgSenderName + " #: " + msgSenderNum + " Msg2you: " + msgToSend
	twilio.SendSMS(from, to, message, "", "")

	println("Sending txt to TK (admin):", msgToSend)

	http.Redirect(w, req, "http://trevorknott.io", 301)
}
