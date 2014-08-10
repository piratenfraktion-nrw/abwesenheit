package configuration

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

const (
	CONFIG_FILENAME = "config.json"
)

var Config = Configuration{}

type Configuration struct {
	Ldap  Ldap   `json:"ldap"`
	MySql MySql  `json:"mysql"`
	Port  string `json:port`
}

type Ldap struct {
	Hostname string `json:"hostname"`
	Port     string `json:"port"`
}

type MySql struct {
	Hostname string `json:"hostname"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
}

func Load() error {

	configFile, err := os.Open(CONFIG_FILENAME)

	if err != nil {
		log.Fatal(err)
		return err
	}

	defer configFile.Close()

	decoder := json.NewDecoder(configFile)
	decoder.Decode(&Config)

	//log.Println(Config)
	return nil
}

func Save() error {

	configFile, err := os.Create(CONFIG_FILENAME)

	if err != nil {
		log.Fatal(err)
		return err
	}

	defer configFile.Close()

	jsonConfig, err := json.Marshal(Config)

	if err != nil {
		log.Fatal(err)
		return err
	}

	n, err := io.WriteString(configFile, string(jsonConfig))

	if err != nil {
		log.Fatal(string(n), err.Error())
		return err
	}

	return nil
}
