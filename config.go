package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Config struct {
	Server struct {
		Address string `json:"address"`
	} `json:"server"`
	DB struct {
		Host     string `json:"host"`
		Port     int    `json:"port"`
		User     string `json:"user"`
		Password string `json:"password"`
		Name     string `json:"db_name"`
	} `json:"db"`
}

func ConfigFromFile(path string) Config {
	file, _ := os.Open(path)
	decoder := json.NewDecoder(file)
	config := new(Config)
	err := decoder.Decode(&config)
	if err != nil {
		// handle it
		log.Panicf("Error creating config from file %s \n", path)
	}
	return *config
}

func (this *Config) DbConnectionString() string {
	return fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		this.DB.Host, this.DB.Port, this.DB.User, this.DB.Password, this.DB.Name)
}
