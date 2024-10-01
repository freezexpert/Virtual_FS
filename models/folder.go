package models

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

func CreateFolder(username, foldername, description string) error {
	user, exists := Users[username]
	if !exists {
		return fmt.Errorf("the %s doesn't exist", username)
	}
	foldername = strings.ToLower(foldername)
	if !IsValidName(foldername) {
		return fmt.Errorf("the %s contains invalid chars", foldername)
	}
	if _, exists := user.Folders[foldername]; exists {
		return fmt.Errorf("the %s has already existed", foldername)
	}
	user.Folders[foldername] = &Folder{Name: foldername, Description: description, CreatedAt: time.Now(), Files: make(map[string]*File)}
	return nil
}

func DeleteFolder(username, foldername string) error {
	user, exists := Users[username]
	if !exists {
		return fmt.Errorf("the %s doesn't exist", username)
	}
	foldername = strings.ToLower(foldername)
	if _, exists := user.Folders[foldername]; !exists {
		return fmt.Errorf("the %s doesn't exist", foldername)
	}
	delete(user.Folders, foldername)
	return nil
}

// Need to add sort function and add Warning return
func ListFolders(username, sortType, order string) ([]*Folder, error) {
	user, exists := Users[username]
	if !exists {
		return nil, fmt.Errorf("the %s doesn't exist", username)
	}
	if len(user.Folders) == 0 {
		return nil, fmt.Errorf("Warning: The user doesn't have any folders")
	}
	fmt.Println("Folders:")
	for _, folder := range user.Folders {
		fmt.Printf("%s - %s - %s\n", folder.Name, folder.Description, folder.CreatedAt.Format("2006-01-02 15:04:05"))
	}
	return []*Folder{}, nil
}

func RenameFolder(username, foldername, newFoldername string) error {
	user, exists := Users[username]
	if !exists {
		return fmt.Errorf("the %s doesn't exist", username)
	}
	foldername = strings.ToLower(foldername)
	newFoldername = strings.ToLower(newFoldername)
	if !IsValidName(newFoldername) {
		return errors.New("the new folder name contains invalid chars")
	}
	folder, exists := user.Folders[foldername]
	if !exists {
		return fmt.Errorf("the %s doesn't exist", foldername)
	}
	if _, exists := user.Folders[newFoldername]; exists {
		return errors.New("the new folder name has already existed")
	}
	folder.Name = newFoldername
	user.Folders[newFoldername] = folder
	delete(user.Folders, foldername)
	return nil
}
