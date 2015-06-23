package watcher

import (
	"log"
	"path/filepath"

	"github.com/minodisk/go-jsonschema/tools/utils"

	"gopkg.in/fsnotify.v1"
)

func Watch(filenames []string, callback func(filename string)) error {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	defer watcher.Close()

	done := make(chan bool)
	dirs := make(map[string]bool)
	for i, filename := range filenames {
		filename = filepath.Clean(filename)
		filenames[i] = filename
		mode, err := utils.FileMode(filename)
		if err != nil {
			return err
		}
		var dir string
		if mode.IsDir() {
			dir = filename
		} else {
			dir = filepath.Dir(filename)
		}
		dirs[dir] = true
	}

	go func() {
		for {
			select {
			case event := <-watcher.Events:
				log.Println(event)
				if event.Op&fsnotify.Create == fsnotify.Create || event.Op&fsnotify.Write == fsnotify.Write {
					if in(filenames, event.Name) {
						log.Printf("[watcher] detect modified: %s", event.Name)
						callback(event.Name)
					}
				}
			case err := <-watcher.Errors:
				log.Printf("[watcher] error: %s", err)
			}
		}
	}()
	for dir, _ := range dirs {
		log.Printf("[watcher] watch dir: %s", dir)
		watcher.Add(dir)
	}

	<-done
	return nil
}

func in(arr []string, elem string) bool {
	for _, e := range arr {
		if e == elem {
			return true
		}
	}
	return false
}
