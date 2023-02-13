package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// MyStruct contains the structure of the data stored into the gojsondb.db
type MyStruct struct {
	Id   int    `json:"id"`
	Key  string `json:"key"`
	Data string `json:"data"`
}

// AddField take in (key, sprint) (data, string) and add to gojsondb.db
func AddField(key, data string) {
	id := CountSize() + 1
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
	file.WriteString("\n")
}

func RemoveField() {

}

// SelectID return an entry string for a specific id
func SelectID(id int) string {
	lastLine := 0
	line := ""
	f, err := os.Open(getFile())
	CheckError("CountSize(1)", err)
	defer f.Close()
	var r io.Reader
	r = f
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		lastLine++
		if lastLine == id {
			line = sc.Text()
			return line
		}
	}
	return line
}

func ModifyField() {

}

// CountSize will return number of rows in the gojsondb.db
func CountSize() int {
	f, err := os.Open(getFile())
	CheckError("CountSize(1)", err)
	defer f.Close()
	var r io.Reader
	r = f
	var count int
	const lineBreak = '\n'
	buf := make([]byte, bufio.MaxScanTokenSize)
	for {
		bufferSize, err := r.Read(buf)
		if err != nil && err != io.EOF {
			return 0
		}
		var buffPosition int
		for {
			i := bytes.IndexByte(buf[buffPosition:], lineBreak)
			if i == -1 || bufferSize == buffPosition {
				break
			}
			buffPosition += i + 1
			count++
		}
		if err == io.EOF {
			break
		}
	}
	return count
}

func FirstField() {

}

// LastField read DB file and return the last entry of gojsondb.db
func LastField() string {
	lastLine := 0
	line := ""
	f, err := os.Open(getFile())
	CheckError("CountSize(1)", err)
	defer f.Close()
	var r io.Reader
	r = f
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		lastLine++
		if lastLine == CountSize() {
			line = sc.Text()
			return line
		}
	}
	return line
}

// EmptyDB - WARNING - this will destroy all data stored in gojsondb.db!
func EmptyDB() {
	jsonFile := getFile()

	delete := os.Remove(jsonFile)
	CheckError("EmptyDB(1)", delete)

	if !fileExists(jsonFile) {
		_, err := os.Create(jsonFile)
		CheckError("EmptyDB(2)", err)
		if !fileExists(jsonFile) {
			CheckError("EmptyDB(3)", err)
			return
		}
	}
}

// fileExists function will check if the gojsondb.db exists and return tru / false
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
