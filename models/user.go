package models

type User struct {
	Name    string             `json:"name"`
	Folders map[string]*Folder `json:"folders"`
}

func Register(name string) error {

	return nil
}
