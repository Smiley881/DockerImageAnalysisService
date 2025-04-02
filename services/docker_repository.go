package services

import (
	"errors"
	"io"
	"net/http"
)

var ErrNotFound = errors.New("Образ не найден")

/* Получение списка сборок */
func getListManifests(input Input) (string, error) {
	baseUrl := "https://" + input.Repository + "/v2/" + input.Name + "/manifests/" + input.Tag

	req, err := http.NewRequest("GET", baseUrl, nil)
	if err != nil {
		return "", err
	}
	req.Header.Add("Accept", "application/vnd.oci.image.manifest.v1+json")
	req.Header.Add("Accept", "application/vnd.oci.image.index.v1+json")

	resp, err := http.DefaultClient.Do(req)
	if resp.StatusCode == 404 {
		return "", ErrNotFound
	}
	defer resp.Body.Close()

	result, errRead := io.ReadAll(resp.Body)
	if errRead != nil {
		return "", errRead
	}
	return string(result), err
}

/* Получение информации об одном из манифестов */
func getBlob(input *Input, digest string) (string, error) {
	baseUrl := "https://" + input.Repository + "/v2/" + input.Name + "/blobs/" + digest

	req, err := http.NewRequest("GET", baseUrl, nil)
	if err != nil {
		return "", err
	}

	resp, err := http.DefaultClient.Do(req)
	if resp.StatusCode == 404 {
		return "", ErrNotFound
	}
	defer resp.Body.Close()

	result, errRead := io.ReadAll(resp.Body)
	if errRead != nil {
		return "", errRead
	}
	return string(result), err
}
