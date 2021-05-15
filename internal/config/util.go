package config

import (
	"fmt"
	"github.com/bharat-rajani/GoHunter/internal/helpers"
	"github.com/spf13/viper"
)

func ReadConfig() (*Configuration, error) {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(helpers.CONFIG_FILE_PATH)

	var config Configuration
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("CONFIG NOT FOUND")
			// TODO: Create Default config
			return nil, err
		} else {
			return nil, err
		}
	}

	err := viper.Unmarshal(&config)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
		return nil, err
	}

	return &config, nil
}
