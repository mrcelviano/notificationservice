package commons

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

var settings Settings

//Settings - настройки для клиентов
type Settings struct {
	DataBase DataBase `json:"dataBase"`
}

func GetSettings() Settings {
	return settings
}

// ConfigInit считывает конфигурационный файл, находящийся по пути filepath
func ConfigInit(filepath string) {
	log.Println("Reading a config file from a path ", filepath, "...")
	cfg, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Println("File error: ", err.Error())
		os.Exit(1)
	}
	err = json.Unmarshal(cfg, &settings)
	if err != nil {
		log.Println("An error occurred while trying to convert the config file from JSON: ", err.Error())
		os.Exit(1)
	} else {
		log.Println("Configuration file read successfully!")
	}
}
