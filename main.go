package main

import (
	"os"
)

// Directory functions
func removeDirectory(username, dir string) error {
	path := "./" + username + "/" + dir
	err := os.RemoveAll(path)
	if err != nil {
		return err
	}
	return nil
}

// func changeDirectory() error {

// }

func getDirectoryList(username string) ([]string, error) {
	path := "./" + username
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}
	var list []string
	for _, entry := range entries {
		if entry.IsDir() {
			list = append(list, entry.Name())
		}
	}
	return list, nil

}

func createDirectory(username, dirName string) error {

	path := username + "/" + dirName
	err := os.MkdirAll(path, 0755)
	if err != nil {
		return err
	}
	return nil
}

// File functions

// func removeFile() error {

// }

// func createFile() error {

// }

// func uploadFile() error {

// }

// func downloadFile() error {

// }

// func getFileList() error {

// }

func main() {

}
