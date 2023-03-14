package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

type SalesforceLoginResponse struct {
	AccessToken string `json:"access_token"`
	InstanceURL string `json:"instance_url"`
}

type SalesforceQueryResult struct {
	TotalSize int `json:"totalSize"`
	Done      bool `json:"done"`
	Records   []map[string]interface{} `json:"records"`
}

func main() {
	// Set up the Salesforce login endpoint URL
	loginEndpoint := "https://test.salesforce.com/services/oauth2/token"

	// Set up the HTTP client
	client := &http.Client{}

	// Set up the Salesforce login request data

	data := url.Values{}
	data.Set("grant_type", "password")
	data.Set("client_id", "3MVG9mVMtbWMH6lsNCBsxz8vJ81zcvJk84NERgyxERJ1xBiQIVjIb10FpzTfM9t5rKvIMXAXY5kTBTVEkVkHi")
	data.Set("client_secret", "006A7FBB9184EB15A7D69F509238B3BE98637247441DD5B13DBB68278070C6DB")
	data.Set("username", "hu+117@ulrich.consulting")
	data.Set("password", "jVc9V6.:Y8FVBt+w") // new pw // connected app disabled (Perform ANSI SQL queries on Customer Data Platform data)


	// Send the Salesforce login request
	req, err := http.NewRequest("POST", loginEndpoint, bytes.NewBufferString(data.Encode()))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// Decode the Salesforce login response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var loginResponse SalesforceLoginResponse
	err = json.Unmarshal(body, &loginResponse)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Print out the Salesforce access token and instance URL
	fmt.Printf("Access token: %s\n", loginResponse.AccessToken)
	fmt.Printf("Instance URL: %s\n", loginResponse.InstanceURL)

	// Send the SOQL query request
	queryEndpoint := loginResponse.InstanceURL + "/services/data/v48.0/query?q=SELECT+Id+FROM+Account"
	req, err = http.NewRequest("GET", queryEndpoint, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	req.Header.Set("Authorization", "Bearer "+loginResponse.AccessToken)

	resp, err = client.Do(req)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// Decode the SOQL query response
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var queryResult SalesforceQueryResult
	err = json.Unmarshal(body, &queryResult)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Print out the SOQL query result
	fmt.Printf("Query result: %v\n", queryResult)
}
