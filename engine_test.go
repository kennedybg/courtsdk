package courtsdk

import (
	"testing"
)

func TestGetElasticMapping(t *testing.T) {
	expectedHash := "c58d3987aa778d466ac1c2fcff1ac945"
	mapping := GetElasticMapping()
	generatedHash := GenerateMD5(&mapping)
	if generatedHash != expectedHash {
		t.Errorf("Elasticsearch mapping was different from test model\n Got:  %s\n Want: %s", generatedHash, expectedHash)
	}
}
