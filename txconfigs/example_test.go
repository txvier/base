package txconfigs

import (
	"fmt"
	"testing"
)

func TestSetConfigLocationForJson(t *testing.T) {
	if err := SetConfigLocation("example.json", AUTOMATIC_ENV_OPTION, WATCH_CONFIG_OPTION); err != nil {
		t.Fatal(err)
	}
	cfg := GetConfigs()
	v := cfg.GetString("APP_NAME")
	fmt.Println(v)
	// output:
	//base in json
}

func TestSetConfigLocationForToml(t *testing.T) {
	if err := SetConfigLocation("example.toml", AUTOMATIC_ENV_OPTION, WATCH_CONFIG_OPTION); err != nil {
		t.Fatal(err)
	}
	cfg := GetConfigs()
	v := cfg.GetString("APP_NAME")
	fmt.Println(v)
	// output:
	//base in toml
}
