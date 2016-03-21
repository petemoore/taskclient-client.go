package tcclient_test

import (
	"fmt"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/taskcluster/taskcluster-client-go/auth"
	"github.com/taskcluster/taskcluster-client-go/tcclient"
)

func ExampleCredentials_CreateTemporaryCredentials() {
	permaCreds := tcclient.NewPermanentCredentials(
		os.Getenv("TASKCLUSTER_CLIENT_ID"),
		os.Getenv("TASKCLUSTER_ACCESS_TOKEN"),
		nil,
	)
	tempCreds, err := permaCreds.CreateTemporaryCredentials(24*time.Hour, "dummy:scope:1", "dummy:scope:2")
	if err != nil {
		// handle error
	}
	fmt.Printf("Temporary creds:\n%q\n", tempCreds)
}

func Test_CreateTemporaryCredentials_WellFormed(t *testing.T) {
	// fake credentials
	permaCreds := tcclient.NewPermanentCredentials(
		"permacred",
		"eHMnHH7PTSqplJSC_qAJ2QKGt8egfvRaqxczIRgOScaw",
		nil,
	)

	tempCreds, err := permaCreds.CreateTemporaryCredentials(24*time.Hour, "scope1")
	if err != nil {
		t.Error(err)
	}

	if tempCreds.AuthorizedScopes != nil {
		t.Errorf("temp creds have AuthorizedScopes!?")
	}

	if tempCreds.ClientID != permaCreds.ClientID {
		t.Errorf("%s != %s", tempCreds.ClientID, permaCreds.ClientID)
	}

	// Certificate and AccessToken are nondeterministic; we rely on other tests
	// to verify them
}

// This clientId/accessToken pair is recognized as valid by the testAutheticate endpoint
var testCreds = tcclient.NewPermanentCredentials(
	"tester",
	"no-secret",
	nil,
)

func checkAuthenticate(t *testing.T, response *auth.TestAuthenticateResponse, err error, expectedClientID string, expectedScopes []string) {

	if err != nil {
		t.Error(err)
		return
	}

	if response.ClientID != expectedClientID {
		t.Errorf("got unexpected clientId %s", response.ClientID)
	}

	if !reflect.DeepEqual(response.Scopes, expectedScopes) {
		t.Errorf("got unexpected scopes %#v", response.Scopes)
	}
}

func Test_PermaCred(t *testing.T) {
	client := auth.New(testCreds)
	response, _, err := client.TestAuthenticate(&auth.TestAuthenticateRequest{
		ClientScopes:   []string{"scope:*"},
		RequiredScopes: []string{"scope:this"},
	})
	checkAuthenticate(t, response, err,
		"tester", []string{"scope:*"})
}

func Test_TempCred(t *testing.T) {
	tempCreds, err := testCreds.CreateTemporaryCredentials(1*time.Hour, "scope:1", "scope:2")
	if err != nil {
		t.Error(err)
		return
	}
	client := auth.New(tempCreds)
	response, _, err := client.TestAuthenticate(&auth.TestAuthenticateRequest{
		ClientScopes:   []string{"scope:*"},
		RequiredScopes: []string{"scope:1"},
	})
	checkAuthenticate(t, response, err,
		"tester", []string{"scope:1", "scope:2"})
}

func Test_NamedTempCred(t *testing.T) {
	tempCreds, err := testCreds.CreateNamedTemporaryCredentials("jimmy", 1*time.Hour,
		"scope:1", "scope:2")
	if err != nil {
		t.Error(err)
		return
	}
	client := auth.New(tempCreds)
	response, _, err := client.TestAuthenticate(&auth.TestAuthenticateRequest{
		ClientScopes:   []string{"scope:*", "auth:create-client:jimmy"},
		RequiredScopes: []string{"scope:1"},
	})
	checkAuthenticate(t, response, err,
		"jimmy", []string{"scope:1", "scope:2"})
}

func Test_TempCred_NoClientId(t *testing.T) {
	baseCreds := tcclient.NewPermanentCredentials("", "no-secret", nil)
	_, err := baseCreds.CreateTemporaryCredentials(1*time.Hour, "s")
	if err == nil {
		t.Errorf("expected error")
	}
}

func Test_TempCred_NoAccessToken(t *testing.T) {
	baseCreds := tcclient.NewPermanentCredentials("tester", "", nil)
	_, err := baseCreds.CreateTemporaryCredentials(1*time.Hour, "s")
	if err == nil {
		t.Errorf("expected error")
	}
}

func Test_TempCred_TooLong(t *testing.T) {
	_, err := testCreds.CreateTemporaryCredentials(32*24*time.Hour, "s")
	if err == nil {
		t.Errorf("expected error")
	}
}

func Test_AuthorizedScopes(t *testing.T) {
	authCreds := *testCreds
	authCreds.AuthorizedScopes = []string{"scope:1", "scope:3"}
	client := auth.New(&authCreds)
	response, _, err := client.TestAuthenticate(&auth.TestAuthenticateRequest{
		ClientScopes:   []string{"scope:*"},
		RequiredScopes: []string{"scope:1"},
	})
	checkAuthenticate(t, response, err,
		"tester", []string{"scope:1", "scope:3"})
}

func Test_TempCredWithAuthorizedScopes(t *testing.T) {
	tempCreds, err := testCreds.CreateTemporaryCredentials(1*time.Hour, "scope:1", "scope:2")
	if err != nil {
		t.Error(err)
		return
	}
	tempCreds.AuthorizedScopes = []string{"scope:1"}
	client := auth.New(tempCreds)
	response, _, err := client.TestAuthenticate(&auth.TestAuthenticateRequest{
		ClientScopes:   []string{"scope:*"},
		RequiredScopes: []string{"scope:1"},
	})
	checkAuthenticate(t, response, err,
		"tester", []string{"scope:1"})
}

func Test_NamedTempCredWithAuthorizedScopes(t *testing.T) {
	tempCreds, err := testCreds.CreateNamedTemporaryCredentials("julie", 1*time.Hour,
		"scope:1", "scope:2")
	if err != nil {
		t.Error(err)
		return
	}
	tempCreds.AuthorizedScopes = []string{"scope:1"} // note: no create-client
	client := auth.New(tempCreds)
	response, _, err := client.TestAuthenticate(&auth.TestAuthenticateRequest{
		ClientScopes:   []string{"scope:*", "auth:create-client:j*"},
		RequiredScopes: []string{"scope:1"},
	})
	checkAuthenticate(t, response, err,
		"julie", []string{"scope:1"})
}
