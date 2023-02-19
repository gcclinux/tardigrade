package main

const DBFile = "tardigrade.db"
const Release = "0.1.0"
const Updated = "Sun 19 Feb 2023 15:15:12 GMT"

// Tardigrade is the main structure
type Tardigrade struct{}

// GetVersion function returns the current release version
func (tar *Tardigrade) GetVersion() (release string) {
	release = Release
	return release
}

// GetUpdated function returns the last updated time
func (tar *Tardigrade) GetUpdated() (updated string) {
	updated = Updated
	return updated
}
