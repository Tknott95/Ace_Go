package TwilioCtrl

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/sfreiberg/gotwilio"
	globals "github.com/tknott95/Ace_Go/Globals"
)

type TxtMsg struct {
	MsgNum  string `json:"msg-num"`
	MsgName string `json:"msg-name"`
	MsgToMe string `json:"msg-to-trev"`
}

func TwilioTest(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var txt_res TxtMsg
	decoder := json.NewDecoder(req.Body)
	decoder.Decode(&txt_res)

	accountSid := globals.TwilioSID
	authToken := globals.TwilioAuthToken
	twilio := gotwilio.NewTwilioClient(accountSid, authToken)

	if txt_res.MsgName == "" {
		txt_res.MsgName = "Nigel"
	}

	if txt_res.MsgNum == "" {
		txt_res.MsgNum = "(970)581-3161"
	}

	from := "+19707142241"
	to := "+19705813161"
	message := "Ace -" + txt_res.MsgName + "- #: " + txt_res.MsgNum + "- msg: " + txt_res.MsgToMe
	twilio.SendSMS(from, to, message, "", "")

	fmt.Println("Ace -" + txt_res.MsgName + "- #: " + txt_res.MsgNum + "- msg: " + txt_res.MsgToMe)

	//http.Redirect(w, req, "http://www.trevorknott.io", 301)
}
