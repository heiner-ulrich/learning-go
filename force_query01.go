
// meldet mich in der instance an und gibt mir das access token
// Voraussetzung: connected App konfigurieren


/*

Create a connected app: In Salesforce, go to Setup > Create > Apps and click the New button under the Connected Apps section. Fill in the required details and ensure that the Enable OAuth Settings checkbox is checked. Make a note of the Consumer Key and Consumer Secret values that are generated for your app, as you will need them later.

Set up a permission set: In Salesforce, go to Setup > Administer > Manage Users > Permission Sets. Create a new permission set and assign the API Enabled permission to it.

Assign the permission set to your user: In Salesforce, go to Setup > Administer > Manage Users > Users. Select the user that you want to use to log in to the API, scroll down to the Permission Set Assignments section, and add the permission set that you created in step 2 to the user.

Verify that the user has API access: In Salesforce, go to Setup > Administer > Manage Users > Users. Select the user that you want to use to log in to the API and ensure that the API Enabled checkbox is checked.


*/


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



func main() {
	// Set up the Salesforce login endpoint URL
	loginEndpoint := "https://test.salesforce.com/services/oauth2/token"

	// Set up the HTTP client
	client := &http.Client{}

	// Set up the Salesforce login request data
	// https://b2plus-wtw--hu.sandbox.my.salesforce.com //
	data := url.Values{}
	data.Set("grant_type", "password")
	data.Set("client_id", "3MVG9mVMtbWMH6lsNCBsxz8vJ81zcvJk84NERgyxERJ1xBiQIVjIb10FpzTfM9t5rKvIMXAXY5kTBTVEkVkHi")
	data.Set("client_secret", "006A7FBB9184EB15A7D69F509238B3BE98637247441DD5B13DBB68278070C6DB")
	data.Set("username", "hu+117@ulrich.consulting")
	data.Set("password", "jVc9V6.:Y8FVBt+w")

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


}



