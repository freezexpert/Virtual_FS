package models

import (
	"fmt"
	"strings"
)

func Register(name string) error {
	name = strings.ToLower(name)
	if _, isExist := Users[name]; isExist {
		return fmt.Errorf("the %s has already existed", name)
	}
	if !IsValidName(name) {
		return fmt.Errorf("the %s contains invalid chars", name)
	}
	Users[name] = &User{Name: name, Folders: make(map[string]*Folder)}
	return nil
}
