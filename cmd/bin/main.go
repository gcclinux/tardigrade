package main

import "fmt"

func main() {

	getOS()

	if AddField("New Entry", "string of data representing a the value") == true {
		fmt.Println("LastField()", LastField("raw"))

		fmt.Println("UniqueID()", UniqueID())

		fmt.Println("CountSize()", CountSize())
	} else {
		fmt.Println("AddField() - Faied because could not find Database", nil)
	}
}
