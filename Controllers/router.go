package srvCtrl

import (
	"net/http"

	"github.com/rs/cors"

	"github.com/julienschmidt/httprouter"
	adminCtrl "github.com/tknott95/Ace_Go/Controllers/AdminCtrl"
	blogCtrl "github.com/tknott95/Ace_Go/Controllers/BlogCtrl"
	langCtrl "github.com/tknott95/Ace_Go/Controllers/LangCtrl"
	sGrid_Ctrl "github.com/tknott95/Ace_Go/Controllers/SgridCtrl"
	twilioCtrl "github.com/tknott95/Ace_Go/Controllers/TwilioCtrl"
	globals "github.com/tknott95/Ace_Go/Globals"
)

func InitServer() {
	mux := httprouter.New()

	/* CORS */
	handler := cors.Default().Handler(mux)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://trevorknott.io"},
		AllowCredentials: true,
	})

	// Insert the middleware
	handler = c.Handler(handler)

	/* CORS END */

	mux.GET("/twil_ctrl", twilioPage)
	mux.POST("/txt", twilioCtrl.TwilioTest)

	mux.GET("/sgrid_ctrl", sGridPage)
	mux.POST("/email", sGrid_Ctrl.SendEmail)

	// PC LANGS
	mux.GET("/", index)
	mux.GET("/pc_langs", langFetch)
	mux.GET("/pc_langs/edit/:lang-id", langCtrl.LangSingleFetch)
	mux.POST("/pc_langs/delete/:lang_id", langCtrl.LangDelete) /* Calls both via. url not form val */
	mux.POST("/pc_langs/add", langCtrl.LangAdd)
	mux.POST("/pc_langs/edit/:l-id", langCtrl.LangUpdate) /* will use formval in blog portion for sure tho */

	// BLOG POSTS
	mux.GET("/blog_posts", blogFetch)
	mux.GET("/blog_posts/edit/:post-id", blogCtrl.SinglePostFetch)
	mux.POST("/blog_posts/edit/:blog-id", blogCtrl.BlogUpdate)
	mux.POST("/blog_posts/add", blogCtrl.BlogPostAdd)
	mux.POST("/blog_posts/update/:blog_id", blogCtrl.BlogUpdate)
	mux.POST("/blog_posts/delete/:post_id/:pic_rmv", blogCtrl.BlogPostDel)

	// Admin Login
	mux.GET("/admin_signin", adminCtrl.AdminPage)
	mux.GET("/admin_logout", adminCtrl.AdminLogout)
	mux.POST("/admin_sigin", adminCtrl.AdminLogin)

	/* UMBRELLA API PORTION */
	/* Will use /api/ always! */
	mux.GET("/api/pc_langs", langCtrl.ApiLangFetch)
	mux.GET("/api/blog_posts", blogCtrl.ApiBlogFetch)

	//http.Handle("/Public/", http.StripPrefix("/Public", http.FileServer(http.Dir("./Public"))))

	mux.NotFound = http.StripPrefix("/Public", http.FileServer(http.Dir("./Public")))

	// handler for serving files
	// mux.ServeFiles("/Public/*filepath", http.Dir("/var/www/Public/"))
	http.ListenAndServe(globals.PortNumber, handler)

}

func index(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	println("📝 Currently on Index page.")

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

func sGridPage(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	println("📝 Currently on SendGrid page.")

	globals.Tmpl.ExecuteTemplate(w, "sgrid_msg.gohtml", nil)
}
