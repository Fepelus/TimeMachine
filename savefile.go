package main

import (
	"crypto/sha1"
	"fmt"
	"io/ioutil"
	"os"
)

type savefile struct {
	path string
}

func (this savefile) hashPathAndLink() error {
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
	hash.hardLink()
	return err
}

func (this savefile) getPath() string {
	return this.path
}

func (this savefile) makeHash() (hashedfile, error) {
	contents, err := ioutil.ReadFile(this.path)
	if err != nil {
		return hashedfile{}, err
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
	hash, _ := this.getHash()
	return fmt.Sprintf("%s/%s/sha1-%s", getObjectsDir(), this.getHashPrefix(), hash)
}

func (this hashedfile) getSnapshotFilename() string {
	return fmt.Sprintf("%s%s", getSnapshotDir(), this.getPath())
}

func (this hashedfile) getHashPrefix() string {
	hash, _ := this.getHash()
	return fmt.Sprintf("%s/%s/%s", hash[0:2], hash[2:4], hash[4:6])
}

func (this hashedfile) copyToObject() error {
	contents, err := ioutil.ReadFile(this.path)
	if err != nil {
		return err
	}
	err = os.MkdirAll(fmt.Sprintf("%s/%s", getObjectsDir(), this.getHashPrefix()), 0766)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(this.getObjectFilename(), contents, 0755)
}

func (this hashedfile) hardLink() error {
	return os.Link(this.getObjectFilename(), this.getSnapshotFilename())
}
