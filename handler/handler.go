package handler

import (
	"html/template"
	"log"
	"net/http"
	"path"
)

func MainHandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/" {
		http.NotFound(w,r)
		return
	}
	tmpl, err := template.ParseFiles(path.Join("views", "form.html"), path.Join("views", "layouts.html"))
	if err != nil {
		log.Println(err)
		http.Error(w, "Error Goblok!!!",http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error ex!!!",http.StatusInternalServerError)
		return
	}

}

func SayHandler(w http.ResponseWriter, r *http.Request){
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
			http.Error(w, "Error Goblok!!!",http.StatusInternalServerError)
			return
		}
		name := r.Form.Get("name")
		tmpl, err := template.ParseFiles(path.Join("views", "say.html"), path.Join("views", "layouts.html"))
		if err != nil {
			log.Println(err)
			http.Error(w, "Error ex!!!",http.StatusInternalServerError)
			return
		}
		err = tmpl.Execute(w,name)
	}
}