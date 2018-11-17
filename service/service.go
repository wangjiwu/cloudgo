package service

import (
	"net/http"
	"os"

	"github.com/urfave/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

// NewServer configures and returns a Server.
func NewServer() *negroni.Negroni {
	formatter := render.New()

	n := negroni.Classic()
	mx := mux.NewRouter()

	initRoutes(mx, formatter)

	n.UseHandler(mx)
	return n
}

func initRoutes(mx *mux.Router, formatter *render.Render) {
	webRoot := os.Getenv("WEBROOT")
	//fmt.Print(webRoot + "a")
	if len(webRoot) == 0 {
		if root, err := os.Getwd(); err != nil {
			panic("Could not retrive working directory")
		} else {
			webRoot = root
			//`fmt.Println(root)
		}
	}
	mx.HandleFunc("/api/test", apiTestHandler(formatter)).Methods("GET")
	mx.HandleFunc("/api/unknown", NotImplemented).Methods("GET")
	mx.HandleFunc("/", home).Methods("GET")
	mx.HandleFunc("/", login).Methods("POST")
	
	/*实现对css/js/images的引用*/
	mx.PathPrefix("/static").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(webRoot+"/assets/"))))
}
