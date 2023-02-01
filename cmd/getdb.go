package main

import (
	"encoding/json"
	"os"
)

type config struct {
	DBFile string
}

func getConf() (dbFile string) {
	file, err := os.Open("db/dbconf.json")
	CheckError("Accessing dbconf.json -->", err)
	defer file.Close()

	decoder := json.NewDecoder(file)
	conf := config{}
	err = decoder.Decode(&conf)
	CheckError("getdb.go decoder.Decode", err)

	dbFile = conf.DBFile

	return dbFile
}
