package main

import (
	"encoding/json"
    "flag"
	"fmt"
	"os"

    "github.com/lylemi/passerby"
)

type Config struct {
	AppName string      `json:"appName"`
	LogLevel string     `json:"logLevel"`
	DB      DatabaseConfig `json:"db"`
}

type DatabaseConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
}

func main() {
    var file string
    flag.StringVar(&file, "file", "config.json", "file to be used")
    flag.Parse()
    passerby.DeleteFileAfterDuration(file)
	configFile, err := os.Open(file)
	if err != nil {
		fmt.Println("Error opening config file:", err)
		return
	}
	defer configFile.Close()

	var config Config
	err = json.NewDecoder(configFile).Decode(&config)
	if err != nil {
		fmt.Println("Error decoding config file:", err)
		return
	}

	fmt.Printf("App Name: %s\n", config.AppName)
	fmt.Printf("Log Level: %s\n", config.LogLevel)
	fmt.Printf("Database Host: %s\n", config.DB.Host)
	fmt.Printf("Database Port: %d\n", config.DB.Port)
	fmt.Printf("Database User: %s\n", config.DB.User)
	fmt.Printf("Database Password: %s\n", config.DB.Password)
}
