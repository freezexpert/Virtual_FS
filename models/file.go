package models

import "Virtual_FS/global"

func CreateFile(username, foldername, filename, description string) error {
	return nil
}

func DeleteFile(username, foldername, filenam string) error {
	return nil
}

func ListFiles(username, foldername, sortType, order string) ([]*global.File, error) {
	return []*global.File{}, nil
}
