package main

import (
	"fmt"
	"io/fs"
	"path/filepath"
)

func pathFunc() {
	fmt.Println(filepath.Base("../file/zaza.txt"))
	fmt.Println(filepath.Base("../file"))
	fmt.Println(filepath.Ext("../file/zaza.txt"))
	fmt.Println(filepath.Dir("../file/zaza.txt"))
	fmt.Println(filepath.IsAbs("../file/zaza.txt"))
	fmt.Println(filepath.Join("a", "b", "c", "d.txt"))
	fmt.Println(filepath.Split("../file/zaza.txt"))
	fmt.Println(filepath.ToSlash("../file/zaza.txt"))
	fmt.Println(filepath.FromSlash("../file/zaza.txt"))

	err := filepath.WalkDir("../../", visit)
	if err != nil {
		fmt.Println("Error walking the path:", err)
	}
}

func visit(path string, d fs.DirEntry, err error) error {
	if err != nil {
		return err
	}

	fmt.Println(" ", path, d.IsDir())
	return nil
}
