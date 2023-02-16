package main

import "fmt"

func main() {
	tar := Tardigrade{}
	fmt.Println(tar.GetUpdated())
	fmt.Println(tar.GetVersion())
}
