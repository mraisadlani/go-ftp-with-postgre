package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Configure struct {
	FTP SetupFTP
	App SetupApp
	Database SetupDatabase
}

type SetupFTP struct {
	FTPUSER string `mapstructure:"FTP_USER"`
	FTPPASS string `mapstructure:"FTP_PASS"`
	FTPHOST string `mapstructure:"FTP_HOST"`
	FTPPORT int `mapstructure:"FTP_PORT"`
}

type SetupApp struct {
	Host string `mapstructure:"HOST"`
	Port int `mapstructure:"PORT"`
}

type SetupDatabase struct {
	DBUSER string `mapstructure:"DB_USER"`
	DBPASS string `mapstructure:"DB_PASS"`
	DBPORT string `mapstructure:"DB_PORT"`
	DBHOST string `mapstructure:"DB_HOST"`
	DBNAME string `mapstructure:"DB_NAME"`
}

var (
	C *Configure
)

func SetupConfiguration() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading config file", err)
		return
	}

	C = new(Configure)
	err := viper.Unmarshal(&C)

	if err != nil {
		log.Fatal("Unable to decode into struct", err)
		return
	}
}
