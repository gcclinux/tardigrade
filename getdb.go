package main

import (
	_ "embed"
	"encoding/json"
)

//go:embed release.json
var releaseData []byte

type ReleaseInfo struct {
	Version string `json:"version"`
	Updated string `json:"updated"`
}

var Release string
var Updated string

func init() {
	var info ReleaseInfo
	if err := json.Unmarshal(releaseData, &info); err == nil {
		Release = info.Version
		Updated = info.Updated
	} else {
		Release = "1.1.0"
		Updated = "Sun Jan 18 09:38:18 PM GMT 2026"
	}
}

// GetVersion function returns the current release version
func AppGetVersion() (release string) {
	release = Release
	return release
}

// GetUpdated function returns the last updated time
func AppGetUpdated() (updated string) {
	updated = Updated
	return updated
}
