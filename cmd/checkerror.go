package main

import "log"

func CheckError(msg string, err error) {
	if err != nil {
		log.Println(msg)
		panic(err)
	}
}
