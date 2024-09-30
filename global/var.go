package global

import "time"

var Users = map[string]*User{}

type User struct {
	Name    string             `json:"name"`
	Folders map[string]*Folder `json:"folders"`
}

type File struct {
	Name        string
	Description string
	CreatedAt   time.Time
}

type Folder struct {
	Name        string
	Description string
	CreatedAt   time.Time
	Files       map[string]*File
}
