package tcclient

import (
	"encoding/base64"
	"encoding/json"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/taskcluster/taskcluster-base-go/jsontest"
)

// TestExtHeaderPermAuthScopes checks that the generated hawk ext http header
// for permanent credentials with authorized scopes listed matches what is
// expected.
func TestExtHeaderPermAuthScopes(t *testing.T) {
	checkExtHeader(
		t,
		NewPermanentCredentials(
			"abc",
			"def",
			[]string{"a", "b", "c"},
		),
		// base64 of `{"authorizedScopes":["a","b","c"]}`
		"eyJhdXRob3JpemVkU2NvcGVzIjpbImEiLCJiIiwiYyJdfQ==",
	)
}

// TestExtHeaderPermNilAuthScopes checks that when permanent credentials are
// provided and the Authorized Scopes are not set (i.e. are nil) that the hawk
// ext header is an empty string.
func TestExtHeaderPermNilAuthScopes(t *testing.T) {
	checkExtHeader(
		t,
		NewPermanentCredentials(
			"abc",
			"def",
			nil,
		),
		"",
	)
}

// TestExtHeaderPermNoAuthScopes checks that when permanent credentials are
// provided and an empty list of authorized scopes is used, that the hawk ext
// http header is explicitly showing an empty list of authorized scopes.
func TestExtHeaderPermNoAuthScopes(t *testing.T) {
	checkExtHeader(
		t,
		NewPermanentCredentials(
			"abc",
			"def",
			[]string{},
		),
		// base64 of `{"authorizedScopes":[]}`
		"eyJhdXRob3JpemVkU2NvcGVzIjpbXX0=",
	)
}

// TestExtHeaderTempAuthScopes checks that the hawk ext header is set to the
// expected value when using temp credentials and an explicit list of
// authorized scopes.
func TestExtHeaderTempAuthScopes(t *testing.T) {
	checkExtHeaderTempCreds(
		t,
		NewPermanentCredentials(
			"abc",
			"def",
			[]string{"a", "b", "c"},
		),
	)
}

// TestExtHeaderTempNilAuthScopes checks that the hawk ext header includes the
// temporary credentials certificate, but no authorized scopes property when
// using temp credentials but not restricting the authorized scopes.
func TestExtHeaderTempNilAuthScopes(t *testing.T) {
	checkExtHeaderTempCreds(
		t,
		NewPermanentCredentials(
			"abc",
			"def",
			nil,
		),
	)
}

// TestExtHeaderTempNoAuthScopes checks that the hawk ext header includes an
// empty list of authorized scopes when an empty list is provided, and that the
// temp credentials certificate is also included.
func TestExtHeaderTempNoAuthScopes(t *testing.T) {
	checkExtHeaderTempCreds(
		t,
		NewPermanentCredentials(
			"abc",
			"def",
			[]string{},
		),
	)
}

type ExtHeaderRawCert struct {
	Certificate      json.RawMessage `json:"certificate"`
	AuthorizedScopes []string        `json:"authorizedScopes"`
}

// checkExtHeaderTempCreds generates temporary credentials from the given
// permanent credentials and then checks what the ext header looks like
// according to getExtHeader function. It base64 decodes the results, and then
// checks that the temporary credentials match the ones given, and then
// evaluates whether authorizedScopes is correct. It checks that if no
// authorized scopes were set, that the authorizedScopes are not set in the
// header; if they are set to anything, including an empty array, that this
// matches what is found in the header.
func checkExtHeaderTempCreds(t *testing.T, permCreds *PermanentCredentials) {
	tempCredentials, err := permCreds.CreateTemporaryCredentials(time.Second*1, "d", "e", "f")
	if err != nil {
		t.Fatalf("Received error when generating temporary credentials: %s", err)
	}
	actualHeader, err := tempCredentials.getExtField()
	if err != nil {
		t.Fatalf("Received error when generating ext header: %s", err)
	}
	decoded, err := base64.StdEncoding.DecodeString(actualHeader)
	if err != nil {
		t.Fatalf("Received error when base64 decoding ext header: %s", err)
	}
	extHeader := new(ExtHeaderRawCert)
	err = json.Unmarshal(decoded, extHeader)
	if err != nil {
		t.Fatalf("Cannot marshal results back into ExtHeader: %s", err)
	}
	if permCreds.AuthorizedScopes == nil {
		if strings.Contains(string(decoded), "authorizedScopes") {
			t.Fatalf("Did not expected authorizedScopes to be in ext header")
		}
	} else {
		if !reflect.DeepEqual(permCreds.AuthorizedScopes, extHeader.AuthorizedScopes) {
			t.Log("Expected AuthorizedScopes in Hawk Ext header to match AuthorizedScopes in credentials, but they didn't.")
			t.Logf("Expected: %q", permCreds.AuthorizedScopes)
			t.Logf("Actual: %q", extHeader.AuthorizedScopes)
			t.Logf("Full ext header: %s", string(decoded))
			t.FailNow()
		}
	}
	jsonCorrect, formattedExpected, formattedActual, err := jsontest.JsonEqual([]byte(tempCredentials.Certificate), extHeader.Certificate)
	if err != nil {
		t.Fatalf("Exception thrown formatting json data!\n%s\n\nStruggled to format either:\n%s\n\nor:\n\n%s", err, tempCredentials.Certificate, string(extHeader.Certificate))
	}

	if !jsonCorrect {
		t.Log("Anticipated json not generated. Expected:")
		t.Logf("%s", formattedExpected)
		t.Log("Actual:")
		t.Logf("%s", formattedActual)
		t.FailNow()
	}
}

// checkExtHeader simply checks if getExtHeader returns the same results as the
// specified expected header.
func checkExtHeader(t *testing.T, creds Credentials, expectedHeader string) {
	actualHeader, err := creds.getExtField()
	if err != nil {
		t.Fatalf("Received error when generating ext header: %s", err)
	}
	if actualHeader != expectedHeader {
		t.Fatalf("Expected header %q but got %q", expectedHeader, actualHeader)
	}
}
