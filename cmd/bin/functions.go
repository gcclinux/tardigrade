package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type MyStruct struct {
	Id   int    `json:"id"`
	Key  string `json:"key"`
	Data string `json:"data"`
}

func GetJsonDB() {

	jsonFile := getFile()
	if !fileExists(jsonFile) {
		fmt.Println("Database does not exist")
		_, err := os.Create(jsonFile)
		CheckError("GetJsonDB(1)", err)
		if !fileExists(jsonFile) {
			CheckError("GetJsonDB(2)", err)
			return
		}
	}
}

func AddField(id int, key, data string) {
	var getStruct = MyStruct{}
	getStruct.Id = id
	getStruct.Key = key
	getStruct.Data = data

	response, err := json.Marshal(getStruct)
	CheckError("Marshal", err)
	fmt.Println(string(response))

	file, err := os.OpenFile(getFile(), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	CheckError("O_APPEND", err)
	file.Write(response)
}

func RemoveField() {

}

func SelectField() {

}

func ModifyField() {

}

func CountSize() {

}

func FirstField() {

}

func LastField() {
	jsonFile := getFile()
	readFile(jsonFile)
}

func EmptyDB() {

}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func ReturnDBfFile() *os.File {
	file, err := os.Open(getFile())
	if err != nil {
		fmt.Println("ReturnDBfFile()", err)
	}
	return file
}

func readFile(fname string) {
	file, err := os.Open(fname)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	buf := make([]byte, 62)
	stat, err := os.Stat(fname)
	start := stat.Size() - 62
	_, err = file.ReadAt(buf, start)
	if err == nil {
		fmt.Printf("%s\n", buf)
	}

}
