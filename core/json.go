package core

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// ReadConfig function reads the settings file ...
func ReadConfig(fn string) (*Commands, error) {

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

// FileExist ...
func FileExist(fn string) bool {

	if _, err := os.Stat(fn); err == nil {

		// path/to/whatever exists
		return true

	} else if os.IsNotExist(err) {

		// path/to/whatever does *not* exist
		return false

	} else {

		return false
	}
}

// CreateConfig create config file ...
func CreateConfig(fn string) error {

	// config structure
	config := []byte(`
	{
		"commands": {
			"concurrent": true,
			"list": [
				{
					"directory": "",
					"concurrent": true,
					"list": [{ "cmd": ""}]
				},
				{
					"directory": "",
					"concurrent": true,
					 "list": [{ "cmd": ""}]
				},
				{
					"directory": "",
					"concurrent": true,
					 "list": [{ "cmd": ""}]
				}
			]
		}
	}
	`)

	configStruct := Commands{}

	err := json.Unmarshal(config, &configStruct)

	if err != nil {
		return err
	}

	cJSON, _ := json.Marshal(configStruct)

	err = ioutil.WriteFile(fn, cJSON, 0644)

	if err != nil {
		return err
	}

	return nil

}
