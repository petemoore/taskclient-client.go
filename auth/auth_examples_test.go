package auth_test

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/taskcluster/taskcluster-client-go/auth"
	"github.com/taskcluster/taskcluster-client-go/tcclient"
)

func Example_scopes() {

	// Note: the API call we will make doesn't need credentials as it supplies public information.
	// However, for the purpose of demonstrating the general case, this is how you can provide
	// credentials for API calls that require them.
	myAuth := auth.New(
		tcclient.NewTemporaryCredentials(
			os.Getenv("TASKCLUSTER_CLIENT_ID"),
			os.Getenv("TASKCLUSTER_ACCESS_TOKEN"),
			os.Getenv("TASKCLUSTER_CERTIFICATE"),
			nil,
		),
	)

	// Look up client details for client id "project/taskcluster/tc-client-go/tests"...
	resp, _, err := myAuth.Client("project/taskcluster/tc-client-go/tests")

	// Handle any errors...
	if err != nil {
		log.Printf("Error occurred: %s", err)
		return
	}

	// Report results...
	fmt.Printf("Client ID:  %v\n", resp.ClientID)
	fmt.Printf("Expires:    %v\n", resp.Expires)
	// Could also print expanded scopes, for example:
	//   fmt.Printf("Expanded Scopes:  %v\n", resp.ExpandedScopes)

	// Output:
	// Client ID:  project/taskcluster/tc-client-go/tests
	// Expires:    3017-02-01T05:00:00.000Z
}

func Example_updateClient() {

	// In this example we will connect to a local auth server running on
	// localhost with authentication disabled. This would also work for
	// connecting to a local taskcluster-proxy instance.
	myAuth := auth.New(
		nil,
	)

	// Disable authentication and set target url to localhost url...
	myAuth.Authenticate = false
	myAuth.BaseURL = "http://localhost:60024/v1"

	// Update client id "b2g-power-tests" with new description and expiry...
	client, cs, err := myAuth.UpdateClient(
		"b2g-power-tests",
		&auth.CreateClientRequest{
			Description: "Grant access to download artifacts for `flame-kk-eng`",
			Expires:     tcclient.Time(time.Now().AddDate(1, 0, 0)),
		},
	)

	// Handle any errors...
	if err != nil {
		log.Printf("Error occurred: %s", err)
		return
	}

	// Report results...
	fmt.Printf("Client Id:        %v\n", client.ClientID)
	fmt.Printf("Created:          %v\n", client.Created)
	fmt.Printf("Description:      %v\n", client.Description)
	fmt.Printf("Expanded Scopes:  %v\n", client.ExpandedScopes)
	fmt.Printf("Expires:          %v\n", client.Expires)
	fmt.Printf("Last Date Used:   %v\n", client.LastDateUsed)
	fmt.Printf("Last Modified:    %v\n", client.LastModified)
	fmt.Printf("Last Rotated:     %v\n", client.LastRotated)

	// if we want, we can also show the raw json that was returned...
	fmt.Println(cs.HTTPResponseBody)
}
