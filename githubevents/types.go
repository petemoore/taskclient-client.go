// This source code file is AUTO-GENERATED by github.com/taskcluster/jsonschema2go

package githubevents

import (
	"encoding/json"
)

type (
	// Message reporting that a GitHub pull request has occurred
	GitHubPullRequestMessage struct {

		// The GitHub `action` which triggered an event.
		//
		// Possible values:
		//   * "assigned"
		//   * "unassigned"
		//   * "labeled"
		//   * "unlabeled"
		//   * "opened"
		//   * "closed"
		//   * "reopened"
		//   * "synchronize"
		Action json.RawMessage `json:"action"`

		// Metadata describing the pull request.
		Details json.RawMessage `json:"details"`

		// The GitHub `organization` which had an event.
		//
		// Syntax:     ^([a-zA-Z0-9-_%]*)$
		// Min length: 1
		// Max length: 100
		Organization string `json:"organization"`

		// The GitHub `repository` which had an event.
		//
		// Syntax:     ^([a-zA-Z0-9-_%]*)$
		// Min length: 1
		// Max length: 100
		Repository string `json:"repository"`

		// Message version
		//
		// Possible values:
		//   * 1
		Version json.RawMessage `json:"version"`
	}

	// Message reporting that a GitHub push has occurred
	GitHubPushMessage struct {

		// Metadata describing the push.
		Details json.RawMessage `json:"details"`

		// The GitHub `organization` which had an event.
		//
		// Syntax:     ^([a-zA-Z0-9-_%]*)$
		// Min length: 1
		// Max length: 100
		Organization string `json:"organization"`

		// The GitHub `repository` which had an event.
		//
		// Syntax:     ^([a-zA-Z0-9-_%]*)$
		// Min length: 1
		// Max length: 100
		Repository string `json:"repository"`

		// Message version
		//
		// Possible values:
		//   * 1
		Version json.RawMessage `json:"version"`
	}
)