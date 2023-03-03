package main

const DBFile = "tardigrade.db"
const Release = "1.0.1"
const Updated = "Fri  3 Mar 20:53:48 GMT 2023"

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
