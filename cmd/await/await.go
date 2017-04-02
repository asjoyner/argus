package main

import (
	"fmt"
	"os"

	"github.com/asjoyner/argus/trigger"
)

func main() {
	w, err := trigger.NewWatcher()
	if err != nil {
		fmt.Printf("error initializing trigger watcher: %+v\n", err)
		os.Exit(1)
	}
	w.Wait()
	fmt.Println("Trigger request received.")
}
