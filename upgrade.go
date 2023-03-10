package main

// Updated Fri  3 Mar 20:53:48 GMT 2023

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/gcclinux/tardigrade-mod"
)

// RunUpgrade function will check release note for current version and then upgrade if available.
// upgrades will only work for compiled binaries available in https://github.com/gcclinux/tardigrade/tree/main/bin
func RunUpgrade() {

	tar := tardigrade.Tardigrade{}
	path := tar.GetOS()
	file := getFile()
	remote_version := ""

	local_int := strings.Replace(AppGetVersion(), ".", "", -1)
	msg, upgrade := versionRemote()

	if upgrade {
		local := fmt.Sprintf("%v%v%v", filepath.Dir(os.Args[0]), string(path), filepath.Base(os.Args[0]))
		old := fmt.Sprintf("%v%v%v%v%v%v", filepath.Dir(os.Args[0]), string(path), "v", local_int, "-", filepath.Base(os.Args[0]))
		remote := fmt.Sprintf("%v%v%v", "https://github.com/gcclinux/tardigrade/blob/main/bin/", file, "?raw=true")

		downloadFile(file, remote)
		replaceFile(local, old)
		replaceFile(file, local)

		msg = fmt.Sprintf("Upgraded %v (%v) to latest version (%v) ....", file, AppGetVersion(), remote_version)
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

func getFile() string {
	FILE_NAME := ""
	if runtime.GOOS == "windows" && runtime.GOARCH == "amd64" {
		FILE_NAME = "tardigrade-win-x86_64.exe"
	} else if runtime.GOOS == "linux" && runtime.GOARCH == "amd64" {
		FILE_NAME = "tardigrade-Linux-x86_64"
	} else if runtime.GOOS == "linux" && runtime.GOARCH == "arm64" {
		FILE_NAME = "tardigrade-Linux-aarch64"
	} else if runtime.GOOS == "darwin" && runtime.GOARCH == "arm64" {
		FILE_NAME = "tardigrade-Darwin-arm64"
	} else {
		log.Println("unknown")
	}
	return FILE_NAME
}
