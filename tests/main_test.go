package tests

import (
	"HyperFlow/configurations"
	"testing"
)

func TestConfiguration(t *testing.T) {
	config := configurations.Config{}
	_, err := config.ConfigureApp()

	if err != nil {
		t.Error(err.Error())
	}
}
