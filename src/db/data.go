package db

import "notes/app/services"

var notes = []services.Note{
	{Title: "Shopping List", Content: "2 hours ago"},
	{Title: "Meeting Notes", Content: "Yesterday"},
	{Title: "Recipe Ideas", Content: "3 days ago"},
	{Title: "Travel Plans", Content: "1 week ago"},
	{Title: "Book Recommendations", Content: "2 weeks ago"},
	{Title: "TV Suggestions", Content: "1 week ago"},
}

func GetNotes() []services.Note {
	return notes
}

func AddNote(title string, content string) {
	var note services.Note
	note.Title = title
	note.Content = content
	notes = append(notes, note)
}
