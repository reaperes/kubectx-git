package config

import (
	"fmt"
	"path/filepath"
	"testing"
)

func Test_ReadConfig(t *testing.T) {
	testConfig := filepath.Join("..", "..", "testdata", "config")
	data, _ := ReadConfig(testConfig)
	fmt.Print(data)
	if len(data) == 0 {
		t.Errorf("Failed to read config file")
	}
}
