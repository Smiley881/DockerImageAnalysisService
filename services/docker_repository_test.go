package services

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"path/filepath"
	"testing"
)

func TestGetManifests_WithTag(t *testing.T) {

	// Array
	input := Input{
		Repository: "dockerhub.timeweb.cloud",
		Name:       "python",
		Tag:        "slim",
	}

	expectedFile, err := os.Open(filepath.Join("..", "resources", "tests", "expected_result_with_tag.json"))
	if err != nil {
		t.Fatalf("Произошла системная ошибка: %v", err)
	}
	expectedJson, err := io.ReadAll(expectedFile)
	if err != nil {
		t.Fatalf("Произошла системная ошибка: %v", err)
	}

	// Act
	result, err := getListManifests(input)

	// Assert
	if err != nil {
		t.Errorf("Произошла ошибка: %v\n", err)
	}

	if string(expectedJson) != result {
		t.Errorf("Получен неверный результат:\nActual:\n%s\nExpected:\n%s\n", result, string(expectedJson))
	}
}

func TestGetManifests_WithoutTag(t *testing.T) {
	// Array
	input := Input{
		Repository: "dockerhub.timeweb.cloud",
		Name:       "python",
		Tag:        "latest",
	}

	expectedFile, err := os.Open(filepath.Join("..", "resources", "tests", "expected_result_without_tag.json"))
	if err != nil {
		t.Fatalf("Произошла системная ошибка: %v", err)
	}
	expectedJson, err := io.ReadAll(expectedFile)
	if err != nil {
		t.Fatalf("Произошла системная ошибка: %v", err)
	}

	// Act
	result, err := getListManifests(input)

	// Assert
	if err != nil {
		t.Errorf("Произошла ошибка: %v\n", err)
	}

	if string(expectedJson) != result {
		t.Errorf("Получен неверный результат:\nActual:\n%s\nExpected:\n%s\n", result, string(expectedJson))
	}

}

func TestGetManifests_IncorrectName(t *testing.T) {

	// Array
	input := Input{
		Repository: "dockerhub.timeweb.cloud",
		Name:       "pytho",
		Tag:        "latest",
	}

	// Act
	result, err := getListManifests(input)

	// Assert
	if result != "" {
		t.Errorf("Получен результат, хотя должна быть пустая строка: %s\n", result)
	}

	if !errors.Is(err, ErrNotFound) {
		t.Errorf("Должна быть получена ошибка Not Found, но получена другая: %v\n", err)
	}
}

func TestGetManifests_IncorrectTag(t *testing.T) {

	// Array
	input := Input{
		Repository: "dockerhub.timeweb.cloud",
		Name:       "python",
		Tag:        "late",
	}

	// Act
	result, err := getListManifests(input)

	// Assert
	if result != "" {
		t.Errorf("Получен результат, хотя должна быть пустая строка: %s\n", result)
	}

	if !errors.Is(err, ErrNotFound) {
		t.Errorf("Должна быть получена ошибка Not Found, но получена другая: %v\n", err)
	}
}

type config struct {
	MediaType string `json:"mediaType"`
	Digest    string `json:"digest"`
	Size      int    `json:"size"`
}

type blob struct {
	Blob config `json:"config"`
}

func TestGetBlobs_Correct(t *testing.T) {

	// Array
	input := Input{
		Repository: "dockerhub.timeweb.cloud",
		Name:       "python",
		Tag:        "latest",
	}

	digest := "sha256:0cea405c3ace86f4480c2986d942fc8258cae70c5ffb1fd70143cb5ba54a208c"

	expectedFile, err := os.Open(filepath.Join("..", "resources", "tests", "expected_result_blobs.json"))
	if err != nil {
		t.Fatalf("Произошла системная ошибка: %v", err)
	}

	expectedByte, err := io.ReadAll(expectedFile)
	if err != nil {
		t.Fatalf("Произошла системная ошибка: %v", err)
	}

	var expectedBlob blob
	var resultBlob blob
	err = json.Unmarshal(expectedByte, &expectedBlob)
	if err != nil {
		t.Fatalf("Произошла системная ошибка: %v", err)
	}

	// Act
	result, err := getBlob(&input, digest)
	errConfig := json.Unmarshal([]byte(result), &resultBlob)

	// Assert
	if errConfig != nil {
		t.Fatalf("Произошла системная ошибка: %v", err)
	}
	if err != nil {
		t.Errorf("Произошла ошибка: %v", err)
	}

	if expectedBlob.Blob.Digest != resultBlob.Blob.Digest || expectedBlob.Blob.MediaType != resultBlob.Blob.MediaType || expectedBlob.Blob.Size != resultBlob.Blob.Size {
		t.Errorf("Получен неверный результат:\nActual:\n%s\nExpected:\n%s\n", string(result), string(expectedByte))
	}
}
