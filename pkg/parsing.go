package services

import (
	"encoding/json"
	"errors"
	"io"
)

var (
	ErrInvalidfFormatInput = errors.New(`Поля "repository" и "name" являются обязательными`)
	ErrInvalidFormatJson   = errors.New(`Неверный формат json`)
)

type Input struct {
	Repository string `json:"repository"`
	Name       string `json:"name"`
	Tag        string `json:"tag"`
}

/* Превращает строку JSON в структуру Input */
func parseJsonToStructInput(inputJson io.Reader, inputStruct *Input) error {

	decoder := json.NewDecoder(inputJson)
	err := decoder.Decode(&inputStruct)
	if err != nil {
		return ErrInvalidFormatJson
	}

	if inputStruct.Name == "" || inputStruct.Repository == "" {
		return ErrInvalidfFormatInput
	}

	if inputStruct.Tag == "" {
		inputStruct.Tag = "latest"
	}

	return nil
}

// ============================================================ //

type Platform struct {
	Architecture string `json:"architecture"`
	Os           string `json:"os"`
}

type Manifest struct {
	Digest   string   `json:"digest"`
	Platform Platform `json:"platform"`
}

type Manifests struct {
	Manifests []Manifest `json:"manifests"`
}

/* Превращает строку JSON в структуру Manifests */
func parseJsonToStructManifests(inputJson io.Reader, manifests *Manifests) error {

	decoder := json.NewDecoder(inputJson)
	err := decoder.Decode(&manifests)
	if err != nil {
		return err
	}
	return nil
}

// ============================================================ //

type Layer struct {
	Size int `json:"size"`
}

type Blob struct {
	Layers []Layer `json:"layers"`
}

/* Превращает строку JSON в структуру Blob */
func parseJsonToStructBlob(inputJson io.Reader, blob *Blob) error {

	decoder := json.NewDecoder(inputJson)
	err := decoder.Decode(&blob)
	if err != nil {
		return err
	}
	return nil
}
