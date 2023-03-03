package main

// Updated Fri  3 Mar 20:53:48 GMT 2023

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func versionCheck(msg string) {
	fmt.Println()
	fmt.Println(msg)
	fmt.Println()
}

func versionRemote() (msg string, upgrade bool) {
	upgrade = false
	remote_version := ""
	file := filepath.Base(os.Args[0])

	lines, err := UrlToLines("https://raw.githubusercontent.com/gcclinux/tardigrade/main/bin/release")
	if err != nil {
		log.Println(err)
	}

	for _, line := range lines {
		remote_version = line
	}

	local_int := strings.Replace(AppGetVersion(), ".", "", -1)
	remote_int := strings.Replace(remote_version, ".", "", -1)

	if local_int == remote_int {
		msg = fmt.Sprintf("Currently \"%v\" is already at the latest version (%v)", file, remote_version)
	} else if local_int > remote_int {
		msg = fmt.Sprintf("Currently \"%v\" (%v) is newer than the online version (%v)", file, AppGetVersion(), remote_version)
	} else if local_int < remote_int {
		msg = fmt.Sprintf("Newer version for \"%v\" (%v) is available for download v(%v)", file, AppGetVersion(), remote_version)
		upgrade = true
	}

	return msg, upgrade
}

func UrlToLines(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return LinesFromReader(resp.Body)
}

func LinesFromReader(r io.Reader) ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
