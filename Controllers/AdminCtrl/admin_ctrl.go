package AdminCtrl

import (
	"fmt"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	mydb "github.com/tknott95/Ace_Go/Controllers/DbCtrl"

	globals "github.com/tknott95/Ace_Go/Globals"
)

func AdminPage(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	println("Admin Login Called")

	globals.Tmpl.ExecuteTemplate(w, "admin_signin.gohtml", nil)
}

func AdminLogin(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	var admin_name string
	var admin_password string

	admin_name = req.FormValue("admin-email")
	admin_password = req.FormValue("admin-password")

	rows, err := mydb.Store.DB.Query(`SELECT admin_email FROM admin_users;`)
	fmt.Println(w, "Established admin_users db connection", nil)
	if err != nil {
		println("Admin user fetch error: ", err)
	}
	defer rows.Close()

	var name string
	var names []string

	// query
	for rows.Next() {
		err = rows.Scan(&name)
		// check(err)

		names = append(names, name)

	}

	if admin_name == names[0] && admin_password == "Welcome1!" {
		os.Setenv("admin", "true")
		println("ADMIN LOGGED IN CORRECTLY")
		http.Redirect(w, req, "/", 301)

		// tpl.ExecuteTemplate(w, "admin_users.gohtml", names)
	} else {
		fmt.Fprintf(w, "ADMIN - Log In Failed")
	}
}

func IsAdminLoggedIn() bool {
	if os.Getenv("admin") == "true" {
		return true
	} else {
		return false

	}
}
