package globals

import (
	"html/template"
)

var PortNumber string = ":" + "8080"

var Tmpl = template.Must(template.ParseGlob("./Views/*"))

var TwilioSID = "ACc2be7e8f4aac8fccf91ee7ae1e51c779"
var TwilioAuthToken = "3f96cfd74a3fd9fe91e670aca9635183"

var SGridApi = "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"

// var CurTime = time.Now()
