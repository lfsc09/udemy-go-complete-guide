package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type Note struct {
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
}

func New(title, body string) (*Note, error) {
	if title == "" || body == "" {
		return nil, errors.New("invalid note input")
	}

	note := &Note{
		Title:     title,
		Body:      body,
		CreatedAt: time.Now(),
	}

	err := note.save()

	if err != nil {
		return nil, err
	}

	return note, nil
}

func (note *Note) Display() {
	fmt.Println("--------")
	fmt.Println("Title:", note.Title)
	fmt.Println("Body:", note.Body)
	fmt.Println("Created At:", note.CreatedAt)
	fmt.Println("--------")
}

func (note *Note) save() error {
	path := "note.json"
	json, err := json.Marshal(note)

	if err != nil {
		return err
	}

	return os.WriteFile(path, json, 0644)
}
