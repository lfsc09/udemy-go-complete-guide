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
	CreatedAt time.Time `json:"created_at"`
}

func New(title string) (*Note, error) {
	if title == "" {
		return nil, errors.New("invalid note input")
	}

	note := &Note{
		Title:     title,
		CreatedAt: time.Now(),
	}

	return note, nil
}

func (note *Note) Display() {
	fmt.Println("--------")
	fmt.Println("Title:", note.Title)
	fmt.Println("Created At:", note.CreatedAt)
	fmt.Println("--------")
}

func (note *Note) Save() error {
	path := "note.json"
	json, err := json.Marshal(note)

	if err != nil {
		return err
	}

	return os.WriteFile(path, json, 0644)
}
