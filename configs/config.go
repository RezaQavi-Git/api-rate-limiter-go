package configs

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
	"strings"
)

func LoadConfig() *Config {
	var configFilePath = GetOsEnv("CONFIG_FILE_PATH", "configs/")
	var configFileName = GetOsEnv("CONFIG_FILE_NAME", "config")

	viper.AddConfigPath(configFileName)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(configFilePath)
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("Couldn't load config from %s%s with error: %s", configFilePath, configFileName, err))
	}
	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		panic(fmt.Sprintf("Couldn't load config from %s%s with error: %s", configFilePath, configFileName, err))
	}
	log.Printf("Successfully loaded config from %s%s", configFilePath, configFileName)

	return &config
}

func GetOsEnv(key string, defaultValue string) string {
	val := os.Getenv(key)
	if val == "" {
		val = defaultValue
	}
	return val
}
