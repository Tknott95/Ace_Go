package globals

import "html/template"

var PortNumber string = ":" + "8080"

var Tmpl = template.Must(template.ParseGlob("./Views/*"))
