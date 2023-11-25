package config

import "github.com/spf13/viper"

type Config struct {
	DBConnectionString string
	DBName             string
	JokesURL           string
	JokesLimit         int
	JokesTimeout       int
	BindAddress        string
}

func GetConfig() (Config, error) {
	viper.AutomaticEnv()

	return Config{
		DBConnectionString: viper.GetString("DB_CONNECTION_STRING"),
		DBName:             viper.GetString("DB_NAME"),
		JokesURL:           viper.GetString("JOKES_URL"),
		JokesLimit:         viper.GetInt("JOKES_LIMIT"),
		JokesTimeout:       viper.GetInt("JOKES_TIMEOUT"),
		BindAddress:        viper.GetString("BIND_ADDRESS"),
	}, nil
}
