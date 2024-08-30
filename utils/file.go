package utils

import (
	"io"
	"os"
)

// CopyFile is a utility function to copy a file from the source path to the destination path.
var CopyFile = copyFile

// CreateDirsIfNotExist is a utility function to create directories if they do not already exist.
var CreateDirsIfNotExist = createDirsIfNotExist

// copyFile copies a file from the source path to the destination path.
// It returns an error if the source file cannot be opened, the destination file cannot be created,
// or if there is an issue during the copy process.
func copyFile(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return err
	}

	return nil
}

// OpenFile opens a file for appending. If the file does not exist, it creates it with write permissions.
// It returns a pointer to the opened file and an error if any occurs during the file opening process.
func OpenFile(path string) (*os.File, error) {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return nil, err
	}
	return file, nil
}

// createDirsIfNotExist checks if the directories at the given paths exist, and creates them if they do not.
// It returns an error if directory creation fails.
func createDirsIfNotExist(paths ...string) error {
	for _, p := range paths {
		if _, err := os.Stat(p); os.IsNotExist(err) {
			err := os.MkdirAll(p, os.ModePerm)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
