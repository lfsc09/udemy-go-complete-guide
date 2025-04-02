package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/lfsc09/udemy-go-complete-guide/section-5/note"
)

func main() {
	title, body := readNote()

	note, err := note.New(title, body)
	if err != nil {
		fmt.Println(err)
		return
	}

	note.Display()
	fmt.Println("Note saved successfully!")
}

func readNote() (string, string) {
	title := promptInput("Enter title")
	body := promptInput("Enter body")

	return title, body
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
