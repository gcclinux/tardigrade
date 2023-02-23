package main

// Built Tue 21 Feb 22:05:28 GMT 2023

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
			} else if os.Args[1] == "-version" {
				fmt.Println()
				vr := tar.GetVersion()
				dt := tar.GetUpdated()
				fmt.Println(dt, "(", vr, ")")
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
				fmt.Println("ERROR - MISSING ARGUMENTS: \n\n>> tardigrade -selectfx \"amount\" \"format\" (raw|json|key|value)")
			} else if os.Args[1] == "-selectlx" {
				fmt.Println()
				fmt.Println("ERROR - MISSING ARGUMENTS: \n\n>> tardigrade -selectlx \"amount\" \"format\" (raw|json|key|value)")
			} else if os.Args[1] == "-change" {
				fmt.Println()
				fmt.Println("ERROR - MISSING ARGUMENTS: \n\n>> tardigrade -change \"id\" \"key data\" \"value date\"")
			} else if os.Args[1] == "-deletef" {
				fmt.Println()
				fmt.Println("ERROR - MISSING ARGUMENTS: \n\n>> tardigrade -deletef \"id\"")
			} else if os.Args[1] == "-help" {
				fmt.Println(`
-createdb		"CREATE new database"
-copydb 		"CREATE backup (copy) of the database"
-deletedb 		"DELETE database"
-deletef 		"DELETE <id> specific row from database"
-insert 		"INSERT <field one> <field two> for new entry"
-change 		"CHANGE <id> <field one> <field two> on existing row "
-selectf 		"SELECT <format> TOP row from database "
-selectl 		"SELECT <format> LAST row from database"
-selectfx 		"SELECT <number> <format> TOP rows from database"
-selectlx 		"SELECT <number> <format> LAST rows from database"
-selecti 		"SELECT <id> specific row from database"
-total 			"SHOW number of entries in database"
-version		"SHOW built date & version`)
				fmt.Println()
			} else {
				fmt.Println()
				fmt.Println("ERROR - INVALID SYNTAX PROVIDED CHECK MANUAL")
			}
		} else if size == 3 {
			if os.Args[1] == "-createdb" {
				fmt.Println("ERROR - REMOVE:", os.Args[2], "\n\n>> tardigrade -createdb")
			} else if os.Args[1] == "-version" {
				fmt.Println("ERROR - REMOVE:", os.Args[2], "\n\n>> tardigrade -version")
			} else if os.Args[1] == "-copydb" {
				fmt.Println("ERROR - REMOVE:", os.Args[2], "\n\n>> tardigrade -copydb")
			} else if os.Args[1] == "-help" {
				fmt.Println("ERROR - REMOVE:", os.Args[2], "\n\n>> tardigrade -help")
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
				fmt.Println("ERROR - MISSING ARGUMENTS: \n\n>> tardigrade -selectfx \"amount\" \"format\" (raw|json|key|value)")
			} else if os.Args[1] == "-selectlx" {
				fmt.Println("ERROR - MISSING ARGUMENTS: \n\n>> tardigrade -selectlx \"amount\" \"format\" (raw|json|key|value)")
			} else if os.Args[1] == "-change" {
				fmt.Println()
				fmt.Println("ERROR - MISSING ARGUMENTS: \n\n>> tardigrade -change \"id\" \"key data\" \"value date\"")
			} else if os.Args[1] == "-deletef" {
				fmt.Println()
				if x, err := strconv.Atoi(os.Args[2]); err == nil {
					msg, status := tar.RemoveField(x)
					fmt.Println(msg, status)
				} else {
					fmt.Println("ERROR - FLAG:(", os.Args[2], ") is not a number!\n\n>> tardigrade -deletef \"number\"")
				}
			} else {
				fmt.Println()
				fmt.Println("ERROR - INVALID SYNTAX PROVIDED CHECK MANUAL")
			}
		} else if size == 4 {
			if os.Args[1] == "-createdb" {
				fmt.Println("ERROR - REMOVE:", os.Args[2], "AND", os.Args[3], "\n\n>> tardigrade -createdb")
			} else if os.Args[1] == "-copydb" {
				fmt.Println("ERROR - REMOVE:", os.Args[2], "AND", os.Args[3], "\n\n>> tardigrade -copydb")
			} else if os.Args[1] == "-version" {
				fmt.Println("ERROR - REMOVE:", os.Args[2], "AND", os.Args[3], "\n\n>> tardigrade -version")
			} else if os.Args[1] == "-help" {
				fmt.Println("ERROR - REMOVE:", os.Args[2], "AND", os.Args[3], "\n\n>> tardigrade -help")
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
				format := os.Args[3]
				if x, err := strconv.Atoi(os.Args[2]); err == nil {
					fmt.Println()
					var format, received = tar.FirstXFields(x, format)
					bytes := received
					var data []MyStruct
					json.Unmarshal(bytes, &data)
					for x := range data {
						if format == "json" {
							out, _ := json.MarshalIndent(&data[x], "", "  ")
							fmt.Printf(string(out))
							fmt.Println()
						} else if format == "value" {
							fmt.Println(string(data[x].Data))
							fmt.Println()
						} else if format == "raw" {
							fmt.Printf("id: %d, key: %v, data: %s\n", data[x].Id, data[x].Key, data[x].Data)
						} else if format == "key" {
							fmt.Printf("%v\n", data[x].Key)
						} else if format == "id" {
							fmt.Println(strconv.Itoa(data[x].Id))
							fmt.Println()
						} else {
							fmt.Printf("Invalid format provided!")
						}
					}
				} else {
					fmt.Println("ERROR - FLAG:(", os.Args[2], ") is not a number!\n\n>> tardigrade -selectfx \"number\" \"format\" (raw|json|key|value)")
				}
			} else if os.Args[1] == "-selectlx" {
				format := os.Args[3]
				if x, err := strconv.Atoi(os.Args[2]); err == nil {
					fmt.Println()
					var format, received = tar.LastXFields(x, format)

					bytes := received
					var data []MyStruct
					json.Unmarshal(bytes, &data)
					for x := range data {
						if format == "json" {
							out, _ := json.MarshalIndent(&data[x], "", "  ")
							fmt.Printf(string(out))
							fmt.Println()
						} else if format == "value" {
							fmt.Println(string(data[x].Data))
							fmt.Println()
						} else if format == "raw" {
							fmt.Printf("id: %d, key: %v, data: %s\n", data[x].Id, data[x].Key, data[x].Data)
						} else if format == "key" {
							fmt.Printf("%v\n", data[x].Key)
						} else if format == "id" {
							fmt.Println(strconv.Itoa(data[x].Id))
							fmt.Println()
						} else {
							fmt.Printf("Invalid format provided!")
						}
					}
				} else {
					fmt.Println("ERROR - FLAG:(", os.Args[2], ") is not a number!\n\n>> tardigrade -selectlx \"number\" \"format\" (raw|json|key|value)")
				}
			} else if os.Args[1] == "-deletef" {
				fmt.Println()
				fmt.Println("ERROR - REMOVE:", os.Args[2], "AND", os.Args[3], "\n\n>> tardigrade -deletef \"amount\"")
			} else {
				fmt.Println()
				fmt.Println("ERROR - INVALID SYNTAX PROVIDED CHECK MANUAL")
			}
		} else if size == 5 {
			if os.Args[1] == "-change" {
				fmt.Println()
				if x, err := strconv.Atoi(os.Args[2]); err == nil {
					change, status := tar.ModifyField(x, os.Args[3], os.Args[4])
					fmt.Println(change, status)
				} else {
					fmt.Println("ERROR - FLAG:(", os.Args[2], ") is not a number!\n\n>> tardigrade -selectfx \"number\"")
				}
				fmt.Println()
			} else {
				fmt.Println()
				fmt.Println("ERROR - INVALID SYNTAX PROVIDED CHECK MANUAL")
			}
		} else if size == 6 {
			fmt.Println()
			fmt.Println("ERROR - INVALID SYNTAX PROVIDED CHECK MANUAL")
		}
	}
}
