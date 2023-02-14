package main

import (
	"fmt"
)

func main() {

	//var tmpRecords MyStruct
	// out, _ := json.MarshalIndent(tmpRecords, "", "	")
	// result := string(out)

	var receivedRecords = LastXFields(10)
	fmt.Println(receivedRecords)

	// NOT FINISHED
}
