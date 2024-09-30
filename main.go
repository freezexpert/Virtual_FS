package main

import (
	"Virtual_FS/models"
	"fmt"
	"os"
	"strings"
)

func main() {
	for {
		var cmd string
		fmt.Printf("# ")
		fmt.Scan(&cmd)
		switch cmd {
		case "register":
			var username string
			fmt.Scan(&username)
			err := models.Register(username)
			if err != nil {
				// Capitalize the first character
				errmsg := strings.ToUpper(string(err.Error()[0])) + string(err.Error()[1:])
				fmt.Fprintln(os.Stderr, "Error: "+errmsg)
			} else {
				fmt.Println("Add", username, "successfully.")
			}

		case "create-folder":
			var username, foldername, description string
			fmt.Scan(&username, &foldername, &description)
			err := models.CreateFolder(username, foldername, description)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Create", foldername, "successfully.")
			}

		case "delete-folder":
			var username, foldername string
			fmt.Scan(&username, &foldername)
			err := models.DeleteFolder(username, foldername)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Delete", foldername, "successfully.")
			}

		case "list-folders":
			var username, sortType, order string
			fmt.Scan(&username, &sortType, &order)
			folders, err := models.ListFolders(username, sortType, order)
			if err != nil {
				fmt.Println("Error:", err)
			}
			fmt.Println("List", folders, "successfully.")

		case "rename-folder":
			var username, foldername, newFolderName string
			fmt.Scan(&username, &foldername, &newFolderName)
			err := models.RenameFolder(username, foldername, newFolderName)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Rename", foldername, "to", newFolderName, "successfully.")
			}

		case "create-file":
			var username, foldername, filename, description string
			fmt.Scan(&username, &foldername, &filename, &description)
			err := models.CreateFile(username, foldername, filename, description)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Create", filename, "in", username+"/"+foldername, "successfully.")
			}

		case "delete-file":
			var username, foldername, filename string
			fmt.Scan(&username, &foldername, &filename)
			err := models.DeleteFile(username, foldername, filename)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Delete", filename, "in", username+"/"+foldername, "successfully.")
			}

		case "list-files":
			var username, foldername, sortType, order string
			fmt.Scan(&username, &foldername, &sortType, &order)
			files, err := models.ListFiles(username, foldername, sortType, order)
			if err != nil {
				fmt.Println("Error:", err)
			}
			fmt.Println(files)

		default:
			fmt.Println("Error: Unrecognized command.")
		}
	}

}
