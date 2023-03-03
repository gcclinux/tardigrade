package main

// Updated Fri  3 Mar 20:53:48 GMT 2023

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gcclinux/tardigrade-mod"
)

type MyStruct struct {
	Id   int
	Key  string
	Data string
}

// Starting to create an example how to parse external arguments or flags:
// Build a binary before using it: $ go build -o tardigrade *.go && ./tardigrade -createdb
func main() {

	tar := tardigrade.Tardigrade{}

	if len(os.Args) == 1 {
		fmt.Println("No arguments parsed! Try: ", filepath.Base(os.Args[0]), " --help")
		fmt.Println()
	} else {
		size := len(os.Args)
		if size == 2 {
			// CreateDB() function
			if os.Args[1] == "--createdb" {
				msg, status := tar.CreateDB()
				fmt.Println()
				fmt.Println(msg, "(", status, ")")
			} else if os.Args[1] == "--version" {
				fmt.Println()
				vr := AppGetVersion()
				dt := AppGetUpdated()
				modv := tar.GetVersion()
				modr := tar.GetUpdated()
				fmt.Println("(App)", dt, "(", vr, ")")
				fmt.Println("(Mod)", modr, "(", modv, ")")
				fmt.Println()
			} else if os.Args[1] == "--copydb" {
				fmt.Println()
				msg, status := tar.CreatedDBCopy()
				fmt.Println(msg, "(", status, ")")
				fmt.Println()
			} else if os.Args[1] == "--deletedb" {
				fmt.Println()
				msg, status := tar.DeleteDB()
				fmt.Println(msg, "(", status, ")")
				fmt.Println()
			} else if os.Args[1] == "--insert" {
				fmt.Println()
				fmt.Println("ERROR - MISSING ARGUMENTS: \n\n>> ", filepath.Base(os.Args[0]), " --insert \"key filed details\" \"data field details\"")
				fmt.Println()
			} else if os.Args[1] == "--selectf" {
				fmt.Println()
				fmt.Println("ERROR - MISSING ARGUMENTS: \n\n>> ", filepath.Base(os.Args[0]), " --selectf \"format\" (raw|json|key|value)")
				fmt.Println()
			} else if os.Args[1] == "--selectl" {
				fmt.Println()
				fmt.Println("ERROR - MISSING ARGUMENTS: \n\n>> ", filepath.Base(os.Args[0]), " --selectl \"format\" (raw|json|key|value)")
				fmt.Println()
			} else if os.Args[1] == "--selecti" {
				fmt.Println()
				fmt.Println("ERROR - MISSING ARGUMENTS: \n\n>> ", filepath.Base(os.Args[0]), " --selecti \"id\" \"format\" (raw|json|key|value)")
				fmt.Println()
			} else if os.Args[1] == "--total" {
				fmt.Println()
				msg := tar.CountSize()
				fmt.Println("(", msg, ")")
			} else if os.Args[1] == "-selectfx" {
				fmt.Println()
				fmt.Println("ERROR - MISSING ARGUMENTS: \n\n>> ", filepath.Base(os.Args[0]), " --selectfx \"amount\" \"format\" (raw|json|key|value)")
				fmt.Println()
			} else if os.Args[1] == "--selectlx" {
				fmt.Println()
				fmt.Println("ERROR - MISSING ARGUMENTS: \n\n>> ", filepath.Base(os.Args[0]), " --selectlx \"amount\" \"format\" (raw|json|key|value)")
				fmt.Println()
			} else if os.Args[1] == "--search" {
				fmt.Println()
				fmt.Println("ERROR - MISSING ARGUMENTS: \n\n>> ", filepath.Base(os.Args[0]), " --search \"word(s)\" \"format\" (raw|json|key|value)")
				fmt.Println()
			} else if os.Args[1] == "--change" {
				fmt.Println()
				fmt.Println("ERROR - MISSING ARGUMENTS: \n\n>> ", filepath.Base(os.Args[0]), " --change \"id\" \"key data\" \"value date\"")
				fmt.Println()
			} else if os.Args[1] == "--deletef" {
				fmt.Println()
				fmt.Println("ERROR - MISSING ARGUMENTS: \n\n>> ", filepath.Base(os.Args[0]), " --deletef \"id\"")
				fmt.Println()
			} else if os.Args[1] == "--help" {
				fmt.Println(`
--upgrade		"Check for newer version and upgrade the compiled application!"
--createdb		"CREATE new database"
--copydb 		"CREATE backup (copy) of the database"
--deletedb 		"DELETE database"
--deletef 		"DELETE <id> specific row from database"
--search 		"SEARCH <Word(s)> <format> match all words and return results"
--selectf 		"SELECT <format> TOP row from database "
--selectl 		"SELECT <format> LAST row from database"
--selectfx 		"SELECT <number> <format> TOP rows from database"
--selectlx 		"SELECT <number> <format> LAST rows from database"
--selecti 		"SELECT <id> <format> return specific row from database"
--insert 		"INSERT <field one> <field two> for new entry"
--change 		"CHANGE <id> <field one> <field two> on existing row "
--total 		"SHOW number of entries in database"
--version		"SHOW local (App) & (Mod) build version & date`)
				fmt.Println()
			} else if os.Args[1] == "--upgrade" {
				fmt.Println()
				RunUpgrade()
				fmt.Println()
			} else {
				fmt.Println()
				fmt.Println("ERROR - INVALID SYNTAX PROVIDED CHECK MANUAL")
			}
		} else if size == 3 {
			if os.Args[1] == "--createdb" {
				fmt.Println()
				fmt.Println("ERROR - REMOVE:", os.Args[2], "\n\n>> ", filepath.Base(os.Args[0]), " --createdb")
				fmt.Println()
			} else if os.Args[1] == "--upgrade" {
				fmt.Println()
				fmt.Println("ERROR - REMOVE:", os.Args[2], "\n\n>> ", filepath.Base(os.Args[0]), " --upgrade")
				fmt.Println()
			} else if os.Args[1] == "--version" {
				fmt.Println()
				fmt.Println("ERROR - REMOVE:", os.Args[2], "\n\n>> ", filepath.Base(os.Args[0]), " --version")
				fmt.Println()
			} else if os.Args[1] == "--copydb" {
				fmt.Println()
				fmt.Println("ERROR - REMOVE:", os.Args[2], "\n\n>> ", filepath.Base(os.Args[0]), " --copydb")
				fmt.Println()
			} else if os.Args[1] == "--help" {
				fmt.Println()
				fmt.Println("ERROR - REMOVE:", os.Args[2], "\n\n>> ", filepath.Base(os.Args[0]), " --help")
				fmt.Println()
			} else if os.Args[1] == "--deletedb" {
				fmt.Println()
				fmt.Println("ERROR - REMOVE:", os.Args[2], "\n\n>> ", filepath.Base(os.Args[0]), " --deletedb")
				fmt.Println()
			} else if os.Args[1] == "--insert" {
				fmt.Println()
				fmt.Println("ERROR - MISSING ARGUMENTS:", os.Args[1], "\n\n>> ", filepath.Base(os.Args[0]), " --insert \"key filed details\" \"data field details\"")
				fmt.Println()
			} else if os.Args[1] == "--selectf" {
				fmt.Println()
				fmt.Println(tar.FirstField(os.Args[2]))
				fmt.Println()
			} else if os.Args[1] == "--selectl" {
				fmt.Println()
				fmt.Println(tar.LastField(os.Args[2]))
				fmt.Println()
			} else if os.Args[1] == "--selecti" {
				fmt.Println()
				fmt.Println("ERROR - MISSING ARGUMENTS: \n\n>> ", filepath.Base(os.Args[0]), " --selecti \"id\" \"format\" (raw|json|key|value)")
				fmt.Println()
			} else if os.Args[1] == "--total" {
				fmt.Println()
				fmt.Println("ERROR - REMOVE:", os.Args[2], "\n\n>> ", filepath.Base(os.Args[0]), " --total")
				fmt.Println()
			} else if os.Args[1] == "--selectfx" {
				fmt.Println()
				fmt.Println("ERROR - MISSING ARGUMENTS: \n\n>> ", filepath.Base(os.Args[0]), " --selectfx \"amount\" \"format\" (raw|json|key|value)")
				fmt.Println()
			} else if os.Args[1] == "--selectlx" {
				fmt.Println()
				fmt.Println("ERROR - MISSING ARGUMENTS: \n\n>> ", filepath.Base(os.Args[0]), " --selectlx \"amount\" \"format\" (raw|json|key|value)")
				fmt.Println()
			} else if os.Args[1] == "--search" {
				fmt.Println()
				fmt.Println("ERROR - MISSING ARGUMENTS: \n\n>> ", filepath.Base(os.Args[0]), " --search \"word(s)\" \"format\" (raw|json|key|value)")
				fmt.Println()
			} else if os.Args[1] == "--change" {
				fmt.Println()
				fmt.Println("ERROR - MISSING ARGUMENTS: \n\n>> ", filepath.Base(os.Args[0]), " --change \"id\" \"key data\" \"value date\"")
				fmt.Println()
			} else if os.Args[1] == "--deletef" {
				fmt.Println()
				if x, err := strconv.Atoi(os.Args[2]); err == nil {
					msg, status := tar.RemoveField(x)
					fmt.Println(msg, status)
				} else {
					fmt.Println("ERROR - FLAG:(", os.Args[2], ") is not a number!\n\n>> ", filepath.Base(os.Args[0]), " --deletef \"number\"")
				}
			} else {
				fmt.Println()
				fmt.Println("ERROR - INVALID SYNTAX PROVIDED CHECK MANUAL")
				fmt.Println()
			}
		} else if size == 4 {
			if os.Args[1] == "--createdb" {
				fmt.Println()
				fmt.Println("ERROR - REMOVE:", os.Args[2], "AND", os.Args[3], "\n\n>> ", filepath.Base(os.Args[0]), " --createdb")
				fmt.Println()
			} else if os.Args[1] == "--upgrade" {
				fmt.Println()
				fmt.Println("ERROR - REMOVE:", os.Args[2], "AND", os.Args[3], "\n\n>> ", filepath.Base(os.Args[0]), " --upgrade")
				fmt.Println()
			} else if os.Args[1] == "--copydb" {
				fmt.Println()
				fmt.Println("ERROR - REMOVE:", os.Args[2], "AND", os.Args[3], "\n\n>> ", filepath.Base(os.Args[0]), " --copydb")
				fmt.Println()
			} else if os.Args[1] == "--version" {
				fmt.Println()
				fmt.Println("ERROR - REMOVE:", os.Args[2], "AND", os.Args[3], "\n\n>> ", filepath.Base(os.Args[0]), " --version")
				fmt.Println()
			} else if os.Args[1] == "--help" {
				fmt.Println()
				fmt.Println("ERROR - REMOVE:", os.Args[2], "AND", os.Args[3], "\n\n>> ", filepath.Base(os.Args[0]), " --help")
				fmt.Println()
			} else if os.Args[1] == "--deletedb" {
				fmt.Println()
				fmt.Println("ERROR - REMOVE:", os.Args[2], "AND", os.Args[3], "\n\n>> ", filepath.Base(os.Args[0]), " --deletedb")
				fmt.Println()
			} else if os.Args[1] == "--insert" {
				fmt.Println()
				status := tar.AddField(os.Args[2], os.Args[3])
				fmt.Println("returned: (", status, ")")
				fmt.Println()
			} else if os.Args[1] == "--selectf" {
				fmt.Println()
				fmt.Println("ERROR - REMOVE:", os.Args[3], "\n\n>> ", filepath.Base(os.Args[0]), " --selectf \"format\" (raw|json|key|value)")
				fmt.Println()
			} else if os.Args[1] == "--selectl" {
				fmt.Println()
				fmt.Println("ERROR - REMOVE:", os.Args[3], "\n\n>> ", filepath.Base(os.Args[0]), " --selectl \"format\" (raw|json|key|value)")
				fmt.Println()
			} else if os.Args[1] == "--selecti" {
				fmt.Println()
				id, _ := strconv.Atoi(os.Args[2])
				format := os.Args[3]
				fmt.Println(tar.SelectByID(id, format))
				fmt.Println()
			} else if os.Args[1] == "--total" {
				fmt.Println()
				fmt.Println("ERROR - REMOVE:", os.Args[2], "AND", os.Args[3], "\n\n>> ", filepath.Base(os.Args[0]), " --total")
				fmt.Println()
			} else if os.Args[1] == "--search" {
				fmt.Println()
				var format, received = tar.SelectSearch(os.Args[2], os.Args[3])
				bytes := received
				var data []MyStruct
				json.Unmarshal(bytes, &data)

				if (strings.Contains(string(received), "Database") && strings.Contains(string(received), "missing")) || (strings.Contains(string(received), "Database") && strings.Contains(string(received), "empty")) {
					fmt.Println(string(received))
					fmt.Println()
				}

				for x := range data {
					if format == "json" {
						out, _ := json.MarshalIndent(&data[x], "", "  ")
						fmt.Printf("%v", string(out))
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
			} else if os.Args[1] == "--selectfx" {
				format := os.Args[3]
				if x, err := strconv.Atoi(os.Args[2]); err == nil {
					fmt.Println()
					var format, received = tar.FirstXFields(x, format)
					bytes := received
					var data []MyStruct
					json.Unmarshal(bytes, &data)
					if (strings.Contains(string(received), "Database") && strings.Contains(string(received), "missing")) || (strings.Contains(string(received), "Database") && strings.Contains(string(received), "empty")) {
						fmt.Println(string(received))
						fmt.Println()
					}
					for x := range data {
						if format == "json" {
							out, _ := json.MarshalIndent(&data[x], "", "  ")
							fmt.Printf("%v", string(out))
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
					fmt.Println("ERROR - FLAG:(", os.Args[2], ") is not a number!\n\n>> ", filepath.Base(os.Args[0]), " --selectfx \"number\" \"format\" (raw|json|key|value)")
				}
			} else if os.Args[1] == "--selectlx" {
				format := os.Args[3]
				if x, err := strconv.Atoi(os.Args[2]); err == nil {
					fmt.Println()
					var format, received = tar.LastXFields(x, format)

					bytes := received
					var data []MyStruct
					json.Unmarshal(bytes, &data)
					if (strings.Contains(string(received), "Database") && strings.Contains(string(received), "missing")) || (strings.Contains(string(received), "Database") && strings.Contains(string(received), "empty")) {
						fmt.Println(string(received))
						fmt.Println()
					}
					for x := range data {
						if format == "json" {
							out, _ := json.MarshalIndent(&data[x], "", "  ")
							fmt.Printf("%v", string(out))
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
					fmt.Println("ERROR - FLAG:(", os.Args[2], ") is not a number!\n\n>> ", filepath.Base(os.Args[0]), " --selectlx \"number\" \"format\" (raw|json|key|value)")
				}
			} else if os.Args[1] == "--deletef" {
				fmt.Println()
				fmt.Println("ERROR - REMOVE:", os.Args[2], "AND", os.Args[3], "\n\n>> ", filepath.Base(os.Args[0]), " --deletef \"amount\"")
				fmt.Println()
			} else {
				fmt.Println()
				fmt.Println("ERROR - INVALID SYNTAX PROVIDED CHECK MANUAL")
				fmt.Println()
			}
		} else if size == 5 {
			if os.Args[1] == "--change" {
				fmt.Println()
				if x, err := strconv.Atoi(os.Args[2]); err == nil {
					change, status := tar.ModifyField(x, os.Args[3], os.Args[4])
					fmt.Println(change, status)
				} else {
					fmt.Println()
					fmt.Println("ERROR - FLAG:(", os.Args[2], ") is not a number!\n\n>> ", filepath.Base(os.Args[0]), " --change \"<id>\" \"<field 1>\" \" field 2\"")
				}
				fmt.Println()
			} else {
				fmt.Println()
				fmt.Println("ERROR - INVALID SYNTAX PROVIDED CHECK MANUAL")
				fmt.Println()
			}
		} else if size == 6 {
			fmt.Println()
			fmt.Println("ERROR - INVALID SYNTAX PROVIDED CHECK MANUAL")
			fmt.Println()
		}
	}
}
