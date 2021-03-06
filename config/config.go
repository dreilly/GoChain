package config

import (
	"encoding/json"
	"os"
	"os/user"
)

// Configuration : config struct
type Configuration struct {
	ChainDirectory string
	PrintLimit     int32
}

// test

// GetConfiguration : returns a configuration based on a config file
// or will create default values if one does not exist
func (conf *Configuration) GetConfiguration() (config Configuration) {
	homeDir, found := getUserDir()
	file, err := os.Open(homeDir + "/gochain.json")
	if !found || err != nil {
		conf.generateDefaultConfig()
	}

	decoder := json.NewDecoder(file)
	decodeError := decoder.Decode(conf)
	if decodeError != nil {
		conf.generateDefaultConfig()
	}

	return
}

func (conf *Configuration) generateDefaultConfig() (config Configuration) {
	conf.ChainDirectory, _ = getUserDir()
	conf.PrintLimit = 128

	return
}

func getUserDir() (homeDir string, found bool) {
	usr, err := user.Current()
	if err != nil {
		return "", false
	}

	return usr.HomeDir, true
}
