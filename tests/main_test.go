package tests

import (
	"HyperFlow/configurations"
	"fmt"
	"testing"
)

func TestConfiguration(t *testing.T) {
	config := configurations.Config{}
	_, err := config.GetConfiguration()

	if err != nil {
		t.Error(err.Error())
	}

	fmt.Println("TEST Success")
}
