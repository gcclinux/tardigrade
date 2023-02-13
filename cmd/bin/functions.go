package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io"
	"os"
	"strconv"
)

// MyStruct contains the structure of the data stored into the gojsondb.db
type MyStruct struct {
	Id   int    `json:"id"`
	Key  string `json:"key"`
	Data string `json:"data"`
}

// AddField take in (key, sprint) (data, string) and add to gojsondb.db
func AddField(key, data string) bool {
	id := CountSize() + 1
	var getStruct = MyStruct{}
	getStruct.Id = id
	getStruct.Key = key
	getStruct.Data = data

	response, err := json.Marshal(getStruct)
	CheckError("Marshal", err)

	file, err := os.OpenFile(getFile(), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	CheckError("O_APPEND", err)
	file.Write(response)
	file.WriteString("\n")

	return true
}

func RemoveField() {

}

// SelectByID function returns an entry string for a specific id in all formats [ raw | json | id | key | value ]
func SelectByID(id int, f string) string {
	lastLine := 0
	line := ""
	result := ""
	file, err := os.Open(getFile())
	CheckError("CountSize(1)", err)
	defer file.Close()
	var r io.Reader
	r = file
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		lastLine++
		if lastLine == id {
			line = sc.Text()
		}
	}

	var s MyStruct
	in := []byte(line)
	err = json.Unmarshal(in, &s)
	CheckError("LastField(2)", err)

	if f == "json" {
		out, _ := json.MarshalIndent(&s, "", "	")
		result = string(out)
	} else if f == "value" {
		result = string(s.Data)
	} else if f == "raw" {
		result = line
	} else if f == "key" {
		result = string(s.Key)
	} else if f == "id" {
		result = strconv.Itoa(s.Id)
	} else {
		result = "Invalid format provided!"
	}
	return result
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

// FirstField returns the first entry of gojsondb.db in all formats [ raw | json | id | key | value ] specify format required
func FirstField(f string) string {

	lastLine := 0
	line := ""
	file, err := os.Open(getFile())
	result := ""
	CheckError("LastField(1)", err)
	defer file.Close()
	var r io.Reader
	r = file
	sc := bufio.NewScanner(r)

	for sc.Scan() {
		lastLine++
		if lastLine == 1 {
			line = sc.Text()
		}
	}

	var s MyStruct
	in := []byte(line)
	err = json.Unmarshal(in, &s)
	CheckError("LastField(2)", err)

	if f == "json" {
		out, _ := json.MarshalIndent(&s, "", "	")
		result = string(out)
	} else if f == "value" {
		result = string(s.Data)
	} else if f == "raw" {
		result = line
	} else if f == "key" {
		result = string(s.Key)
	} else if f == "id" {
		result = strconv.Itoa(s.Id)
	} else {
		result = "Invalid format provided!"
	}

	return result

}

// LastField returns the last entry of gojsondb.db in all formats [ raw | json | id | key | value ] specify format required
func LastField(f string) string {
	lastLine := 0
	line := ""
	file, err := os.Open(getFile())
	result := ""
	CheckError("LastField(1)", err)
	defer file.Close()
	var r io.Reader
	r = file
	sc := bufio.NewScanner(r)

	for sc.Scan() {
		lastLine++
		if lastLine == CountSize() {
			line = sc.Text()
		}
	}

	var s MyStruct
	in := []byte(line)
	err = json.Unmarshal(in, &s)
	CheckError("LastField(2)", err)

	if f == "json" {
		out, _ := json.MarshalIndent(&s, "", "	")
		result = string(out)
	} else if f == "value" {
		result = string(s.Data)
	} else if f == "raw" {
		result = line
	} else if f == "key" {
		result = string(s.Key)
	} else if f == "id" {
		result = strconv.Itoa(s.Id)
	} else {
		result = "Invalid format provided!"
	}

	return result
}

// EmptyDB - WARNING - this will destroy all data stored in gojsondb.db!
func EmptyDB() bool {
	jsonFile := getFile()

	delete := os.Remove(jsonFile)
	CheckError("EmptyDB(1)", delete)

	if !fileExists(jsonFile) {
		_, err := os.Create(jsonFile)
		CheckError("EmptyDB(2)", err)
		if !fileExists(jsonFile) {
			CheckError("EmptyDB(3)", err)
			return false
		}
	}
	return true
}

// fileExists function will check if the gojsondb.db exists and return true / false
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
