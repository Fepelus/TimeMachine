/*
Copyright 2014 Patrick Borgeest
The MIT Licence.  See /LICENSE.txt
*/
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

const (
	DEFAULT_SOURCE_ROOT  = "/MyWorld"
	DEFAULT_TARGET_ROOT  = "/TimeMachine"
	SNAPSHOTS_DIR        = "/snapshots"
	OBJECTS_DIR          = "/objects"
	SNAPSHOT_TIME_FORMAT = "2006-01-02T1304"
)

var (
	source_root string
	target_root string
)

func handleCommandline() {
	flag.StringVar(&source_root, "sourcedir", DEFAULT_SOURCE_ROOT, "Source directory to be backed up")
	flag.StringVar(&target_root, "targetdir", DEFAULT_TARGET_ROOT, "Target directory to up")
	flag.Parse()
}
func main() {
	handleCommandline()
	if !isTargetReady(getSnapshotsDir(), getObjectsDir()) {
		log.Fatalf("Target %s is not set up with snapshots and objects\n", target_root)
		os.Exit(1)
	}
	handleDir(source_root)
}

func getObjectsDir() string {
	return fmt.Sprintf("%s%s", target_root, OBJECTS_DIR)
}

var now = time.Now().Format(SNAPSHOT_TIME_FORMAT)

func getSnapshotsDir() string {
	return fmt.Sprintf("%s%s", target_root, SNAPSHOTS_DIR)
}

func getSnapshotDir() string {
	return fmt.Sprintf("%s/%s", getSnapshotsDir(), now)
}
