package pkg

import "github.com/spf13/viper"

type Config struct {
	TelegramToken string
	AccessKey     string
}

func Init() (*Config, error) {
	var cnf Config
	if err := parseEnv(&cnf); err != nil {
		return nil, err
	}
	return &cnf, nil
}

func parseEnv(cnf *Config) error {
	if err := viper.BindEnv("token"); err != nil {
		return err
	}
	if err := viper.BindEnv("access_key"); err != nil {
		return err
	}
	cnf.TelegramToken = viper.GetString("token")
	cnf.AccessKey = viper.GetString("access_key")
	return nil
}
