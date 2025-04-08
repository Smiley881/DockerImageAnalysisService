package services

import (
	"os"
	"path/filepath"
	"testing"
)

func TestGetManifestLinuxAmd64(t *testing.T) {

	// Array
	input, err := os.Open(filepath.Join("..", "resources", "tests", "expected_result_with_tag.json"))
	if err != nil {
		t.Fatalf("Произошла системная ошибка: %v", err)
	}

	var manifests Manifests
	err = parseJsonToStructManifests(input, &manifests)
	if err != nil {
		t.Fatalf("Произошла системная ошибка: %v", err)
	}

	expectedResult := "sha256:0cea405c3ace86f4480c2986d942fc8258cae70c5ffb1fd70143cb5ba54a208c"

	// Act
	result, err := getManifestLinuxAmd64(&manifests)

	// Assert
	if err != nil {
		t.Errorf("Произошла ошибка: %v", err)
	}

	if result != expectedResult {
		t.Errorf("Ожидалось - %s, а получено - %s", expectedResult, result)
	}
}
