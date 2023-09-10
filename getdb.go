package main

const Release = "1.0.2"
const Updated = "Sun 10 Sep 19:49:18 BST 2023"

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
