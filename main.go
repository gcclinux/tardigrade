package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
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
			} else if os.Args[1] == "-selectf" {
				fmt.Println()
				fmt.Println("ERROR - MISSING ARGUMENTS: \n\n>> tardigrade -selectf \"format\" (raw|json|key|value)")
			} else if os.Args[1] == "-selectl" {
				fmt.Println()
				fmt.Println("ERROR - MISSING ARGUMENTS: \n\n>> tardigrade -selectl \"format\" (raw|json|key|value)")
			} else if os.Args[1] == "-selecti" {
				fmt.Println()
				fmt.Println("ERROR - MISSING ARGUMENTS: \n\n>> tardigrade -selecti \"id\" \"format\" (raw|json|key|value)")
			} else if os.Args[1] == "-total" {
				fmt.Println()
				msg := tar.CountSize()
				fmt.Println("(", msg, ")")
			} else if os.Args[1] == "-selectfx" {
				fmt.Println()
				fmt.Println("ERROR - MISSING ARGUMENTS: \n\n>> tardigrade -selectfx \"amount\"")
			} else if os.Args[1] == "-selectlx" {
				fmt.Println()
				fmt.Println("ERROR - MISSING ARGUMENTS: \n\n>> tardigrade -selectlx \"amount\"")
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
			} else if os.Args[1] == "-selectf" {
				fmt.Println()
				fmt.Println(tar.FirstField(os.Args[2]))
			} else if os.Args[1] == "-selectl" {
				fmt.Println()
				fmt.Println(tar.LastField(os.Args[2]))
			} else if os.Args[1] == "-selecti" {
				fmt.Println()
				fmt.Println("ERROR - MISSING ARGUMENTS: \n\n>> tardigrade -selecti \"id\" \"format\" (raw|json|key|value)")
			} else if os.Args[1] == "-total" {
				fmt.Println("ERROR - REMOVE:", os.Args[2], "\n\n>> tardigrade -total")
			} else if os.Args[1] == "-selectfx" {
				if x, err := strconv.Atoi(os.Args[2]); err == nil {
					fmt.Println()
					var received = tar.FirstXFields(x)

					bytes := received
					var data []MyStruct
					size := len(data)
					json.Unmarshal(bytes, &data)

					if size == 1 {
						fmt.Printf("id: %v, key: %v, data: %s", data[0].Id, data[0].Key, data[0].Data)
					} else {
						for x := range data {
							fmt.Printf("id: %v, key: %v, data: %s", data[x].Id, data[x].Key, data[x].Data)
							fmt.Println()
						}
					}
				} else {
					fmt.Println("ERROR - FLAG:(", os.Args[2], ") is not a number!\n\n>> tardigrade -selectfx \"number\"")
				}
			} else if os.Args[1] == "-selectlx" {
				if x, err := strconv.Atoi(os.Args[2]); err == nil {
					fmt.Println()
					var received = tar.LastXFields(x)

					bytes := received
					var data []MyStruct
					size := len(data)
					json.Unmarshal(bytes, &data)

					if size == 1 {
						fmt.Printf("id: %v, key: %v, data: %s", data[0].Id, data[0].Key, data[0].Data)
					} else {
						for x := range data {
							fmt.Printf("id: %v, key: %v, data: %s", data[x].Id, data[x].Key, data[x].Data)
							fmt.Println()
						}
					}
				} else {
					fmt.Println("ERROR - FLAG:(", os.Args[2], ") is not a number!\n\n>> tardigrade -selectlx \"number\"")
				}
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
			} else if os.Args[1] == "-selectf" {
				fmt.Println()
				fmt.Println("ERROR - REMOVE:", os.Args[3], "\n\n>> tardigrade -selectf \"format\" (raw|json|key|value)")
			} else if os.Args[1] == "-selectl" {
				fmt.Println()
				fmt.Println("ERROR - REMOVE:", os.Args[3], "\n\n>> tardigrade -selectl \"format\" (raw|json|key|value)")
			} else if os.Args[1] == "-select" {
				fmt.Println()
				id, _ := strconv.Atoi(os.Args[2])
				fmt.Println(tar.SelectByID(id, os.Args[3]))
			} else if os.Args[1] == "-total" {
				fmt.Println("ERROR - REMOVE:", os.Args[2], "AND", os.Args[3], "\n\n>> tardigrade -total")
			} else if os.Args[1] == "-selectfx" {
				fmt.Println()
				fmt.Println("ERROR - REMOVE:", os.Args[2], "AND", os.Args[3], "\n\n>> tardigrade -selectfx \"amount\"")
			} else if os.Args[1] == "-selectlx" {
				fmt.Println()
				fmt.Println("ERROR - REMOVE:", os.Args[2], "AND", os.Args[3], "\n\n>> tardigrade -selectlx \"amount\"")
			}
		} else if size == 5 {
			fmt.Println()
			fmt.Println("ERROR - INVALID SYNTAX PROVIDED CHECK MANUAL")
		}
	}
}
