package helper

import "github.com/spf13/viper"

type EnvConfig struct {
	Port                   string `mapstructure:"PORT"`
	DBSource               string `mapstructure:"DB_SOURCE"`
	MigrationURL           string `mapstructure:"MIGRATION_URL"`
	GoogleClientID         string `mapstructure:"GOOGLE_OAUTH_CLIENT_ID"`
	GoogleClientSecret     string `mapstructure:"GOOGLE_OAUTH_CLIENT_SECRET"`
	GoogleOAuthRedirectUrl string `mapstructure:"GOOGLE_OAUTH_REDIRECT_URL"`
	FrontEndOrigin string `mapstructure:"FRONTEND_ORIGIN"`
}

func LoadEnv(path string) (env EnvConfig, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&env)
	return
}
