package services

import (
	"strings"
	"testing"
)

func TestParseJsonInputToStruct_WithTag(t *testing.T) {

	// Array
	var result Input
	input := strings.NewReader(`{"repository":"dockerhub.timeweb.cloud","name":"python","tag":"slim"}`)
	expectedResult := Input{
		Repository: "dockerhub.timeweb.cloud",
		Name:       "python",
		Tag:        "slim",
	}

	// Act
	err := parseJsonToStruct_Input(input, &result)

	// Assert
	if err != nil {
		t.Errorf("Произошла ошибка: %v\n", err)
	}

	if result.Name != expectedResult.Name || result.Repository != expectedResult.Repository || result.Tag != expectedResult.Tag {
		t.Errorf(
			"Получен неверный результат:\nActual:\n- Repository: %s\n- Name: %s\n- Tag: %s\nExpected:\n- Repository: %s\n- Name: %s\n- Tag: %s",
			result.Repository, result.Name, result.Tag, expectedResult.Repository, expectedResult.Name, expectedResult.Tag,
		)
	}
}

func TestParseJsonInputToStruct_WithoutTag(t *testing.T) {

	// Array
	var result Input
	input := strings.NewReader(`{"repository":"dockerhub.timeweb.cloud","name":"python"}`)
	expectedResult := Input{
		Repository: "dockerhub.timeweb.cloud",
		Name:       "python",
		Tag:        "latest",
	}

	// Act
	err := parseJsonToStruct_Input(input, &result)

	// Assert
	if err != nil {
		t.Errorf("Произошла ошибка: %v\n", err)
	}

	if result.Name != expectedResult.Name || result.Repository != expectedResult.Repository || result.Tag != expectedResult.Tag {
		t.Errorf(
			"Получен неверный результат:\nActual:\n- Repository: %s\n- Name: %s\n- Tag: %s\nExpected:\n- Repository: %s\n- Name: %s\n- Tag: %s",
			result.Repository, result.Name, result.Tag, expectedResult.Repository, expectedResult.Name, expectedResult.Tag,
		)
	}
}
