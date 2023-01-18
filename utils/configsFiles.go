package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Config struct {
	DatabaseConnection string `json:"databaseConnection"`
	Port               string `json:"port"`
}

func ReadConfigs() Config {
	file, _ := ioutil.ReadFile("./configs/database.json")

	data := Config{}

	err := json.Unmarshal([]byte(file), &data)
	if err != nil {
		log.Panic(err)
	}
	return data
}
