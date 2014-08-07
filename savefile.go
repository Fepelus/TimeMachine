/*
Copyright 2014 Patrick Borgeest
The MIT Licence.  See /LICENSE.txt
*/
package main

import (
	"bufio"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
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
	f, err := os.Open(this.path)
	if err != nil {
		return hashedfile{}, err
	}
	defer f.Close()
	var reader *bufio.Reader
	reader = bufio.NewReader(f)
	sha1 := sha1.New()
	_, err = io.Copy(sha1, reader)
	if err != nil {
		return hashedfile{}, err
	}
	hash := hex.EncodeToString(sha1.Sum(nil))

	return hashedfile{this, hash}, nil
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
	err := os.MkdirAll(fmt.Sprintf("%s/%s", getObjectsDir(), this.getHashPrefix()), 0766)
	if err != nil {
		return err
	}
	return copyFile(this.path, this.getObjectFilename())
}

func (this hashedfile) hardLink() error {
	return os.Link(this.getObjectFilename(), this.getSnapshotFilename())
}

func copyFile(src, dest string) error {
	sf, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sf.Close()

	df, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer df.Close()

	_, err = io.Copy(df, sf)
	return err
}
