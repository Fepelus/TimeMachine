package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func visit(path string, f os.FileInfo, err error) error {
	if f.IsDir() {
		return os.MkdirAll(fmt.Sprintf("%s%s", getSnapshotsDir(), path), 0766)
	}
	hashPathAndLink(path)
	return nil
}

func handleDir(source string) {
	_ = filepath.Walk(root, visit)
}
