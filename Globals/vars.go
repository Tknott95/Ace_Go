package globals

import (
	"html/template"
	"time"
)

var PortNumber string = ":" + "8080"

var Tmpl = template.Must(template.ParseGlob("./Views/*"))

var CurTime = time.Now()
