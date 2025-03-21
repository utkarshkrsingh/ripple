package config

import (
	"errors"
	"strings"

	"github.com/spf13/viper"
)

var Config *viper.Viper
var Sections = map[string]bool{}

func InitReadConfig() error {

    Config = viper.New()
    Config.SetConfigName("ripple")
    Config.SetConfigType("toml")
    Config.AddConfigPath(".")

    if err := Config.ReadInConfig(); err != nil {
        return errors.New("unable to read ripple.toml")
    }

    GetAllSections()

    return nil
}

func GetAllSections() {
    for _, k := range Config.AllKeys() {
        splitString := strings.Split(k, ".")
        Sections[splitString[0]] = true
    }
}
