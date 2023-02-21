package main

// Built Tue 21 Feb 22:05:28 GMT 2023

const DBFile = "tardigrade.db"
const Release = "0.1.2"
const Updated = "Tue 21 Feb 22:05:28 GMT 2023"

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
