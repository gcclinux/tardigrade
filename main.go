package main

import (
	"fmt"
	"os"
)

// Starting to create an example how to parse external arguments or flags:
// Build a binary before using it: $ go build -o tardigrade *.go && ./tardigrade -createdb

// TOBE Continued!

func main() {

	tar := Tardigrade{}

	if len(os.Args) == 1 {
		fmt.Println("No arguments parsed! Try: tardigrade -help")
	} else {
		size := len(os.Args)
		if size == 2 {
			// CreateDB() function
			if os.Args[1] == "-createdb" {
				msg, status := tar.CreateDB()
				fmt.Println()
				fmt.Println(msg, "(", status, ")")
			} else if os.Args[1] == "-copydb" {
				fmt.Println()
				msg, status := tar.CreatedDBCopy()
				fmt.Println(msg, "(", status, ")")
			} else if os.Args[1] == "-deletedb" {
				fmt.Println()
				msg, status := tar.DeleteDB()
				fmt.Println(msg, "(", status, ")")
			} else if os.Args[1] == "-insert" {
				fmt.Println()
				fmt.Println("ERROR - MISSING ARGUMENTS: \n\n>> tardigrade -insert \"key filed details\" \"data field details\"")
			}
		} else if size == 3 {
			if os.Args[1] == "-createdb" {
				fmt.Println("ERROR - REMOVE:", os.Args[2], "\n\n>> tardigrade -createdb")
			} else if os.Args[1] == "-copydb" {
				fmt.Println("ERROR - REMOVE:", os.Args[2], "\n\n>> tardigrade -copydb")
			} else if os.Args[1] == "-deletedb" {
				fmt.Println("ERROR - REMOVE:", os.Args[2], "\n\n>> tardigrade -deletedb")
			} else if os.Args[1] == "-insert" {
				fmt.Println("ERROR - MISSING ARGUMENTS:", os.Args[1], "\n\n>> tardigrade -insert \"key filed details\" \"data field details\"")
			}
		} else if size == 4 {
			if os.Args[1] == "-createdb" {
				fmt.Println("ERROR - REMOVE:", os.Args[2], "AND", os.Args[3], "\n\n>> tardigrade -createdb")
			} else if os.Args[1] == "-copydb" {
				fmt.Println("ERROR - REMOVE:", os.Args[2], "AND", os.Args[3], "\n\n>> tardigrade -copydb")
			} else if os.Args[1] == "-deletedb" {
				fmt.Println("ERROR - REMOVE:", os.Args[2], "AND", os.Args[3], "\n\n>> tardigrade -deletedb")
			} else if os.Args[1] == "-insert" {
				fmt.Println()
				status := tar.AddField(os.Args[2], os.Args[3])
				fmt.Println("returned: (", status, ")")
			}
		}
	}
}
