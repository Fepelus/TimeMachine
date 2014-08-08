/*
Copyright 2014 Patrick Borgeest
MIT Licence.  See /LICENSE.txt
*/
package main

import (
	"os"
	"testing"
)

func touch(path string) {
	_, e := os.Open(path)
	if os.IsNotExist(e) {
		f, _ := os.Create(path)
		defer f.Close()
	}
}

func TestCreate(t *testing.T) {
	testsavefile := savefile{"/tmp"}
	if testsavefile.getPath() != "/tmp" {
		t.Error("Should get path back but got: ", testsavefile.getPath())
	}
}

func TestHash(t *testing.T) {
	filename := "/tmp/bob"
	touch(filename)
	testsavefile := savefile{filename}
	testhash, err := testsavefile.makeHash()
	if err != nil {
		t.Error("first get hash threw error:", err)
	}
	hash, err := testhash.getHash()
	if err != nil {
		t.Error("second get hash threw error:", err)
	}
	if hash != "da39a3ee5e6b4b0d3255bfef95601890afd80709" {
		t.Error("Should get hash back but got: ", hash)
	}
	if testhash.getHashPrefix() != "da/39/a3" {
		t.Error("Should get 'da/39/a3' back but got: ", testhash.getHashPrefix())
	}
	if testhash.getObjectFilename() != "/objects/da/39/a3/sha1-da39a3ee5e6b4b0d3255bfef95601890afd80709" {
		t.Error("Should get object filename but got: ", testhash.getObjectFilename())
	}
	if testhash.getSnapshotFilename() != getSnapshotDir()+filename {
		t.Error("Should get snapshot filename but got: ", testhash.getSnapshotFilename())
	}
}
