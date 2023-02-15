package main

import (
	"fmt"
	"io"
	"os"
)

// CreatedDBCopy creates a copy of the Database and store in UserHomeDir()
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
	fmt.Println("PATH: ", dst)
	return true
}

// CreateDB - This function will create a database file if it does not exist and return true | false
func CreateDB() bool {
	jsonFile := getFile()

	if !fileExists(jsonFile) {
		_, err := os.Create(jsonFile)
		CheckError("CreateDB(2)", err)
		if !fileExists(jsonFile) {
			CheckError("CreateDB(3)", err)
			return false
		}
	}
	return true
}

// DeleteDB - WARNING - this function delete the database file return true | false
func DeleteDB() bool {
	jsonFile := getFile()

	if fileExists(jsonFile) {
		delete := os.Remove(jsonFile)
		CheckError("DeleteDB(1)", delete)
		if fileExists(jsonFile) {
			return false
		}
	}
	return true
}

// fileExists function will check if the database exists and return true / false
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// EmptyDB function - WARNING - this will destroy the database and all data stored in it!
func EmptyDB() bool {

	status := true

	if DeleteDB() {
		if !CreateDB() {
			status = false
		}
	} else {
		status = false
	}
	return status
}
