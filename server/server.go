package server

import (
	"fmt"
	"log"
	"net/http"

	//"github.com/julienschmidt/httprouter"
	"github.com/gorilla/mux"
)

func stdout(out string) {
	fmt.Println(out)
}

/*
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")
var templates = template.Must(template.ParseFiles("template/edit.html", "template/view.html"))

func rebuildTmpl() {
	stdout("rebuil template")
	templates = template.Must(template.ParseFiles("template/edit.html", "template/view.html"))
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {

	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, "rendererro : \n "+err.Error(), http.StatusInternalServerError)
	}
}
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s !", r.URL.Path[1:])
}

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		rt, ok := r.URL.Query()["rt"]

		if ok && rt[0] == "1" {
			rebuildTmpl()
		}

		fn(w, r, m[2])
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := LoadPage(title)
	if err != nil {
		// fmt.Fprintf(w, "%s", err)
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	// fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
	// t, _ := template.ParseFiles("template/view.html")
	// t.Execute(w, p)
	renderTemplate(w, "view", p)
}

//loads the page
func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := LoadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}

	// fmt.Fprintf(w, "<h1>Editing %s</h1>"+
	// 	"<form action=\"/save/%s\" method=\"POST\">"+
	// 	"<textarea name=\"body\">%s</textarea><br>"+
	// 	"<input type=\"submit\" value=\"Save\">"+
	// 	"</form>", p.Title, p.Title, p.Body)
	// t, _ := template.ParseFiles("template/edit.html")
	// t.Execute(w, p)
	renderTemplate(w, "edit", p)

}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {

	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.SavePage()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Welcome!\n")
}
*/

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/hello", helloHandler).Methods("GET")

	// Declare the static file directory and point it to the
	// directory we just made
	staticFileDirectory := http.Dir("static")

	// Declare the handler, that routes requests to their respective filename.
	// The fileserver is wrapped in the `stripPrefix` method, because we want to
	// remove the "/assets/" prefix when looking for files.
	// For example, if we type "/assets/index.html" in our browser, the file server
	// will look for only "index.html" inside the directory declared above.
	// If we did not strip the prefix, the file server would look for
	// "./assets/assets/index.html", and yield an error
	staticFileHandler := http.StripPrefix("/static/", http.FileServer(staticFileDirectory))

	// The "PathPrefix" method acts as a matcher, and matches all routes starting
	// with "/assets/", instead of the absolute route itself
	r.PathPrefix("/static/").Handler(staticFileHandler).Methods("GET")

	return r
}

//Run runserver
func Run() {

	//router := httprouter.New()
	//router.GET("/", Index)
	// router.GET("/view/", viewHandler)
	// router.GET("/static/")
	//router.NotFound = http.FileServer(http.Dir("static"))

	//https://stackoverflow.com/questions/43601359/how-do-i-serve-css-and-js-in-go
	//http.Handle("/", http.FileServer(http.Dir("static/css")))
	// http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// http.HandleFunc("/view/", makeHandler(viewHandler))
	// http.HandleFunc("/edit/", makeHandler(editHandler))
	// http.HandleFunc("/save/", makeHandler(saveHandler))

	r := newRouter()

	stdout("server run on 8080")
	log.Fatal(http.ListenAndServe(":8080", r))

}
