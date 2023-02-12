package main

import "log"

// CheckError receive a message and the error code, if there is a problem panic and display issues
func CheckError(msg string, err error) {
	if err != nil {
		log.Println(msg)
		panic(err)
	}
}
