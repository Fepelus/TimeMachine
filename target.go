package main

import "os"

func isTargetReady(snapshots string, objects string) bool {
	return doesFileExist(snapshots) && doesFileExist(objects)
}

func doesFileExist(filename string) bool {
	_, err := os.Open(filename)
	return !os.IsNotExist(err)
}
