// trigger sends a trigger request via network broadcast.
package main

import (
	"fmt"
	"os"

	"github.com/asjoyner/argus/trigger"
)

func main() {
	err := trigger.Trigger()
	if err != nil {
		fmt.Printf("error sending trigger request: %+v\n", err)
		os.Exit(1)
	}
	fmt.Println("Trigger request sent.")
}
