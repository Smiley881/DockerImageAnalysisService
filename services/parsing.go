package services

import (
	"encoding/json"
	"errors"
	"io"
)

var ErrInvalidfFormat = errors.New(`invalidfFormat: Поля "repository" и "name" являются обязательными`)

type Input struct {
	Repository string `json:"repository"`
	Name       string `json:"name"`
	Tag        string `json:"tag"`
}

/* Превращает строку JSON в структуру Input */
func parseJsonToStruct_Input(inputJson io.Reader, inputStruct *Input) error {

	decoder := json.NewDecoder(inputJson)
	err := decoder.Decode(&inputStruct)
	if err != nil {
		return err
	}

	if inputStruct.Name == "" || inputStruct.Repository == "" {
		return ErrInvalidfFormat
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
func parseJsonToStruct_Manifests(inputJson io.Reader, manifests *Manifests) error {

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
func parseJsonToStruct_Blob(inputJson io.Reader, blob *Blob) error {

	decoder := json.NewDecoder(inputJson)
	err := decoder.Decode(&blob)
	if err != nil {
		return err
	}
	return nil
}
