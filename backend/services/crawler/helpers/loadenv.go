package helpers

import "github.com/spf13/viper"


type EnvConfig struct {
	GRPCPort string `mapstructure:"GRPC_PORT"`
	DBSource               string `mapstructure:"DB_SOURCE"`
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