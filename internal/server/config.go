package server

import "github.com/spf13/viper"

type Config struct {
	Port int `mapstructure:"serverPort"`
}

func NewConfig(env string) (*Config, error) {
	viper.SetConfigFile("config/" + env + ".json")
	viper.SetConfigType("json")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	config := &Config{}
	err = viper.Unmarshal(config)
	return config, err
}
