package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type File struct {
	Type string `json:"type"`
	Name string `json:"name"`
	Children *File `json:"children,omitempty"`
}

func OSReadDir(root string) (string) {
	var files string
	f, err := os.Open(root)
	if err != nil {
		return files
	}
	fileInfo, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		return files
	}

	for _, file := range fileInfo {
		return file.Name()
	}
	return files
}

func function(path string) File {
	d := File{}
	err := filepath.Walk(path,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			d.Name = info.Name()

			if info.IsDir() {
				d.Type = "directory"

				f := function(filepath.Join(path, OSReadDir(path)))
				var c *File = &f
				d.Children = c

				output, _ := json.MarshalIndent(d, "", "     ")
				fmt.Println(string(output))
			} else {
				d.Type = "file"
				d.Children = nil
			}
			return nil
		})
	if err != nil {
		log.Println(err)
	}

	return d
}

func main() {
	function("./task3")
}
