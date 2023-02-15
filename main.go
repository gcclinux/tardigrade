package main

import "fmt"

// This main function is only used for testing it will be deleted from the final version as this becomes a module.

func main() {

	// Check current information in ROW 2 BEFORE CHANGE
	fmt.Println(SelectByID(2, "raw"))

	// Modify ROW 2 with new information provided in key & value
	var change = ModifyField(2, "Updated key", "Updated data set with new inforation")
	fmt.Println("Changed: ", change)

	// Check current information in ROW 2 AFTER CHANGE
	fmt.Println(SelectByID(2, "raw"))
}
