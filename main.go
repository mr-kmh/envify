package main

import (
	"io/fs"
	"log"
	"os"
	"strings"
)

func readCWD() []fs.DirEntry {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("can't read current directory, error %s", err)
	}
	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatalf("can't read directory %s, error %s", dir, err)
	}
	return files
}

func filterDotEnv(files []fs.DirEntry) []fs.DirEntry {
	dotEnvFiles := make([]fs.DirEntry, 0)
	for _, file := range files {
		if strings.HasPrefix(file.Name(), ".env") {
			dotEnvFiles = append(dotEnvFiles, file)
		}
	}
	return dotEnvFiles
}

func main() {
	for _, file := range filterDotEnv(readCWD()) {
		log.Println(file.Name())
	}
}
