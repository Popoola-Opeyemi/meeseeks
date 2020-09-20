package core

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// ReadSettings function reads the settings file ...
func ReadSettings(fn string) (*Commands, error) {

	// initializing the settings struct
	settings := &Commands{}

	// opening the file to read
	file, err := os.Open(fn)

	if err != nil {
		return settings, err
	}

	// close file when done
	defer file.Close()

	// reading the contents of the file as bytes
	bytFile, err := ioutil.ReadAll(file)

	if err != nil {
		return settings, err
	}

	// unmarshal the file into settings
	if err := json.Unmarshal(bytFile, &settings); err != nil {
		return settings, err
	}

	return settings, nil
}
