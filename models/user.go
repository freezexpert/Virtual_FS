package models

import (
	"Virtual_FS/global"
	"fmt"
	"regexp"
	"strings"
)

func Register(name string) error {
	name = strings.ToLower(name)
	if _, isExist := global.Users[name]; isExist {
		return fmt.Errorf("the %s has already existed", name)
	}
	re := regexp.MustCompile(`[^A-Za-z0-9'-]`)
	if re.MatchString(name) {
		return fmt.Errorf("the %s contains invalid chars", name)
	}
	global.Users[name] = &global.User{Name: name, Folders: make(map[string]*global.Folder)}
	return nil
}
