/*
Copyright 2014 Patrick Borgeest
MIT Licence.  See /LICENSE.txt
*/
package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func visit(path string, f os.FileInfo, err error) error {
	if err != nil {
		log.Println("Error calling: ", path, err)
	}
	if f == nil {
		log.Println("No file: ", path)
		return nil
	}
	if f.IsDir() {
		return os.MkdirAll(fmt.Sprintf("%s%s", getSnapshotDir(), path), 0766)
	}
	if f.Mode().IsRegular() {
		pathfile := savefile{path}
		errr := pathfile.hashPathAndLink()
		if errr != nil {
			log.Println("ERROR", path, errr)
		}
	}
	return nil
}

func handleDir(source string) {
	_ = filepath.Walk(source, visit)
}
