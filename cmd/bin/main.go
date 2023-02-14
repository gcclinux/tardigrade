package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	var tmpRecords MyStruct
	out, _ := json.MarshalIndent(tmpRecords, "", "	")
	result := string(out)
	fmt.Println(LastXFields(10, "json"))
	fmt.Println(result)

	// NOT FINISHED
}
