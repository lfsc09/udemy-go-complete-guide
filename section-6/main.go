package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/lfsc09/udemy-go-complete-guide/section-6/note"
	"github.com/lfsc09/udemy-go-complete-guide/section-6/todo"
)

type saver interface {
	Save() error
}

type outputable interface {
	saver
	Display()
}

func main() {
	note := newNote()
	err := output(note)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Saved successfully!")
	}

	todo := newTodo()
	err = output(todo)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Saved successfully!")
	}
}

func output(obj outputable) error {
	obj.Display()
	return obj.Save()
}

func newNote() *note.Note {
	title := promptInput("Enter note title")

	note, err := note.New(title)

	if err != nil {
		return nil
	}
	return note
}

func newTodo() *todo.Todo {
	title := promptInput("Enter todo title")

	todo, err := todo.New(title)

	if err != nil {
		return nil
	}
	return todo
}

func promptInput(prompt string) string {
	fmt.Printf("%v: ", prompt)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
