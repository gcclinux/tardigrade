package main

import "fmt"

func main() {

	fmt.Println("AddField(): ",
		AddField("New Entry", "string of data representing a the value"))

	fmt.Println("SelectByID()", SelectByID(10, "raw"))

	fmt.Println("SelectByID()", SelectByID(10, "value"))
}
