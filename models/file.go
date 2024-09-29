package models

import "time"

type File struct {
	Name        string
	Description string
	CreatedAt   time.Time
}

func CreateFile(username, foldername, filename, description string) error {
	return nil
}

func DeleteFile(username, foldername, filenam string) error {
	return nil
}

func ListFiles(username, foldername, sortType, order string) ([]*File, error) {
	return []*File{}, nil
}
