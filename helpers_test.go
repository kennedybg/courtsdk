package courtsdk

import (
	"os"
	"strconv"
	"testing"
)

func TestGetEnvInt(t *testing.T) {
	testValue := 1
	os.Setenv("TEST_ENV_INT_VALUE", strconv.Itoa(testValue))
	envValue := GetEnvInt("TEST_ENV_INT_VALUE", 0)
	if envValue != testValue {
		t.Errorf("Env var was incorrect\n Got:  %d\n Want: %d", envValue, testValue)
	}
}

func TestGetEnvString(t *testing.T) {
	testValue := "TST"
	os.Setenv("TEST_ENV_STR_VALUE", testValue)
	envValue := GetEnvString("TEST_ENV_STR_VALUE", "TST")
	if envValue != testValue {
		t.Errorf("Env var was incorrect\n Got:  %s\n Want: %s", envValue, testValue)
	}
}

func TestGenerateMD5(t *testing.T) {
	expectedHash := "86fb269d190d2c85f6e0468ceca42a20"
	phrase := "Hello world!"
	generatedHash := GenerateMD5(&phrase)
	if generatedHash != expectedHash {
		t.Errorf("MD5 checksum different from expected\n Got:  %s\n Want: %s", generatedHash, expectedHash)
	}
}

func TestGetElasticMapping(t *testing.T) {
	expectedHash := "c58d3987aa778d466ac1c2fcff1ac945"
	mapping := GetElasticMapping()
	generatedHash := GenerateMD5(&mapping)
	if generatedHash != expectedHash {
		t.Errorf("Elasticsearch mapping was different from test model\n Got:  %s\n Want: %s", generatedHash, expectedHash)
	}
}
