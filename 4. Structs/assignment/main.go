package main

import (
	"bufio"
	"fmt"
	"os"

	"example.com/note"
)

func main() {
	fmt.Println("Hello World")
	title := getUserData("Note title: ")
	content := getUserData("Note content: ")
	fmt.Printf("Title: %v\nContent: \n\n%v\n", title, content)

	fmt.Println("Saving Note...")
	userNote, err := note.New(title, content)
	if err != nil{
		fmt.Println(err)
		return
	}
	userNote.Save()
	fmt.Println("Successfully to save note")

}

func getUserData(prompt string) string{
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil{
		return ""
	}
	return text
}