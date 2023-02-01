package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Update takes a key and a value if matches DBFile it stores in the config
func Update(k, v string) {
	jsonFile, err := os.Open("db/dbconf.json")
	CheckError("updateconf.go open json --> ", err)
	defer jsonFile.Close()

	decoder := json.NewDecoder(jsonFile)
	conf := config{}

	err = decoder.Decode(&conf)
	CheckError("updateconf.go (decoder.Decode) --> ", err)

	data := map[string]interface{}{}
	if k == "DBFile" && v != "" {
		data = map[string]interface{}{
			k: v,
		}
	}

	myPreJson := make(map[string]string)
	for key, value := range data {
		myPreJson[key] = fmt.Sprintf("%v", value)
	}

	jsonArrVal, _ := json.MarshalIndent(myPreJson, "", "    ")

	file, err := os.Create("db/dbconf.json")
	CheckError("updateconf.go create json --> ", err)
	defer file.Close()

	file.Write(jsonArrVal)
}
