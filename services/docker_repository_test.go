package services

import (
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
		t.Fatalf("Произошла ошибка: %v", err)
	}
	expectedJson, err := io.ReadAll(expectedFile)
	if err != nil {
		t.Fatalf("Произошла ошибка: %v", err)
	}

	// Act
	result, err := getImageManifests(input)

	// Assert
	if err != nil {
		t.Errorf("Произошла ошибка: %v\n", err)
	}

	if string(expectedJson) != result {
		t.Errorf("Получен неверный результат: \n%s\n", result)
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
		t.Fatalf("Произошла ошибка: %v", err)
	}
	expectedJson, err := io.ReadAll(expectedFile)
	if err != nil {
		t.Fatalf("Произошла ошибка: %v", err)
	}

	// Act
	result, err := getImageManifests(input)

	// Assert
	if err != nil {
		t.Errorf("Произошла ошибка: %v\n", err)
	}

	if string(expectedJson) != result {
		t.Errorf("Получен неверный результат: \n%s\n", result)
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
	result, err := getImageManifests(input)

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
	result, err := getImageManifests(input)

	// Assert
	if result != "" {
		t.Errorf("Получен результат, хотя должна быть пустая строка: %s\n", result)
	}

	if !errors.Is(err, ErrNotFound) {
		t.Errorf("Должна быть получена ошибка Not Found, но получена другая: %v\n", err)
	}
}
