package main

import (
	"encoding/json"
	"fmt"
)

// This main function is only used for testing it will be deleted from the final version as this becomes a module.

func main() {
	// Delete the field id 2 row
	fmt.Println("Return: ", RemoveField(2))

	// Check results print last 3 rows
	var received = FirstXFields(2)
	bytes := received
	var data []MyStruct
	json.Unmarshal(bytes, &data)

	for l := range data {
		fmt.Printf("id: %v, key: %v, data: %s", data[l].Id, data[l].Key, data[1].Data)
		fmt.Println()
	}
}
