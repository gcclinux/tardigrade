package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"strconv"
	"time"
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
		emptyDB()
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

// createdDBCopy creates a copy of the Database before RemoveField() runs to campture all error or issues
func CreatedDBCopy() bool {

	dirname, err := os.UserHomeDir()
	CheckError("CreatedDBCopy(0)", err)
	target := "gojsontmp.db"

	src := getFile()
	fin, err := os.Open(src)
	CheckError("CreatedDBCopy(1)", err)
	defer fin.Close()

	dst := fmt.Sprintf("%s%s%s", dirname, string(getOS()), target)
	buf := make([]byte, 1024)
	tmp, err := os.Create(dst)
	CheckError("CreatedDBCopy(2)", err)
	defer tmp.Close()

buffering:
	for {
		n, err := fin.Read(buf)
		if err != nil && err != io.EOF {
			CheckError("CreatedDBCopy(3)", err)
			return false
		}

		if n == 0 {
			fin.Close()
			tmp.Close()
			break buffering
		}

		if _, err := tmp.Write(buf[:n]); err != nil {
			CheckError("CreatedDBCopy(4)", err)
			return false
		}
	}
	fmt.Println(dst)
	return true
}

// RemoveField - WARNING: takes a REGEX input and remove add matching string (Can not be undone)
func RemoveField(text string) bool {

	fmt.Println("Check String:", text)

	CreatedDBCopy()
	dirname, err := os.UserHomeDir()
	CheckError("RemoveField(0)", err)
	src := "gojsontmp.db"

	fpath := fmt.Sprintf("%s%s%s", dirname, string(getOS()), src)
	f, err := os.Open(fpath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var bs []byte
	buf := bytes.NewBuffer(bs)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if scanner.Text() != ":9," {
			fmt.Println("Not match:", scanner.Text())
			_, err := buf.Write(scanner.Bytes())
			if err != nil {
				log.Fatal(err)
			}
			_, err = buf.WriteString("\n")
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(fpath, buf.Bytes(), 0666)
	if err != nil {
		log.Fatal(err)
	}

	return false
}

// SelectByID function returns an entry string for a specific id in all formats [ raw | json | id | key | value ]
func SelectByID(id int, f string) string {
	lastLine := 0
	line := ""
	result := ""
	file, err := os.Open(getFile())
	CheckError("CountSize(1)", err)
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

func ModifyField() bool {
	//TODO
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

// LastXFields returns last X number of entries in gojsondb.db in all formats [ raw | json | id | key | value ]
// specify format and number of fileds to return LastXFields(10,"raw")
func LastXFields(count int, f string) []MyStruct {

	var tmpRecords MyStruct
	var allRecords []MyStruct
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
			err = json.Unmarshal(in, &tmpRecords)
			CheckError("LastXFields(2)", err)

			if f == "json" {
				allRecords = append(allRecords, tmpRecords)
			}
			// else if f == "value" {
			// 	result = string(tmpRecords.Data)
			// 	fmt.Println(result)
			// } else if f == "raw" {
			// 	result = line
			// 	fmt.Println(result)
			// } else if f == "key" {
			// 	result = string(tmpRecords.Key)
			// 	fmt.Println(result)
			// } else if f == "id" {
			// 	result = strconv.Itoa(tmpRecords.Id)
			// 	fmt.Println(result)
			// } else {
			// 	result = "Invalid format provided!"
			// 	fmt.Println(result)
			// }
		}
	}

	return allRecords

}

// FirstField returns the first entry of gojsondb.db in all formats [ raw | json | id | key | value ]
// specify format required FirstField("json")
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

// LastField returns the last entry of gojsondb.db in all formats [ raw | json | id | key | value ] specify format required
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

// EmptyDB - WARNING - this will destroy all data stored in gojsondb.db!
func emptyDB() bool {
	jsonFile := getFile()

	if fileExists(jsonFile) {
		delete := os.Remove(jsonFile)
		CheckError("emptyDB(1)", delete)
	}

	if !fileExists(jsonFile) {
		_, err := os.Create(jsonFile)
		CheckError("emptyDB(2)", err)
		if !fileExists(jsonFile) {
			CheckError("emptyDB(3)", err)
			return false
		}
	}
	time.Sleep(2000)
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
