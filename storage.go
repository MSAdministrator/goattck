package goattck

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
)

var (
	DefaultFileName = "./enterprise-attack.json"
)

type storage interface {
	ExistsLocally() (bool, error)
	DownloadAndSave(url string) (RawEnterpriseAttck, error)
	Download(url string) (RawEnterpriseAttck, error)
	Save(enterprise RawEnterpriseAttck) error
	Retrieve() (RawEnterpriseAttck, error)
}

var _ (storage) = new(Storage)

type Storage struct {
}

func (s Storage) ExistsLocally() (bool, error) {
	if _, err := os.Stat(DefaultFileName); err == nil {
		return true, nil
	} else {
		return false, err
	}
}

// Downloads and stores json data to local disk
func (s Storage) DownloadAndSave(url string) (RawEnterpriseAttck, error) {
	// If it doesn't exist locally, we download it
	raw, err := s.Download(url)
	if err != nil {
		return raw, err
	}
	err = s.Save(raw)
	if err != nil {
		return RawEnterpriseAttck{}, err
	}
	return raw, nil
}

// Downloads the enterprise-attack json from the provided URL
func (s Storage) Download(url string) (RawEnterpriseAttck, error) {
	resp, err := http.Get(url)
	if err != nil {
		return RawEnterpriseAttck{}, err
	}
	defer resp.Body.Close()
	return s.unmarshall(resp.Body)
}

// Encodes and store the provided data object to the provided path
func (s Storage) Save(enterprise RawEnterpriseAttck) error {
	file, err := os.Create(DefaultFileName)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(enterprise)
}

// Retrieves the RawEnterpriseAttck from local disk using the provided path
func (s Storage) Retrieve() (RawEnterpriseAttck, error) {
	var raw RawEnterpriseAttck

	file, err := os.Open(DefaultFileName)
	if err != nil {
		return raw, err
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return raw, err
	}
	if err := json.Unmarshal(bytes, &raw); err != nil {
		return raw, err
	}
	return raw, nil
}

func (s Storage) unmarshall(data io.Reader) (RawEnterpriseAttck, error) {
	eAttck := RawEnterpriseAttck{}
	byteValue, err := io.ReadAll(data)
	if err != nil {
		return eAttck, err
	}
	bytesData := []byte(byteValue)
	if err != nil {
		return eAttck, err
	}

	json.Unmarshal(bytesData, &eAttck)
	return eAttck, nil
}
