package models

func CreateFile(username, foldername, filename, description string) error {
	return nil
}

func DeleteFile(username, foldername, filenam string) error {
	return nil
}

func ListFiles(username, foldername, sortType, order string) ([]*File, error) {
	return []*File{}, nil
}
