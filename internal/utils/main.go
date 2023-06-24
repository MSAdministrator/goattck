package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/msadministrator/goattck/internal/models"

	"github.com/msadministrator/goattck/internal/logger"
)

var slogger = logger.NewLogger(logger.Info, true)

const attackURL = "https://raw.githubusercontent.com/mitre/cti/master/enterprise-attack/enterprise-attack.json"

// Fetch MITRE ATT&CK data
func Fetch() (models.EnterpriseAttck, error) {
	slogger.Info("Fetching MITRE ATT&CK...")

	res, err := FetchURL(attackURL)
	if err != nil {
		slogger.Fatal("error fetching MITRE ATT&CK")
		return models.EnterpriseAttck{}, err
	}
	return res, nil
}

// FetchURL returns HTTP response body
func FetchURL(url string) (models.EnterpriseAttck, error) {
	resp, err := http.Get(attackURL)
	if err != nil {
		fmt.Println("No response from request")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Can not read body")
	}
	target := models.EnterpriseAttck{}
	if err := json.Unmarshal(body, &target); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}
	//fmt.Println(target.SpecVersion)
	return target, nil
}

func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
