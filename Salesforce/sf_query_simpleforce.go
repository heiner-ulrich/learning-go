package main

import (
	"github.com/simpleforce/simpleforce"
	"fmt"
)

var (
	sfURL      = "https://ulrichconsulting--ucdev02.sandbox.lightning.force.com/"
	sfUser     = "ulrich.consulting@pboedition.com.ucdev02"
	sfPassword = "T{pSPd=p3W3m[BPC?]Uz4"
	sfToken    = "tqtn8ylBvFj489FtSXO0lDheY" 
)

func main() {
	createClient()
	//Query()
}



func createClient() *simpleforce.Client {
	client := simpleforce.NewClient(sfURL, simpleforce.DefaultClientID, simpleforce.DefaultAPIVersion)
	if client == nil {
		// handle the error
		fmt.Println("Error in createClient()")
		return nil
	}
	fmt.Println(client)
	return client
	

	err := client.LoginPassword(sfUser, sfPassword, sfToken)
	if err != nil {
		// handle the error
		fmt.Println("Error in Login")
		return nil
	}

	// Do some other stuff with the client instance if needed.

	return client
}


func Query() {
	client := simpleforce.NewClient(sfURL, simpleforce.DefaultClientID, simpleforce.DefaultAPIVersion)
	client.LoginPassword(sfUser, sfPassword, sfToken)

	q := "select name from account limit 20"
	result, err := client.Query(q) // Note: for Tooling API, use client.Tooling().Query(q)
	if err != nil {
		// handle the error
		return
	}

	for _, record := range result.Records {
		// access the record as SObjects.
		fmt.Println(record)
	}
}
