/*
Copyright 2014 Patrick Borgeest
The MIT Licence.  See /LICENSE.txt
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
	if f.IsDir() {
		return os.MkdirAll(fmt.Sprintf("%s%s", getSnapshotDir(), path), 0766)
	}
	pathfile := savefile{path}
	errr := pathfile.hashPathAndLink()
	if errr != nil {
		log.Println("ERROR", path, errr)
	}
	return nil
}

func handleDir(source string) {
	_ = filepath.Walk(source, visit)
}
