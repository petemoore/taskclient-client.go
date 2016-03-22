// The following code is AUTO-GENERATED. Please DO NOT edit.
// To update this generated code, run the following command:
// in the /codegenerator/model subdirectory of this project,
// making sure that `${GOPATH}/bin` is in your `PATH`:
//
// go install && go generate
//
// This package was generated from the schema defined at
// http://references.taskcluster.net/auth/v1/api.json

// Authentication related API end-points for TaskCluster and related
// services. These API end-points are of interest if you wish to:
//   * Authenticate request signed with TaskCluster credentials,
//   * Manage clients and roles,
//   * Inspect or audit clients and roles,
//   * Gain access to various services guarded by this API.
//
// ### Clients
// The authentication service manages _clients_, at a high-level each client
// consists of a `clientId`, an `accessToken`, scopes, and some metadata.
// The `clientId` and `accessToken` can be used for authentication when
// calling TaskCluster APIs.
//
// The client's scopes control the client's access to TaskCluster resources.
// The scopes are *expanded* by substituting roles, as defined below.
// Every client has an implicit scope named `assume:client-id:<clientId>`,
// allowing additional access to be granted to the client without directly
// editing the client's scopes.
//
// ### Roles
// A _role_ consists of a `roleId`, a set of scopes and a description.
// Each role constitutes a simple _expansion rule_ that says if you have
// the scope: `assume:<roleId>` you get the set of scopes the role has.
// Think of the `assume:<roleId>` as a scope that allows a client to assume
// a role.
//
// As in scopes the `*` kleene star also have special meaning if it is
// located at the end of a `roleId`. If you have a role with the following
// `roleId`: `my-prefix*`, then any client which has a scope staring with
// `assume:my-prefix` will be allowed to assume the role.
//
// As previously mentioned each client gets the scope:
// `assume:client-id:<clientId>`, it trivially follows that you can create a
// role with the `roleId`: `client-id:<clientId>` to assign additional
// scopes to a client. You can also create a role `client-id:user-*`
// if you wish to assign a set of scopes to all clients whose `clientId`
// starts with `user-`.
//
// ### Guarded Services
// The authentication service also has API end-points for delegating access
// to some guarded service such as AWS S3, or Azure Table Storage.
// Generally, we add API end-points to this server when we wish to use
// TaskCluster credentials to grant access to a third-party service used
// by many TaskCluster components.
//
// See: http://docs.taskcluster.net/auth/api-docs
//
// How to use this package
//
// First create an Auth object:
//
//  myAuth := auth.New(&tcclient.Credentials{ClientId: "myClientId", AccessToken: "myAccessToken"})
//
// and then call one or more of myAuth's methods, e.g.:
//
//  data, callSummary, err := myAuth.ListClients(.....)
// handling any errors...
//  if err != nil {
//  	// handle error...
//  }
//
// TaskCluster Schema
//
// The source code of this go package was auto-generated from the API definition at
// http://references.taskcluster.net/auth/v1/api.json together with the input and output schemas it references, downloaded on
// Thu, 17 Mar 2016 at 19:11:00 UTC. The code was generated
// by https://github.com/taskcluster/taskcluster-client-go/blob/master/build.sh.
package auth

import (
	"net/url"
	"time"

	"github.com/taskcluster/taskcluster-client-go/tcclient"
)

type Auth tcclient.ConnectionData

// Returns a pointer to Auth, configured to run against production.  If you
// wish to point at a different API endpoint url, set BaseURL to the preferred
// url. Authentication can be disabled (for example if you wish to use the
// taskcluster-proxy) by setting Authenticate to false (in which case creds is
// ignored).
//
// For example:
//  creds := &tcclient.Credentials{
//  	ClientId:    os.Getenv("TASKCLUSTER_CLIENT_ID"),
//  	AccessToken: os.Getenv("TASKCLUSTER_ACCESS_TOKEN"),
//  	Certificate: os.Getenv("TASKCLUSTER_CERTIFICATE"),
//  }
//  myAuth := auth.New(creds)                              // set credentials
//  myAuth.Authenticate = false                            // disable authentication (creds above are now ignored)
//  myAuth.BaseURL = "http://localhost:1234/api/Auth/v1"   // alternative API endpoint (production by default)
//  data, callSummary, err := myAuth.ListClients(.....)    // for example, call the ListClients(.....) API endpoint (described further down)...
//  if err != nil {
//  	// handle errors...
//  }
func New(credentials *tcclient.Credentials) *Auth {
	myAuth := Auth(tcclient.ConnectionData{
		Credentials:  credentials,
		BaseURL:      "https://auth.taskcluster.net/v1",
		Authenticate: true,
	})
	return &myAuth
}

// Get a list of all clients.  With `prefix`, only clients for which
// it is a prefix of the clientId are returned.
//
// See http://docs.taskcluster.net/auth/api-docs/#listClients
func (myAuth *Auth) ListClients(prefix string) (*ListClientResponse, error) {
	v := url.Values{}
	v.Add("prefix", prefix)
	cd := tcclient.ConnectionData(*myAuth)
	responseObject, _, err := (&cd).APICall(nil, "GET", "/clients/", new(ListClientResponse), v)
	return responseObject.(*ListClientResponse), err
}

// Get information about a single client.
//
// See http://docs.taskcluster.net/auth/api-docs/#client
func (myAuth *Auth) Client(clientId string) (*GetClientResponse, error) {
	cd := tcclient.ConnectionData(*myAuth)
	responseObject, _, err := (&cd).APICall(nil, "GET", "/clients/"+url.QueryEscape(clientId), new(GetClientResponse), nil)
	return responseObject.(*GetClientResponse), err
}

// Create a new client and get the `accessToken` for this client.
// You should store the `accessToken` from this API call as there is no
// other way to retrieve it.
//
// If you loose the `accessToken` you can call `resetAccessToken` to reset
// it, and a new `accessToken` will be returned, but you cannot retrieve the
// current `accessToken`.
//
// If a client with the same `clientId` already exists this operation will
// fail. Use `updateClient` if you wish to update an existing client.
//
// The caller's scopes must satisfy `scopes`.
//
// Required scopes:
//   * auth:create-client:<clientId>
//
// See http://docs.taskcluster.net/auth/api-docs/#createClient
func (myAuth *Auth) CreateClient(clientId string, payload *CreateClientRequest) (*CreateClientResponse, error) {
	cd := tcclient.ConnectionData(*myAuth)
	responseObject, _, err := (&cd).APICall(payload, "PUT", "/clients/"+url.QueryEscape(clientId), new(CreateClientResponse), nil)
	return responseObject.(*CreateClientResponse), err
}

// Reset a clients `accessToken`, this will revoke the existing
// `accessToken`, generate a new `accessToken` and return it from this
// call.
//
// There is no way to retrieve an existing `accessToken`, so if you loose it
// you must reset the accessToken to acquire it again.
//
// Required scopes:
//   * auth:reset-access-token:<clientId>
//
// See http://docs.taskcluster.net/auth/api-docs/#resetAccessToken
func (myAuth *Auth) ResetAccessToken(clientId string) (*CreateClientResponse, error) {
	cd := tcclient.ConnectionData(*myAuth)
	responseObject, _, err := (&cd).APICall(nil, "POST", "/clients/"+url.QueryEscape(clientId)+"/reset", new(CreateClientResponse), nil)
	return responseObject.(*CreateClientResponse), err
}

// Update an exisiting client. The `clientId` and `accessToken` cannot be
// updated, but `scopes` can be modified.  The caller's scopes must
// satisfy all scopes being added to the client in the update operation.
// If no scopes are given in the request, the client's scopes remain
// unchanged
//
// Required scopes:
//   * auth:update-client:<clientId>
//
// See http://docs.taskcluster.net/auth/api-docs/#updateClient
func (myAuth *Auth) UpdateClient(clientId string, payload *CreateClientRequest) (*GetClientResponse, error) {
	cd := tcclient.ConnectionData(*myAuth)
	responseObject, _, err := (&cd).APICall(payload, "POST", "/clients/"+url.QueryEscape(clientId), new(GetClientResponse), nil)
	return responseObject.(*GetClientResponse), err
}

// Enable a client that was disabled with `disableClient`.  If the client
// is already enabled, this does nothing.
//
// This is typically used by identity providers to re-enable clients that
// had been disabled when the corresponding identity's scopes changed.
//
// Required scopes:
//   * auth:enable-client:<clientId>
//
// See http://docs.taskcluster.net/auth/api-docs/#enableClient
func (myAuth *Auth) EnableClient(clientId string) (*GetClientResponse, error) {
	cd := tcclient.ConnectionData(*myAuth)
	responseObject, _, err := (&cd).APICall(nil, "POST", "/clients/"+url.QueryEscape(clientId)+"/enable", new(GetClientResponse), nil)
	return responseObject.(*GetClientResponse), err
}

// Disable a client.  If the client is already disabled, this does nothing.
//
// This is typically used by identity providers to disable clients when the
// corresponding identity's scopes no longer satisfy the client's scopes.
//
// Required scopes:
//   * auth:disable-client:<clientId>
//
// See http://docs.taskcluster.net/auth/api-docs/#disableClient
func (myAuth *Auth) DisableClient(clientId string) (*GetClientResponse, error) {
	cd := tcclient.ConnectionData(*myAuth)
	responseObject, _, err := (&cd).APICall(nil, "POST", "/clients/"+url.QueryEscape(clientId)+"/disable", new(GetClientResponse), nil)
	return responseObject.(*GetClientResponse), err
}

// Delete a client, please note that any roles related to this client must
// be deleted independently.
//
// Required scopes:
//   * auth:delete-client:<clientId>
//
// See http://docs.taskcluster.net/auth/api-docs/#deleteClient
func (myAuth *Auth) DeleteClient(clientId string) error {
	cd := tcclient.ConnectionData(*myAuth)
	_, _, err := (&cd).APICall(nil, "DELETE", "/clients/"+url.QueryEscape(clientId), nil, nil)
	return err
}

// Get a list of all roles, each role object also includes the list of
// scopes it expands to.
//
// See http://docs.taskcluster.net/auth/api-docs/#listRoles
func (myAuth *Auth) ListRoles() (*ListRolesResponse, error) {
	cd := tcclient.ConnectionData(*myAuth)
	responseObject, _, err := (&cd).APICall(nil, "GET", "/roles/", new(ListRolesResponse), nil)
	return responseObject.(*ListRolesResponse), err
}

// Get information about a single role, including the set of scopes that the
// role expands to.
//
// See http://docs.taskcluster.net/auth/api-docs/#role
func (myAuth *Auth) Role(roleId string) (*GetRoleResponse, error) {
	cd := tcclient.ConnectionData(*myAuth)
	responseObject, _, err := (&cd).APICall(nil, "GET", "/roles/"+url.QueryEscape(roleId), new(GetRoleResponse), nil)
	return responseObject.(*GetRoleResponse), err
}

// Create a new role.
//
// The caller's scopes must satisfy the new role's scopes.
//
// If there already exists a role with the same `roleId` this operation
// will fail. Use `updateRole` to modify an existing role.
//
// Required scopes:
//   * auth:create-role:<roleId>
//
// See http://docs.taskcluster.net/auth/api-docs/#createRole
func (myAuth *Auth) CreateRole(roleId string, payload *CreateRoleRequest) (*GetRoleResponse, error) {
	cd := tcclient.ConnectionData(*myAuth)
	responseObject, _, err := (&cd).APICall(payload, "PUT", "/roles/"+url.QueryEscape(roleId), new(GetRoleResponse), nil)
	return responseObject.(*GetRoleResponse), err
}

// Update an existing role.
//
// The caller's scopes must satisfy all of the new scopes being added, but
// need not satisfy all of the client's existing scopes.
//
// Required scopes:
//   * auth:update-role:<roleId>
//
// See http://docs.taskcluster.net/auth/api-docs/#updateRole
func (myAuth *Auth) UpdateRole(roleId string, payload *CreateRoleRequest) (*GetRoleResponse, error) {
	cd := tcclient.ConnectionData(*myAuth)
	responseObject, _, err := (&cd).APICall(payload, "POST", "/roles/"+url.QueryEscape(roleId), new(GetRoleResponse), nil)
	return responseObject.(*GetRoleResponse), err
}

// Delete a role. This operation will succeed regardless of whether or not
// the role exists.
//
// Required scopes:
//   * auth:delete-role:<roleId>
//
// See http://docs.taskcluster.net/auth/api-docs/#deleteRole
func (myAuth *Auth) DeleteRole(roleId string) error {
	cd := tcclient.ConnectionData(*myAuth)
	_, _, err := (&cd).APICall(nil, "DELETE", "/roles/"+url.QueryEscape(roleId), nil, nil)
	return err
}

// Return an expanded copy of the given scopeset, with scopes implied by any
// roles included.
//
// See http://docs.taskcluster.net/auth/api-docs/#expandScopes
func (myAuth *Auth) ExpandScopes(payload *SetOfScopes) (*SetOfScopes, error) {
	cd := tcclient.ConnectionData(*myAuth)
	responseObject, _, err := (&cd).APICall(payload, "GET", "/scopes/expand", new(SetOfScopes), nil)
	return responseObject.(*SetOfScopes), err
}

// Return the expanded scopes available in the request, taking into account all sources
// of scopes and scope restrictions (temporary credentials, assumeScopes, client scopes,
// and roles).
//
// See http://docs.taskcluster.net/auth/api-docs/#currentScopes
func (myAuth *Auth) CurrentScopes() (*SetOfScopes, error) {
	cd := tcclient.ConnectionData(*myAuth)
	responseObject, _, err := (&cd).APICall(nil, "GET", "/scopes/current", new(SetOfScopes), nil)
	return responseObject.(*SetOfScopes), err
}

// Stability: *** EXPERIMENTAL ***
//
// Get temporary AWS credentials for `read-write` or `read-only` access to
// a given `bucket` and `prefix` within that bucket.
// The `level` parameter can be `read-write` or `read-only` and determines
// which type of credentials are returned. Please note that the `level`
// parameter is required in the scope guarding access.
//
// The credentials are set to expire after an hour, but this behavior is
// subject to change. Hence, you should always read the `expires` property
// from the response, if you intend to maintain active credentials in your
// application.
//
// Please note that your `prefix` may not start with slash `/`. Such a prefix
// is allowed on S3, but we forbid it here to discourage bad behavior.
//
// Also note that if your `prefix` doesn't end in a slash `/`, the STS
// credentials may allow access to unexpected keys, as S3 does not treat
// slashes specially.  For example, a prefix of `my-folder` will allow
// access to `my-folder/file.txt` as expected, but also to `my-folder.txt`,
// which may not be intended.
//
// Required scopes:
//   * auth:aws-s3:<level>:<bucket>/<prefix>
//
// See http://docs.taskcluster.net/auth/api-docs/#awsS3Credentials
func (myAuth *Auth) AwsS3Credentials(level, bucket, prefix string) (*AWSS3CredentialsResponse, error) {
	cd := tcclient.ConnectionData(*myAuth)
	responseObject, _, err := (&cd).APICall(nil, "GET", "/aws/s3/"+url.QueryEscape(level)+"/"+url.QueryEscape(bucket)+"/"+url.QueryEscape(prefix), new(AWSS3CredentialsResponse), nil)
	return responseObject.(*AWSS3CredentialsResponse), err
}

// Returns a signed URL for AwsS3Credentials, valid for the specified duration.
//
// Required scopes:
//   * auth:aws-s3:<level>:<bucket>/<prefix>
//
// See AwsS3Credentials for more details.
func (myAuth *Auth) AwsS3Credentials_SignedURL(level, bucket, prefix string, duration time.Duration) (*url.URL, error) {
	cd := tcclient.ConnectionData(*myAuth)
	return (&cd).SignedURL("/aws/s3/"+url.QueryEscape(level)+"/"+url.QueryEscape(bucket)+"/"+url.QueryEscape(prefix), nil, duration)
}

// Get a shared access signature (SAS) string for use with a specific Azure
// Table Storage table.  Note, this will create the table, if it doesn't
// already exist.
//
// Required scopes:
//   * auth:azure-table-access:<account>/<table>
//
// See http://docs.taskcluster.net/auth/api-docs/#azureTableSAS
func (myAuth *Auth) AzureTableSAS(account, table string) (*AzureSharedAccessSignatureResponse, error) {
	cd := tcclient.ConnectionData(*myAuth)
	responseObject, _, err := (&cd).APICall(nil, "GET", "/azure/"+url.QueryEscape(account)+"/table/"+url.QueryEscape(table)+"/read-write", new(AzureSharedAccessSignatureResponse), nil)
	return responseObject.(*AzureSharedAccessSignatureResponse), err
}

// Returns a signed URL for AzureTableSAS, valid for the specified duration.
//
// Required scopes:
//   * auth:azure-table-access:<account>/<table>
//
// See AzureTableSAS for more details.
func (myAuth *Auth) AzureTableSAS_SignedURL(account, table string, duration time.Duration) (*url.URL, error) {
	cd := tcclient.ConnectionData(*myAuth)
	return (&cd).SignedURL("/azure/"+url.QueryEscape(account)+"/table/"+url.QueryEscape(table)+"/read-write", nil, duration)
}

// Validate the request signature given on input and return list of scopes
// that the authenticating client has.
//
// This method is used by other services that wish rely on TaskCluster
// credentials for authentication. This way we can use Hawk without having
// the secret credentials leave this service.
//
// See http://docs.taskcluster.net/auth/api-docs/#authenticateHawk
func (myAuth *Auth) AuthenticateHawk(payload *HawkSignatureAuthenticationRequest) (*HawkSignatureAuthenticationResponse, error) {
	cd := tcclient.ConnectionData(*myAuth)
	responseObject, _, err := (&cd).APICall(payload, "POST", "/authenticate-hawk", new(HawkSignatureAuthenticationResponse), nil)
	return responseObject.(*HawkSignatureAuthenticationResponse), err
}

// Stability: *** EXPERIMENTAL ***
//
// Utility method to test client implementations of TaskCluster
// authentication.
//
// Rather than using real credentials, this endpoint accepts requests with
// clientId `tester` and accessToken `no-secret`. That client's scopes are
// based on `clientScopes` in the request body.
//
// The request is validated, with any certificate, authorizedScopes, etc.
// applied, and the resulting scopes are checked against `requiredScopes`
// from the request body. On success, the response contains the clientId
// and scopes as seen by the API method.
//
// See http://docs.taskcluster.net/auth/api-docs/#testAuthenticate
func (myAuth *Auth) TestAuthenticate(payload *TestAuthenticateRequest) (*TestAuthenticateResponse, error) {
	cd := tcclient.ConnectionData(*myAuth)
	responseObject, _, err := (&cd).APICall(payload, "POST", "/test-authenticate", new(TestAuthenticateResponse), nil)
	return responseObject.(*TestAuthenticateResponse), err
}

// Stability: *** EXPERIMENTAL ***
//
// Utility method similar to `testAuthenticate`, but with the GET method,
// so it can be used with signed URLs (bewits).
//
// Rather than using real credentials, this endpoint accepts requests with
// clientId `tester` and accessToken `no-secret`. That client's scopes are
// `['test:*', 'auth:create-client:test:*']`.  The call fails if the
// `test:authenticate-get` scope is not available.
//
// The request is validated, with any certificate, authorizedScopes, etc.
// applied, and the resulting scopes are checked, just like any API call.
// On success, the response contains the clientId and scopes as seen by
// the API method.
//
// This method may later be extended to allow specification of client and
// required scopes via query arguments.
//
// See http://docs.taskcluster.net/auth/api-docs/#testAuthenticateGet
func (myAuth *Auth) TestAuthenticateGet() (*TestAuthenticateResponse, error) {
	cd := tcclient.ConnectionData(*myAuth)
	responseObject, _, err := (&cd).APICall(nil, "GET", "/test-authenticate-get/", new(TestAuthenticateResponse), nil)
	return responseObject.(*TestAuthenticateResponse), err
}

// Stability: *** EXPERIMENTAL ***
//
// Documented later...
//
// **Warning** this api end-point is **not stable**.
//
// See http://docs.taskcluster.net/auth/api-docs/#ping
func (myAuth *Auth) Ping() error {
	cd := tcclient.ConnectionData(*myAuth)
	_, _, err := (&cd).APICall(nil, "GET", "/ping", nil, nil)
	return err
}
