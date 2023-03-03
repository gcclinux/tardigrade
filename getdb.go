package main

const DBFile = "tardigrade.db"
const Release = "1.0.0"
const Updated = "Fri 03 Mar 12:15:49 GMT 2023"

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
