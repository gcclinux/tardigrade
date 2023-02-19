package main

// Built Sun 19 Feb 20:40:57 GMT 2023

import "log"

// CheckError function takes in a string and the error code!
func CheckError(msg string, err error) {
	if err != nil {
		log.Println(msg)
		panic(err)
	}
}
