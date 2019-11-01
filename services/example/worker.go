package example

import (
	"fmt"
)

// Worker do what i say
type Worker struct{}

// Start worker
func (w Worker) Start() {
	fmt.Println("Im started differently, and im working on something")
}
