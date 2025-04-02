package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type Todo struct {
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
}

func New(title string) (*Todo, error) {
	if title == "" {
		return nil, errors.New("invalid todo input")
	}

	todo := &Todo{
		Title:     title,
		CreatedAt: time.Now(),
	}

	return todo, nil
}

func (todo *Todo) Display() {
	fmt.Println("--------")
	fmt.Println("Title:", todo.Title)
	fmt.Println("Created At:", todo.CreatedAt)
	fmt.Println("--------")
}

func (todo *Todo) Save() error {
	path := "todo.json"
	json, err := json.Marshal(todo)

	if err != nil {
		return err
	}

	return os.WriteFile(path, json, 0644)
}
