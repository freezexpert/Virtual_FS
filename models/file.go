package models

import (
	"fmt"
	"strings"
	"time"
)

func CreateFile(username, foldername, filename, description string) error {
	user, exists := Users[username]
	if !exists {
		return fmt.Errorf("the %s doesn't exist", username)
	}
	foldername = strings.ToLower(foldername)
	folder, exists := user.Folders[foldername]
	if !exists {
		return fmt.Errorf("the %s doesn't exist", foldername)
	}
	filename = strings.ToLower(filename)
	if !IsValidName(filename) {
		return fmt.Errorf("the %s contains invalid chars", filename)
	}
	if _, exists := folder.Files[filename]; exists {
		return fmt.Errorf("the %s has already existed", filename)
	}
	folder.Files[filename] = &File{Name: filename, Description: description, CreatedAt: time.Now()}
	return nil
}

func DeleteFile(username, foldername, filename string) error {
	user, exists := Users[username]
	if !exists {
		return fmt.Errorf("the %s doesn't exist", username)
	}
	foldername = strings.ToLower(foldername)
	folder, exists := user.Folders[foldername]
	if !exists {
		return fmt.Errorf("the %s doesn't exist", foldername)
	}
	filename = strings.ToLower(filename)
	if _, exists := folder.Files[filename]; !exists {
		return fmt.Errorf("the %s doesn't exist", filename)
	}
	delete(folder.Files, filename)
	return nil
}

func ListFiles(username, foldername, sortType, order string) ([]*File, error) {
	return []*File{}, nil
}
