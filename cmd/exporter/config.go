package main

import (
	"os"
	"log"
	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	ChassisList []interface{}
	Username string
	Password string
}

func GetConf(file string) (*Config, error) {
	obj := make(map[string]any)

	yamlFile, err := os.ReadFile(file)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	err = yaml.Unmarshal(yamlFile, obj)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return &Config{
		ChassisList: obj["chassis"].([]interface{}),
		Username: obj["username"].(string),
		Password: obj["password"].(string),
	}, nil
}
