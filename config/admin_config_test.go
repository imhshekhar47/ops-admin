package config

import "testing"

var (
	testConfig *AdminConfiguration = &AdminConfiguration {
		Core: CoreConfiguration {
			Version: "0.0.0",
		},
		Hostname: "localhost",
		Uuid: "localhost",
	}
)


func TestCoreConfig(t *testing.T) {
	if nil == testConfig {
		t.Fail()
	}

	if testConfig.Core.Version != "0.0.0" {
		t.Errorf("Failed to get test configuration")
	}
}