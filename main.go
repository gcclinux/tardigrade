package main

import (
	"fmt"
)

// This main function is only used for testing it will be deleted from the final version as this becomes a module.

func main() {
	tar := Tardigrade{}
	fmt.Println(tar.GetUpdated())
	fmt.Println(tar.GetVersion())
}
