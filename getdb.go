package main

// Built Mon 27 Feb 20:15:33 GMT 2023

const DBFile = "tardigrade.db"
const Release = "0.1.5"
const Updated = "Mon 27 Feb 22:58:49 GMT 2023"

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
