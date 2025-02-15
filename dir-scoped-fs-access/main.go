package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: dir-scoped-fs-access <file>")
		os.Exit(1)
	}

	file := os.Args[1]

	root, err := os.OpenRoot("files")
	if err != nil {
		panic(err)
	}

	defer root.Close()

	f, err := root.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	info, err := f.Stat()
	if err != nil {
		panic(err)
	}

	fmt.Printf("File: %s, size :%d\n", info.Name(), info.Size())
}
