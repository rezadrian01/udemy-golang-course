package note

import (
	"encoding/json"
	"errors"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title string `json:"title"`
	Content string `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func (n *Note) Save() error{
	filename := strings.ReplaceAll(n.Title, " ", "_")
	filename = strings.TrimSuffix(filename, "\n")
	filename = strings.TrimSuffix(filename, "\r") + ".json"

	jsonData, err := json.Marshal(n)
	if err != nil{
		return err
	}

	return os.WriteFile(filename, jsonData, 0644)
}

func New(title, content string) (*Note, error){
	if strings.TrimSpace(title) == "" || strings.TrimSpace(content) == ""{
		return nil, errors.New("title and content are required")
	}

	return &Note{
		title, 
		content, 
		time.Now(),
	}, nil
}