package services

import (
	"encoding/json"
	"errors"
	"io"
)

var (
	ErrInvalidfFormat = errors.New(`invalidfFormat: Поля "repository" и "name" являются обязательными`)
	ErrNotFound       = errors.New("not found: Образ не найден")
)

type Input struct {
	Repository string `json:"repository"`
	Name       string `json:"name"`
	Tag        string `json:"tag"`
}

func parseJsonInputToStruct(inputJson io.Reader) (Input, error) {
	var inputStruct Input

	decoder := json.NewDecoder(inputJson)
	err := decoder.Decode(&inputStruct)
	if err != nil {
		return inputStruct, err
	}

	if inputStruct.Name == "" || inputStruct.Repository == "" {
		return inputStruct, ErrInvalidfFormat
	}

	if inputStruct.Tag == "" {
		inputStruct.Tag = "latest"
	}

	return inputStruct, nil
}

func ImageDownloadSize(input io.ReadCloser) (string, error) {

	inputStruct, err := parseJsonInputToStruct(input)
	if err != nil {
		return "", err
	}

	inputByte, err := json.Marshal(&inputStruct)
	if err != nil {
		return "", err
	}
	// req := bytes.NewReader(inputByte)

	return string(inputByte), nil
}
