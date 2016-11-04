package main

import (
	"log"
	"os"

	"github.com/fsnotify/fsnotify"
)

func getWatched() []string {
	args := os.Args[1:]

	return args
}

func main() {
	// Check to make sure we've got some valid inputs. If not, no need to proceed.
	files := getWatched()

	if len(files) == 0 {
		log.Fatal("No valid files to watch specified.")
	}

	watcher, err := fsnotify.NewWatcher()
	done := make(chan bool)

	if err != nil {
		log.Fatal(err)
	}

	// Concurrent goroutine that awaits events and does Smart Things when they occur.
	go func() {
		log.Println("Awaiting changes...")
		for {
			select {
			case event := <-watcher.Events:
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("modified file:", event.Name)
				}
			case err := <-watcher.Errors:
				log.Println("Error: ", err)
				close(done)
			}
		}
	}()

	defer watcher.Close()

	for _, filename := range files {
		log.Println("Watching", filename)
		err = watcher.Add(filename)
	}

	<-done
}
