// The following code is AUTO-GENERATED. Please DO NOT edit.
// To update this generated code, run the following command:
// in the /codegenerator/model subdirectory of this project,
// making sure that `${GOPATH}/bin` is in your `PATH`:
//
// go install && go generate
//
// This package was generated from the schema defined at
// http://references.taskcluster.net/secrets/v1/api.json

// The secrets service provides a simple key/value store for small bits of secret
// data.  Access is limited by scopes, so values can be considered secret from
// those who do not have the relevant scopes.
//
// Secrets also have an expiration date, and once a secret has expired it can no
// longer be read.  This is useful for short-term secrets such as a temporary
// service credential or a one-time signing key.
//
// See: https://docs.taskcluster.net/reference/core/secrets/api-docs
//
// How to use this package
//
// First create a Secrets object:
//
//  secrets := tcsecrets.New(nil)
//
// and then call one or more of secrets's methods, e.g.:
//
//  err := secrets.Set(.....)
//
// handling any errors...
//
//  if err != nil {
//  	// handle error...
//  }
//
// Taskcluster Schema
//
// The source code of this go package was auto-generated from the API definition at
// http://references.taskcluster.net/secrets/v1/api.json together with the input and output schemas it references, downloaded on
// Tue, 10 Apr 2018 at 17:21:00 UTC. The code was generated
// by https://github.com/taskcluster/taskcluster-client-go/blob/master/build.sh.
package tcsecrets

import (
	"net/url"
	"time"

	tcclient "github.com/taskcluster/taskcluster-client-go"
)

const (
	DefaultBaseURL = "https://secrets.taskcluster.net/v1"
)

type Secrets tcclient.Client

// New returns a Secrets client, configured to run against production. Pass in
// nil to create a client without authentication. The
// returned client is mutable, so returned settings can be altered.
//
//  secrets := tcsecrets.New(nil)                              // client without authentication
//  secrets.BaseURL = "http://localhost:1234/api/Secrets/v1"   // alternative API endpoint (production by default)
//  err := secrets.Set(.....)                                  // for example, call the Set(.....) API endpoint (described further down)...
//  if err != nil {
//  	// handle errors...
//  }
func New(credentials *tcclient.Credentials) *Secrets {
	return &Secrets{
		Credentials:  credentials,
		BaseURL:      DefaultBaseURL,
		Authenticate: credentials != nil,
	}
}

// NewFromEnv returns a Secrets client with credentials taken from the environment variables:
//
//  TASKCLUSTER_CLIENT_ID
//  TASKCLUSTER_ACCESS_TOKEN
//  TASKCLUSTER_CERTIFICATE
//
// If environment variables TASKCLUSTER_CLIENT_ID is empty string or undefined
// authentication will be disabled.
func NewFromEnv(credentials *tcclient.Credentials) *Secrets {
	c := tcclient.CredentialsFromEnvVars()
	return &Secrets{
		Credentials:  c,
		BaseURL:      DefaultBaseURL,
		Authenticate: c.ClientID != "",
	}
}

// Set the secret associated with some key.  If the secret already exists, it is
// updated instead.
//
// Required scopes:
//   secrets:set:<name>
//
// See https://docs.taskcluster.net/reference/core/secrets/api-docs#set
func (secrets *Secrets) Set(name string, payload *Secret) error {
	cd := tcclient.Client(*secrets)
	_, _, err := (&cd).APICall(payload, "PUT", "/secret/"+url.QueryEscape(name), nil, nil)
	return err
}

// Delete the secret associated with some key.
//
// Required scopes:
//   secrets:set:<name>
//
// See https://docs.taskcluster.net/reference/core/secrets/api-docs#remove
func (secrets *Secrets) Remove(name string) error {
	cd := tcclient.Client(*secrets)
	_, _, err := (&cd).APICall(nil, "DELETE", "/secret/"+url.QueryEscape(name), nil, nil)
	return err
}

// Read the secret associated with some key.  If the secret has recently
// expired, the response code 410 is returned.  If the caller lacks the
// scope necessary to get the secret, the call will fail with a 403 code
// regardless of whether the secret exists.
//
// Required scopes:
//   secrets:get:<name>
//
// See https://docs.taskcluster.net/reference/core/secrets/api-docs#get
func (secrets *Secrets) Get(name string) (*Secret, error) {
	cd := tcclient.Client(*secrets)
	responseObject, _, err := (&cd).APICall(nil, "GET", "/secret/"+url.QueryEscape(name), new(Secret), nil)
	return responseObject.(*Secret), err
}

// Returns a signed URL for Get, valid for the specified duration.
//
// Required scopes:
//   secrets:get:<name>
//
// See Get for more details.
func (secrets *Secrets) Get_SignedURL(name string, duration time.Duration) (*url.URL, error) {
	cd := tcclient.Client(*secrets)
	return (&cd).SignedURL("/secret/"+url.QueryEscape(name), nil, duration)
}

// List the names of all secrets.
//
// By default this end-point will try to return up to 1000 secret names in one
// request. But it **may return less**, even if more tasks are available.
// It may also return a `continuationToken` even though there are no more
// results. However, you can only be sure to have seen all results if you
// keep calling `listTaskGroup` with the last `continuationToken` until you
// get a result without a `continuationToken`.
//
// If you are not interested in listing all the members at once, you may
// use the query-string option `limit` to return fewer.
//
// See https://docs.taskcluster.net/reference/core/secrets/api-docs#list
func (secrets *Secrets) List(continuationToken, limit string) (*SecretsList, error) {
	v := url.Values{}
	if continuationToken != "" {
		v.Add("continuationToken", continuationToken)
	}
	if limit != "" {
		v.Add("limit", limit)
	}
	cd := tcclient.Client(*secrets)
	responseObject, _, err := (&cd).APICall(nil, "GET", "/secrets", new(SecretsList), v)
	return responseObject.(*SecretsList), err
}

// Respond without doing anything.
// This endpoint is used to check that the service is up.
//
// See https://docs.taskcluster.net/reference/core/secrets/api-docs#ping
func (secrets *Secrets) Ping() error {
	cd := tcclient.Client(*secrets)
	_, _, err := (&cd).APICall(nil, "GET", "/ping", nil, nil)
	return err
}