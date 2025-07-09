package main

import (
	"fmt"
	"net/http"
	"notes/app/components"
	"notes/app/db"

	"github.com/a-h/templ"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	// UI
	http.Handle("/home", templ.Handler(components.Home()))
	http.Handle("/inbox", templ.Handler(components.Inbox(db.GetNotes())))
	http.Handle("/note/new", templ.Handler(components.NewNote()))

	// API
	http.Handle("/api/note", templ.Handler(components.NewNote()))

	fmt.Println("Server starting at http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server failed:", err)
	}
}

type AddNoteRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}
