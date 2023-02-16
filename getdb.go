package tardigrade

import (
	"encoding/json"
	"os"
)

// Tardigrade is the main structure
type Tardigrade struct{}

type config struct {
	DBFile  string
	Release string
	Updated string
}

// getFile fucntion returns the database path
func (tar *Tardigrade) getFile() (dbFile string) {
	file, err := os.Open("dbconf.json")
	CheckError("Accessing dbconf.json -->", err)
	defer file.Close()

	decoder := json.NewDecoder(file)
	conf := config{}
	err = decoder.Decode(&conf)
	CheckError("getdb.go decoder.Decode", err)

	dbFile = conf.DBFile

	return dbFile
}

// GetVersion function returns the current release version
func (tar *Tardigrade) GetVersion() (release string) {
	file, err := os.Open("dbconf.json")
	CheckError("Accessing dbconf.json -->", err)
	defer file.Close()

	decoder := json.NewDecoder(file)
	conf := config{}
	err = decoder.Decode(&conf)
	CheckError("getdb.go decoder.Decode", err)

	release = conf.Release

	return release
}

// GetUpdated function returns the last updated time
func (tar *Tardigrade) GetUpdated() (updated string) {
	file, err := os.Open("dbconf.json")
	CheckError("Accessing dbconf.json -->", err)
	defer file.Close()

	decoder := json.NewDecoder(file)
	conf := config{}
	err = decoder.Decode(&conf)
	CheckError("getdb.go decoder.Decode", err)

	updated = conf.Updated

	return updated
}
