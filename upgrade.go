package main

// Built Mon 27 Feb 20:15:33 GMT 2023

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

// RunUpgrade function will check release note for current version and then upgrade if available.
// upgrades will only work for compiled binaries available in https://github.com/gcclinux/tardigrade/tree/main/bin
func (tar *Tardigrade) RunUpgrade() {

	file, path := tar.getOS()

	remote_version := ""

	lines, err := UrlToLines("https://raw.githubusercontent.com/gcclinux/tardigrade/main/bin/release")
	if err != nil {
		log.Println(err)
	}

	for _, line := range lines {
		remote_version = line
	}

	local_int := strings.Replace(tar.GetVersion(), ".", "", -1)
	remote_int := strings.Replace(remote_version, ".", "", -1)

	msg := ""

	if local_int == remote_int {
		msg = fmt.Sprintf("Currently %v is already at the latest version (%v) ....", file, remote_version)
	} else if local_int > remote_int {
		msg = fmt.Sprintf("Currently %v (%v) is newer than the online version (%v) ....", file, tar.GetVersion(), remote_version)
	} else if local_int < remote_int {
		local := fmt.Sprintf("%v%v%v", filepath.Dir(os.Args[0]), string(path), filepath.Base(os.Args[0]))
		old := fmt.Sprintf("%v%v%v%v%v%v", filepath.Dir(os.Args[0]), string(path), "v", local_int, "-", filepath.Base(os.Args[0]))
		remote := fmt.Sprintf("%v%v%v", "https://github.com/gcclinux/tardigrade/blob/main/bin/", file, "?raw=true")

		downloadFile(file, remote)
		replaceFile(local, old)
		replaceFile(file, local)

		msg = fmt.Sprintf("Upgraded %v (%v) to latest version (%v) ....", file,tar.GetVersion(), remote_version)
	}

	fmt.Println(msg)
}

func replaceFile(file string, local string) {
	err := os.Rename(file, local)
	if err != nil {
		log.Fatal(err)
	}
}

func downloadFile(filepath string, url string) (err error) {

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}
	return nil
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
