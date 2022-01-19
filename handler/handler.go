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