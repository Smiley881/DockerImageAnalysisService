package services

import (
	"strings"
	"testing"
)

func TestParseJsonToStruct_Input_WithTag(t *testing.T) {

	// Array
	var result Input
	input := strings.NewReader(`{"repository":"dockerhub.timeweb.cloud","name":"python","tag":"slim"}`)
	expectedResult := Input{
		Repository: "dockerhub.timeweb.cloud",
		Name:       "python",
		Tag:        "slim",
	}

	// Act
	err := parseJsonToStructInput(input, &result)

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

func TestParseJsonToStruct_Manifests(t *testing.T) {

	// Array
	var result Manifests
	input := strings.NewReader(`{"manifests":[{"digest":"sha256:0cea405c3ace86f4480c2986d942fc8258cae70c5ffb1fd70143cb5ba54a208c","platform":{"architecture":"amd64","os":"linux"}},{"digest":"sha256:f791f0f4857173e4d3d1590d9de7ab2507b4803bfcf0995a39111d5ab6633d06","platform":{"architecture":"unknown","os":"unknown"}}]}`)
	expectedResult := Manifests{
		Manifests: []Manifest{
			{
				Digest: "sha256:0cea405c3ace86f4480c2986d942fc8258cae70c5ffb1fd70143cb5ba54a208c",
				Platform: Platform{
					Architecture: "amd64",
					Os:           "linux",
				},
			},
			{
				Digest: "sha256:f791f0f4857173e4d3d1590d9de7ab2507b4803bfcf0995a39111d5ab6633d06",
				Platform: Platform{
					Architecture: "unknown",
					Os:           "unknown",
				},
			},
		},
	}

	// Act
	err := parseJsonToStructManifests(input, &result)

	// Assert
	if err != nil {
		t.Errorf("Произошла ошибка: %v\n", err)
	}

	if len(result.Manifests) != 2 {
		t.Fatal("Получен неверный результат")
	}

	checkFirst := result.Manifests[0].Digest != expectedResult.Manifests[0].Digest || result.Manifests[0].Platform.Os != expectedResult.Manifests[0].Platform.Os || result.Manifests[0].Platform.Architecture != expectedResult.Manifests[0].Platform.Architecture
	checkSecond := result.Manifests[1].Digest != expectedResult.Manifests[1].Digest || result.Manifests[1].Platform.Os != expectedResult.Manifests[1].Platform.Os || result.Manifests[1].Platform.Architecture != expectedResult.Manifests[1].Platform.Architecture

	if checkFirst || checkSecond {
		t.Error("Получен неверный результат")
	}
}

func TestParseJsonToStruct_Blob(t *testing.T) {

	// Array
	var result Blob
	input := strings.NewReader(`{"layers":[{"size":28204865},{"size":3511511},{"size":12582814},{"size":249}]}`)
	expectedResult := Blob{
		Layers: []Layer{{Size: 28204865}, {Size: 3511511}, {Size: 12582814}, {Size: 249}},
	}

	// Act
	err := parseJsonToStructBlob(input, &result)

	// Assert
	if err != nil {
		t.Errorf("Произошла ошибка: %v\n", err)
	}

	if len(result.Layers) != 4 {
		t.Fatal("Получен неверный результат")
	}

	if result.Layers[0].Size != expectedResult.Layers[0].Size || result.Layers[1].Size != expectedResult.Layers[1].Size || result.Layers[2].Size != expectedResult.Layers[2].Size || result.Layers[3].Size != expectedResult.Layers[3].Size {
		t.Fatal("Получен неверный результат")
	}
}
