package services

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

var (
	ErrInvalidfFormat = errors.New(`InvalidfFormat: Поля "repository" и "name" являются обязательными`)
	ErrNotFound       = errors.New("Not found: Образ не найден")
)

type Input struct {
	Repository string `json:"repository"`
	Name       string `json:"name"`
	Tag        string `json:"tag"`
}

func getImageMainfests(input Input) (string, error) {
	baseUrl := "https://" + input.Repository + "/v2/" + input.Name + "manifests" + input.Tag

	req, err := http.NewRequest("GET", baseUrl, nil)
	if err != nil {
		return "", err
	}
	req.Header.Add("Accept", "application/vnd.oci.image.manifest.v1+json")
	req.Header.Add("Accept", "application/vnd.oci.image.index.v1+json")

	resp, err := http.DefaultClient.Do(req)
	if resp.StatusCode == 404 {
		return "", ErrNotFound
	} else {
		result, errRead := io.ReadAll(req.Body)
		if errRead != nil {
			return "", errRead
		}
		return string(result), err
	}

}

func ImageDownloadSize(input io.ReadCloser) (string, error) {

	var inputStruct Input
	decoder := json.NewDecoder(input)
	err := decoder.Decode(&inputStruct)
	if err != nil {
		return "", err
	}

	if inputStruct.Name == "" || inputStruct.Repository == "" {
		return "", ErrInvalidfFormat
	}

	if inputStruct.Tag == "" {
		inputStruct.Tag = "latest"
	}

	inputByte, err := json.Marshal(&inputStruct)
	if err != nil {
		return "", err
	}
	// req := bytes.NewReader(inputByte)

	return string(inputByte), nil
}
