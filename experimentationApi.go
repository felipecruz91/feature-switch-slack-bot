package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

// Experiment struct
type Experiment struct {
	Name    string      `json:"name"`
	Created interface{} `json:"created"`
	Updated interface{} `json:"updated"`
	For     struct {
		All                         bool          `json:"all"`
		Buckets                     []interface{} `json:"buckets"`
		Users                       []interface{} `json:"users"`
		ExcludesPreExistingEntities bool          `json:"excludesPreExistingEntities"`
	} `json:"for"`
	OnForAll bool `json:"onForAll"`
	Links    []struct {
		Rel  string `json:"rel"`
		Href string `json:"href"`
	} `json:"links"`
}

var expAuthKey = os.Getenv("Z_AUTH_KEY_QA")
var ExperimentationAPIURLQA = os.Getenv("EXP_API_URL_QA")

// GetExperiment performs a /GET HTTP request to retrieve information about the experiment.
func GetExperiment(experimentName string) (experiment Experiment, err error) {
	var exp Experiment
	var URL = ExperimentationAPIURLQA + experimentName

	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		log.Print(err)
		return exp, err
	}

	// Add headers
	req.Header.Add("Z-Auth-Key", expAuthKey)
	req.Header.Add("Content-Type", "application/json; charset=UTF-8")

	// Perform request
	client := &http.Client{Timeout: time.Duration(30 * time.Second)}
	resp, err := client.Do(req)
	if err != nil {
		log.Print(err)
		return exp, err
	}

	if resp.StatusCode == http.StatusNotFound {
		err := fmt.Errorf("Experiment '%s' not found", experimentName)
		log.Print(err)
		return exp, err
	}

	// Read response
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Print(err)
		return exp, err
	}

	// Prase Json string into Experiment struct
	if err := json.Unmarshal(body, &exp); err != nil {
		log.Print(err)
		return exp, err
	}

	return exp, nil
}
