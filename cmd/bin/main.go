package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	var received = LastXFields(2)

	bytes := received
	var data []MyStruct
	json.Unmarshal(bytes, &data)

	for l := range data {
		fmt.Printf("id: %v, key: %v, data: %s", data[l].Id, data[l].Key, data[1].Data)
		fmt.Println()
	}
}
