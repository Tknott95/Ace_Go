package srvCtrl

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	adminCtrl "github.com/tknott95/MasterGo/Controllers/AdminCtrl"
	blogCtrl "github.com/tknott95/MasterGo/Controllers/BlogCtrl"
	langCtrl "github.com/tknott95/MasterGo/Controllers/LangCtrl"
	// sGrid_Ctrl "github.com/tknott95/MasterGo/Controllers/SgridCtrl"
	twilioCtrl "github.com/tknott95/MasterGo/Controllers/TwilioCtrl"
	globals "github.com/tknott95/MasterGo/Globals"
)

func InitServer() {
	mux := httprouter.New()

	mux.GET("/twil_ctrl", twilioPage)
	mux.POST("/txt", twilioCtrl.TwilioTest)

	// PC LANGS
	mux.GET("/", index)
	mux.GET("/pc_langs", langFetch)
	mux.POST("/pc_langs/delete/:lang_id", langCtrl.LangDelete) /* Calls both via. url not form val */
	mux.POST("/pc_langs/add", langCtrl.LangAdd)                /* will use formval in blog portion for sure tho */

	// BLOG POSTS
	mux.GET("/blog_posts", blogFetch)
	mux.POST("/blog_posts/add", blogCtrl.BlogPostAdd)
	mux.POST("/blog_posts/delete/:post_id/:pic_rmv", blogCtrl.BlogPostDel)

	// Admin Login
	mux.GET("/admin_signin", adminCtrl.AdminPage)
	mux.POST("/admin_sigin", adminCtrl.AdminLogin)

	/* UMBRELLA API PORTION */
	/* Will use /api/ always! */
	mux.GET("/api/pc_langs", langCtrl.ApiLangFetch)
	mux.GET("/api/blog_posts", blogCtrl.ApiBlogFetch)

	//http.Handle("/Public/", http.StripPrefix("/Public", http.FileServer(http.Dir("./Public"))))

	mux.NotFound = http.StripPrefix("/Public", http.FileServer(http.Dir("./Public")))

	// handler for serving files
	// mux.ServeFiles("/Public/*filepath", http.Dir("/var/www/Public/"))
	http.ListenAndServe(globals.PortNumber, mux)

}

func index(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	println("📝 Currently on Index page.")

	// sGrid_Ctrl.SendEmail()

	globals.Tmpl.ExecuteTemplate(w, "index.gohtml", nil)
}

/* Base Fetch ALL Query Pages Rendered Here  */

func langFetch(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	println("📝 Currently on Language Control page.")

	globals.Tmpl.ExecuteTemplate(w, "langs_fetch.gohtml", langCtrl.FetchLangs())
}

func blogFetch(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	println("📝 Currently on Blog Post Control page.")

	globals.Tmpl.ExecuteTemplate(w, "blog_control.gohtml", blogCtrl.BlogPostFetch())
}

func twilioPage(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	println("📝 Currently on Twilio page.")

	globals.Tmpl.ExecuteTemplate(w, "twilio_msg.gohtml", nil)
}
