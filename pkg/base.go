package services

import (
	"io"
	"strings"
)

/* Возвращение сборки на linux/amd64 */
func getManifestLinuxAmd64(manifests *Manifests) (string, error) {
	var digest string

	for i := range manifests.Manifests {
		if manifests.Manifests[i].Platform.Architecture == "amd64" && manifests.Manifests[i].Platform.Os == "linux" {
			digest = manifests.Manifests[i].Digest
			break
		}
	}

	if digest == "" {
		return "", ErrNotFound
	}
	return digest, nil
}

/* Возвращение списка сборок в виде структуры */
func GetImageManifests(input io.ReadCloser, inputStruct Input) (Manifests, error) {

	var manifests Manifests

	manifestList, err := getListManifests(inputStruct)
	if err != nil {
		return manifests, err
	}

	err = parseJsonToStruct_Manifests(strings.NewReader(manifestList), &manifests)
	if err != nil {
		return manifests, err
	}

	return manifests, nil
}

/* Возвращение определенной сборки в виде структуры */
func GetImageBlob(digest string, inputStruct Input) (Blob, error) {

	var blob Blob
	blobString, err := getBlob(&inputStruct, digest)
	if err != nil {
		return blob, err
	}

	err = parseJsonToStruct_Blob(strings.NewReader(blobString), &blob)
	if err != nil {
		return blob, err
	}

	return blob, nil
}

type BaseResult struct {
	LayersCount int `json:"layers_count"`
	TotalSize   int `json:"total_size"`
}

/* Возвращает сводную информацию об образе (количество слоев и общий вес)  */
func ImageDownloadSize(input io.ReadCloser) (BaseResult, error) {

	var result BaseResult
	var inputStruct Input
	err := parseJsonToStruct_Input(input, &inputStruct)
	if err != nil {
		return result, err
	}

	manifests, err := GetImageManifests(input, inputStruct)
	if err != nil {
		return result, err
	}

	digest, err := getManifestLinuxAmd64(&manifests)
	if err != nil {
		return result, err
	}

	blob, err := GetImageBlob(digest, inputStruct)
	if err != nil {
		return result, err
	}

	result.LayersCount = len(blob.Layers)

	for i := range result.LayersCount {
		result.TotalSize += blob.Layers[i].Size
	}

	return result, nil
}
