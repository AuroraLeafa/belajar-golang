package main

import (
	"embed"
	_ "embed"
	"fmt"
	"io/fs"
	"os"
)

//go:embed test/version.txt
var version string

//go:embed test/logo.png
var logo []byte

//go:embed test/files/*.txt
var path embed.FS

func main() {
	fmt.Println(version)

	err := os.WriteFile("logo.png", logo, fs.ModePerm)
	if err != nil {
		panic(err)
	}

	dirEntry, err := path.ReadDir("test/files")
	if err != nil {
		panic(err)
	}
	for _, entry := range dirEntry {
		if !entry.IsDir() {
			fmt.Println(entry.Name())
			file, _ := path.ReadFile("test/files/" + entry.Name())
			fmt.Println(string(file))
		}
	}
}
