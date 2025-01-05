package configurations

import (
	"HyperFlow/configurations/models"
	"encoding/json"
	"errors"
	"io"
	"log"
	"os"
)

var configurationFileName string = "configurations/settings.json"

type Config struct{}

func (config *Config) GetConfiguration() (*models.Settings, error) {
	if _, err := os.Stat(configurationFileName); os.IsNotExist(err) {
		log.Printf("%s dosyası bulunamadı", configurationFileName)
		return nil, errors.New("config dosyası bulunamadı")
	}

	file, err := os.Open(configurationFileName)
	if err != nil {
		log.Printf("Dosya açılırken hata oluştu: %v", err)
		return nil, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		log.Printf("Dosya okuma hatası: %v", err)
		return nil, err
	}

	var settings models.Settings
	err = json.Unmarshal(data, &settings)
	if err != nil {
		log.Printf("JSON çözümleme hatası: %v", err)
		return nil, err
	}

	return &settings, nil
}
