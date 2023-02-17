package main

const DBFile = "tardigrade.db"
const Release = "0.0.5"
const Updated = "Fri Feb 17 22:41:31 GMTST 2023"

// Tardigrade is the main structure
type Tardigrade struct{}

type config struct {
	DBFile  string
	Release string
	Updated string
}

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
