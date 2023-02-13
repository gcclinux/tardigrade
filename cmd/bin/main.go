package main

import "fmt"

func main() {

	AddField("one", "string data test")
	AddField("two", "string data test II")
	AddField("three", "string data testIII")
	AddField("four", "string data test IV")

	fmt.Println("\nLast Line: ", LastField())
	fmt.Println("Count Lines: ", CountSize())

	fmt.Println("\n Line 10:", SelectID(10))

}
