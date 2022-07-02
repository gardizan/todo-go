package main

import (
	"html/template"
	"log"
	"net/http"
)

var tmpl *template.Template

type Todo struct {
	Item  string
	Feito bool
}

type PaginaData struct {
	Titulo string
	Todos  []Todo
}

func todo(w http.ResponseWriter, r *http.Request) {
	data := PaginaData{
		Titulo: "Lista de TO DO",
		Todos: []Todo{
			{Item: "Instalar Go", Feito: true},
			{Item: "Variáveis Go", Feito: true},
			{Item: "Servidor Http Go", Feito: true},
			{Item: "TO DO em Go", Feito: false},
		},
	}

	tmpl.Execute(w, data)
}

func main() {
	mux := http.NewServeMux()
	tmpl = template.Must(template.ParseFiles("templates/index.gohtml"))

	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	mux.HandleFunc("/todo", todo)
	log.Fatal(http.ListenAndServe(":3001", mux))
}
