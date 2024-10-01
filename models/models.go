package models

import (
	"regexp"
	"time"
)

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

func IsValidName(name string) bool {
	re := regexp.MustCompile(`[^A-Za-z0-9'-]`)
	if re.MatchString(name) {
		return false
	} else {
		return true
	}
}
