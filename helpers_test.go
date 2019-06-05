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
