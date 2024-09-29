package models

import "time"

type Folder struct {
	Name        string
	Description string
	CreatedAt   time.Time
	Files       map[string]*File
}

func CreateFolder(user, folder, description string) error {
	return nil
}

func DeleteFolder(user, folder string) error {
	return nil
}

func ListFolders(user, sortType, order string) ([]*Folder, error) {
	return []*Folder{}, nil
}

func RenameFolder(username, foldername, newFolderName string) error {
	return nil
}
