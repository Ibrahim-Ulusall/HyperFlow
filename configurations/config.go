package configurations

import (
	"HyperFlow/configurations/models"
	"fmt"

	"github.com/spf13/viper"
)

var path string = "E:\\go\\configurations\\"

type Config struct{}

func (config *Config) ConfigureApp() (*models.Settings, error) {
	viper := viper.New()
	viper.AddConfigPath(path)
	viper.SetConfigName("settings")
	viper.SetConfigType("json")

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("configuration file read error: %w", err)
	}
	var setting models.Settings
	if err := viper.Unmarshal(&setting); err != nil {
		return nil, fmt.Errorf("unable to unmarshal config: %w", err)
	}
	return &setting, nil
}
