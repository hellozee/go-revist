package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: <program_name> <root_directory>")
		return
	}
	watcher := NewWatcher(os.Args[1], 2*time.Second)
	f := func(name string, status FileStatus) {
		fmt.Println(status, " ", name)
	}
	watcher.Start(f)
}

type FileStatus int

const (
	create FileStatus = iota
	modified
	erased
)

func (fs FileStatus) String() string {
	return [...]string{"Created", "Modified", "Erased"}[fs]
}

// Watcher The Watcher data structure
type Watcher struct {
	paths      map[string]time.Time
	searchPath string
	delay      time.Duration
}

// Start  To start the file watcher
func (w *Watcher) Start(action func(name string, status FileStatus)) {

	for file := range w.paths {
		if _, err := os.Stat(file); err != nil {
			delete(w.paths, file)
			action(file, erased)
		}
	}

	pathIterator := func(p string, info os.FileInfo, err error) error {
		lastWrite := info.ModTime()

		if _, ok := w.paths[p]; !ok {
			w.paths[p] = lastWrite
			action(p, create)
			return nil
		}

		if w.paths[p] != lastWrite {
			w.paths[p] = lastWrite
			action(p, modified)
		}

		return nil
	}

	for {
		time.Sleep(w.delay)
		//look for changes
		err := filepath.Walk(w.searchPath, pathIterator)
		if err != nil {
			fmt.Errorf("Unable to walk through the directory")
		}
	}
}

// NewWatcher  The Watcher constructor
func NewWatcher(dir string, waitDelay time.Duration) *Watcher {
	w := Watcher{
		searchPath: dir,
		delay:      waitDelay,
		paths:      make(map[string]time.Time),
	}

	// list all the present paths
	writeTimings := func(p string, info os.FileInfo, err error) error {
		w.paths[p] = info.ModTime()
		return nil
	}
	err := filepath.Walk(w.searchPath, writeTimings)

	if err != nil {
		fmt.Errorf("Unable to create the Watcher")
	}

	return &w
}
