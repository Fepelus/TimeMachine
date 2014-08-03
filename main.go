package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

const (
	SOURCE_ROOT   = "/MyWorld"
	TARGET_ROOT   = "/TimeMachine"
	SNAPSHOTS_DIR = "/snapshots"
	OBJECTS_DIR   = "/objects"
)

func main() {
	if !isTargetReady(getSnapshotsDir(), getObjectsDir()) {
		log.Fatalf("Target %s is not set up with snapshots and objects\n", TARGET_ROOT)
		os.Exit(1)
	}
	handleDir(SOURCE_ROOT)
}

func getObjectsDir() string {
	return fmt.Sprintf("%s%s", TARGET_ROOT, OBJECTS_DIR)
}

var now = time.Now().Format("2006-01-02T03:04:05")

func getSnapshotsDir() string {
	return fmt.Sprintf("%s%s%s", TARGET_ROOT, SNAPSHOTS_DIR, now)
}
