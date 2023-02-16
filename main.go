package main

import (
	"encoding/json"
	"fmt"
)

// This main function is only used for testing it will be deleted from the final version as this becomes a module.

func main() {

	tar := Tardigrade{}
	var received = tar.LastXFields(1)

	bytes := received
	var data []MyStruct
	json.Unmarshal(bytes, &data)

	fmt.Printf("id: %v, key: %v, data: %s", data[0].Id, data[0].Key, data[0].Data)

	// for x := range data {
	// 	fmt.Printf("id: %v, key: %v, data: %s", data[x].Id, data[x].Key, data[x].Data)
	// 	fmt.Println()
	// }

	// func (*Tardigrade).FirstXFields(count int) []byte
	// func (*Tardigrade).GetUpdated() (updated string)
	// func (*Tardigrade).GetVersion() (release string)
	// func (*Tardigrade).LastXFields(count int) []byte
	// func (*Tardigrade).ModifyField(id int, k string, v string) bool
	// func (*Tardigrade).RemoveField(id int) (string, bool)

}
