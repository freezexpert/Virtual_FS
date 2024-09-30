package models

import "time"

var Users = map[string]*User{}

type User struct {
	Name    string
	Folders map[string]*Folder
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
