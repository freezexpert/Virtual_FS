package models

import "Virtual_FS/global"

func CreateFolder(user, folder, description string) error {
	
	return nil
}

func DeleteFolder(user, folder string) error {
	return nil
}

func ListFolders(user, sortType, order string) ([]*global.Folder, error) {
	return []*global.Folder{}, nil
}

func RenameFolder(username, foldername, newFolderName string) error {
	return nil
}
