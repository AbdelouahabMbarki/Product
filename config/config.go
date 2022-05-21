package config

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/viper"
)

type Configurations struct {
	Debug    string
	Server   ServerConfigurations
	Database DatabaseConfigurations
}

type ServerConfigurations struct {
	Port string `yaml:"port"`
}
type DatabaseConfigurations struct {
	Url string `yaml:"url"`
}

func LoadConfig(path string) (config Configurations, err error) {
	viper.SetConfigName("config")
	viper.AddConfigPath(path)
	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()
	viper.SetConfigType("yml")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}
	// Set undefined variables
	err = viper.Unmarshal(&config)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}
	spew.Dump(config)
	return
}
