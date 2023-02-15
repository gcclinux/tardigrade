package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io"
	"log"
	"os"
	"runtime"
	"strconv"
)

// MyStruct contains the structure of the data stored into the gojsondb.db
type MyStruct struct {
	Id   int    `json:"id"`
	Key  string `json:"key"`
	Data string `json:"data"`
}

func getOS() rune {

	PATH_SEPARATOR := '/'

	if runtime.GOOS == "windows" {
		PATH_SEPARATOR = '\\'
	} else if runtime.GOOS == "linux" {
		PATH_SEPARATOR = '/'
	} else if runtime.GOOS == "darwin" {
		PATH_SEPARATOR = '/'
	} else {
		log.Println("unknown")
	}

	return PATH_SEPARATOR
}

// AddField take in (key, sprint) (data, string) and add to gojsondb.db
func AddField(key, data string) bool {

	if !fileExists(getFile()) {
		CreateDB()
		if !fileExists(getFile()) {
			return false
		}
	}

	id := UniqueID() + 1
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

// RemoveField function takes an unique field id as an input and remove the matching field entry
func RemoveField(id int) bool {

	line := SelectByID(id, "raw")

	//CreatedDBCopy()
	// dirname, err := os.UserHomeDir()
	// CheckError("RemoveField(0)", err)
	// src := "gojsontmp.db"
	// fpath := fmt.Sprintf("%s%s%s", dirname, string(getOS()), src)

	fpath := getFile()
	f, err := os.Open(fpath)
	CheckError("RemoveField(1)", err)

	var bs []byte
	buf := bytes.NewBuffer(bs)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if scanner.Text() != line {
			_, err := buf.Write(scanner.Bytes())
			CheckError("RemoveField(2)", err)
			_, err = buf.WriteString("\n")
			CheckError("RemoveField(3)", err)
		}
	}
	if err := scanner.Err(); err != nil {
		CheckError("RemoveField(4)", err)
	}

	err = os.WriteFile(fpath, buf.Bytes(), 0666)
	CheckError("RemoveField(5)", err)
	f.Close()

	return true
}

// SelectByID function returns an entry string for a specific id in all formats [ raw | json | id | key | value ]
func SelectByID(id int, f string) string {
	lastLine := 0
	line := ""
	result := ""
	file, err := os.Open(getFile())
	CheckError("SelectByID(1)", err)
	defer file.Close()
	var r io.Reader = file
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

// ModifyField takes ID, Key, Value (all 3 fields) and update with information provided in k & V
func ModifyField(id int, k, v string) bool {

	//line := SelectByID(id, "raw")

	return false
}

// CountSize will return number of rows in the gojsondb.db
func CountSize() int {
	f, err := os.Open(getFile())
	CheckError("CountSize(1)", err)
	defer f.Close()
	var r io.Reader = f
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

// UniqueID function returns an int for the last used UniqueID to AutoIncrement in the AddField()
func UniqueID() int {
	lastID, err := strconv.Atoi(LastField("id"))
	CheckError("UniqueID()", err)
	return lastID
}

// FirstXFields returns first X number of entries from database in byte[] format
//
// Example:
// specify number of fields to return FirstXFields(2)
func FirstXFields(count int) []byte {

	var allRecords []MyStruct
	xFields := new(MyStruct)
	var tmpStruct MyStruct
	lastLine := 0
	start := 1
	end := count
	line := ""

	file, err := os.Open(getFile())
	CheckError("FirstXFields(1)", err)

	defer file.Close()
	var r io.Reader = file
	sc := bufio.NewScanner(r)

	for sc.Scan() {
		lastLine++
		if lastLine >= start && lastLine <= end {
			line = sc.Text()
			in := []byte(line)

			err = json.Unmarshal(in, &tmpStruct)
			CheckError("FirstXFields(2)", err)

			xFields.Id = tmpStruct.Id
			xFields.Key = string(tmpStruct.Key)
			xFields.Data = string(tmpStruct.Data)
			allRecords = append(allRecords, *xFields)
		}
	}

	allRecord, err := json.Marshal(allRecords)
	CheckError("FirstXFields(3)", err)

	return allRecord
}

// LastXFields returns last X numbers of entries from db in byte[] format
//
// Example:
// specify number of fields to return LastXFields(2)
func LastXFields(count int) []byte {

	var allRecords []MyStruct
	count = count - 1
	xFields := new(MyStruct)
	var tmpStruct MyStruct
	lastLine := 0
	start := CountSize() - count
	end := CountSize()
	line := ""

	file, err := os.Open(getFile())
	CheckError("LastXFields(1)", err)

	defer file.Close()
	var r io.Reader = file
	sc := bufio.NewScanner(r)

	for sc.Scan() {
		lastLine++
		if lastLine >= start && lastLine <= end {
			line = sc.Text()
			in := []byte(line)

			err = json.Unmarshal(in, &tmpStruct)
			CheckError("LastXFields(2)", err)

			xFields.Id = tmpStruct.Id
			xFields.Key = string(tmpStruct.Key)
			xFields.Data = string(tmpStruct.Data)
			allRecords = append(allRecords, *xFields)
		}
	}

	allRecord, err := json.Marshal(allRecords)
	CheckError("LastXFields(3)", err)

	return allRecord
}

// FirstField returns the first entry in the database in all formats [ raw | json | id | key | value ],
// must specify format required Example: FirstField("json")
func FirstField(f string) string {

	lastLine := 0
	line := ""
	file, err := os.Open(getFile())
	result := ""
	CheckError("LastField(1)", err)
	defer file.Close()
	var r io.Reader = file
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

// LastField returns the last entry of the database in all formats [ raw | json | id | key | value ] specify format required
func LastField(f string) string {
	lastLine := 0
	line := ""
	file, err := os.Open(getFile())
	result := ""
	CheckError("LastField(1)", err)
	defer file.Close()
	var r io.Reader = file
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
