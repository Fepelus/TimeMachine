package main

import (
	"crypto/sha1"
	"fmt"
	"io/ioutil"
)

type hash interface {
	getHash() (string, error)
}
type savefile struct {
	path string
}

func (this savefile) hashPathAndLink(path string) error {
	hash, err := this.makeHash()
	if err != nil {
		return err
	}
	// make this transactional
	if !doesFileExist(hash.getObjectFilename()) {
		err = hash.copyToObject()
		if err != nil {
			return err
		}
	}
	return err
}

func (this savefile) getPath() string {
	this.path
}

func (this savefile) getHash() (string, error) {
	hashed, err := this.makeHash()
	if err != nil {
		return "", err
	}
	return hashed.getHash()
}
func (this savefile) makeHash() (hashedfile, error) {
	contents, err := ioutil.ReadFile(this.path)
	if err != nil {
		return nil, err
	}
	h := sha1.New()
	h.Write(contents)
	bs := h.Sum(nil)
	return hashedfile{this, fmt.Sprintf("%x", bs)}, nil
}

func (this hashedfile) getHash() (string, error) {
	return this.hash, nil
}

type hashedfile struct {
	savefile
	hash string
}

func (this hashedfile) getObjectFilename() string {
	return fmt.Sprintf("%s/%s/%s", getObjectsDir(), this.getHashPrefix(), this.getHash())
}

func (this hashedfile) getSnapshotFilename() string {
	return fmt.Sprintf("%s/%s", getSnapshotsDir(), this.getPath(), this.getHash())
}
