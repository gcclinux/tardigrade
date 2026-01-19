package main

// Updated Fri  3 Mar 20:53:48 GMT 2023

import (
	"fmt"
	"os/exec"
	"runtime"
)

// RunUpgrade function opens browser to latest release page
func RunUpgrade() {
	url := "https://github.com/gcclinux/tardigrade/releases/latest/"
	fmt.Println("Opening browser to download latest version...")
	fmt.Println(url)
	
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	
	if err != nil {
		fmt.Println("Could not open browser. Please visit:")
		fmt.Println(url)
	}
}
