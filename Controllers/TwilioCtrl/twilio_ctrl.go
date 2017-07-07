package TwilioCtrl

import "github.com/sfreiberg/gotwilio"

func TwilioTest() {
	accountSid := "ACc2be7e8f4aac8fccf91ee7ae1e51c779"
	authToken := "3f96cfd74a3fd9fe91e670aca9635183"
	twilio := gotwilio.NewTwilioClient(accountSid, authToken)

	from := "+19707142241"
	to := "+19705813161"
	message := "Welcome to gotwilio! Hit homepage of golang admin board!"
	twilio.SendSMS(from, to, message, "", "")
}
