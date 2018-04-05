// This source code file is AUTO-GENERATED by github.com/taskcluster/jsonschema2go

package hooks

import (
	"encoding/json"
	"errors"

	tcclient "github.com/taskcluster/taskcluster-client-go"
)

type (
	// Information about an unsuccessful firing of the hook
	//
	// See http://schemas.taskcluster.net/hooks/v1/hook-status.json#/properties/lastFire/oneOf[1]
	FailedFire struct {

		// The error that occurred when firing the task.  This is typically,
		// but not always, an API error message.
		//
		// See http://schemas.taskcluster.net/hooks/v1/hook-status.json#/properties/lastFire/oneOf[1]/properties/error
		Error json.RawMessage `json:"error"`

		// Possible values:
		//   * "error"
		//
		// See http://schemas.taskcluster.net/hooks/v1/hook-status.json#/properties/lastFire/oneOf[1]/properties/result
		Result string `json:"result"`

		// The time the task was created.  This will not necessarily match `task.created`.
		//
		// See http://schemas.taskcluster.net/hooks/v1/hook-status.json#/properties/lastFire/oneOf[1]/properties/time
		Time tcclient.Time `json:"time"`
	}

	// Definition of a hook that can create tasks at defined times.
	//
	// See http://schemas.taskcluster.net/hooks/v1/create-hook-request.json#
	HookCreationRequest struct {

		// Use of this field is deprecated; use `deadline: {$fromNow: ..}` in the task template instead.
		//
		// Default:    "1 day"
		//
		// See http://schemas.taskcluster.net/hooks/v1/create-hook-request.json#/properties/deadline
		Deadline string `json:"deadline,omitempty"`

		// Use of this field is deprecated; use `expires: {$fromNow: ..}` in the task template instead.
		//
		// Default:    "3 months"
		//
		// See http://schemas.taskcluster.net/hooks/v1/create-hook-request.json#/properties/expires
		Expires string `json:"expires,omitempty"`

		// Syntax:     ^([a-zA-Z0-9-_]*)$
		// Min length: 1
		// Max length: 22
		//
		// See http://schemas.taskcluster.net/hooks/v1/create-hook-request.json#/properties/hookGroupId
		HookGroupID string `json:"hookGroupId,omitempty"`

		// Max length: 255
		//
		// See http://schemas.taskcluster.net/hooks/v1/create-hook-request.json#/properties/hookId
		HookID string `json:"hookId,omitempty"`

		// See http://schemas.taskcluster.net/hooks/v1/create-hook-request.json#/properties/metadata
		Metadata struct {

			// Long-form of the hook's purpose and behavior
			//
			// Max length: 32768
			//
			// See http://schemas.taskcluster.net/hooks/v1/create-hook-request.json#/properties/metadata/properties/description
			Description string `json:"description"`

			// Whether to email the owner on an error creating the task.
			//
			// Default:    true
			//
			// See http://schemas.taskcluster.net/hooks/v1/create-hook-request.json#/properties/metadata/properties/emailOnError
			EmailOnError bool `json:"emailOnError,omitempty"`

			// Human readable name of the hook
			//
			// Max length: 255
			//
			// See http://schemas.taskcluster.net/hooks/v1/create-hook-request.json#/properties/metadata/properties/name
			Name string `json:"name"`

			// Email of the person or group responsible for this hook.
			//
			// Max length: 255
			//
			// See http://schemas.taskcluster.net/hooks/v1/create-hook-request.json#/properties/metadata/properties/owner
			Owner string `json:"owner"`
		} `json:"metadata"`

		// Definition of the times at which a hook will result in creation of a task.
		// If several patterns are specified, tasks will be created at any time
		// specified by one or more patterns.
		//
		// Default:    []
		//
		// See http://schemas.taskcluster.net/hooks/v1/create-hook-request.json#/properties/schedule
		Schedule []string `json:"schedule,omitempty"`

		// Template for the task definition.  This is rendered using [JSON-e](https://taskcluster.github.io/json-e/)
		// as described in https://docs.taskcluster.net/reference/core/taskcluster-hooks/docs/firing-hooks to produce
		// a task definition that is submitted to the Queue service.
		//
		// See http://schemas.taskcluster.net/hooks/v1/create-hook-request.json#/properties/task
		Task json.RawMessage `json:"task"`

		// Default:    map["additionalProperties":%!q(bool=false) "type":"object"]
		//
		// See http://schemas.taskcluster.net/hooks/v1/create-hook-request.json#/properties/triggerSchema
		TriggerSchema json.RawMessage `json:"triggerSchema,omitempty"`
	}

	// Definition of a hook that will create tasks when defined events occur.
	//
	// See http://schemas.taskcluster.net/hooks/v1/hook-definition.json#
	HookDefinition struct {

		// Use of this field is deprecated; use `deadline: {$fromNow: ..}` in the task template instead.
		//
		// Default:    "1 day"
		//
		// See http://schemas.taskcluster.net/hooks/v1/hook-definition.json#/properties/deadline
		Deadline string `json:"deadline"`

		// Use of this field is deprecated; use `expires: {$fromNow: ..}` in the task template instead.
		//
		// Default:    "3 months"
		//
		// See http://schemas.taskcluster.net/hooks/v1/hook-definition.json#/properties/expires
		Expires string `json:"expires"`

		// Syntax:     ^([a-zA-Z0-9-_]*)$
		// Min length: 1
		// Max length: 22
		//
		// See http://schemas.taskcluster.net/hooks/v1/hook-definition.json#/properties/hookGroupId
		HookGroupID string `json:"hookGroupId"`

		// Max length: 255
		//
		// See http://schemas.taskcluster.net/hooks/v1/hook-definition.json#/properties/hookId
		HookID string `json:"hookId"`

		// See http://schemas.taskcluster.net/hooks/v1/hook-definition.json#/properties/metadata
		Metadata struct {

			// Long-form of the hook's purpose and behavior
			//
			// Max length: 32768
			//
			// See http://schemas.taskcluster.net/hooks/v1/hook-definition.json#/properties/metadata/properties/description
			Description string `json:"description"`

			// Whether to email the owner on an error creating the task.
			//
			// Default:    true
			//
			// See http://schemas.taskcluster.net/hooks/v1/hook-definition.json#/properties/metadata/properties/emailOnError
			EmailOnError bool `json:"emailOnError,omitempty"`

			// Human readable name of the hook
			//
			// Max length: 255
			//
			// See http://schemas.taskcluster.net/hooks/v1/hook-definition.json#/properties/metadata/properties/name
			Name string `json:"name"`

			// Email of the person or group responsible for this hook.
			//
			// Max length: 255
			//
			// See http://schemas.taskcluster.net/hooks/v1/hook-definition.json#/properties/metadata/properties/owner
			Owner string `json:"owner"`
		} `json:"metadata"`

		// Definition of the times at which a hook will result in creation of a task.
		// If several patterns are specified, tasks will be created at any time
		// specified by one or more patterns.  Note that tasks may not be created
		// at exactly the time specified.
		//                     {$ref: "http://schemas.taskcluster.net/hooks/v1/schedule.json"}
		//
		// See http://schemas.taskcluster.net/hooks/v1/hook-definition.json#/properties/schedule
		Schedule json.RawMessage `json:"schedule"`

		// Template for the task definition.  This is rendered using [JSON-e](https://taskcluster.github.io/json-e/)
		// as described in https://docs.taskcluster.net/reference/core/taskcluster-hooks/docs/firing-hooks to produce
		// a task definition that is submitted to the Queue service.
		//
		// See http://schemas.taskcluster.net/hooks/v1/hook-definition.json#/properties/task
		Task json.RawMessage `json:"task"`

		// See http://schemas.taskcluster.net/hooks/v1/hook-definition.json#/properties/triggerSchema
		TriggerSchema json.RawMessage `json:"triggerSchema"`
	}

	// List of `hookGroupIds`.
	//
	// See http://schemas.taskcluster.net/hooks/v1/list-hook-groups-response.json#
	HookGroups struct {

		// See http://schemas.taskcluster.net/hooks/v1/list-hook-groups-response.json#/properties/groups
		Groups []string `json:"groups"`
	}

	// List of hooks
	//
	// See http://schemas.taskcluster.net/hooks/v1/list-hooks-response.json#
	HookList struct {

		// See http://schemas.taskcluster.net/hooks/v1/list-hooks-response.json#/properties/hooks
		Hooks []HookDefinition `json:"hooks"`
	}

	// A description of when a hook's task will be created, and the next scheduled time
	//
	// See http://schemas.taskcluster.net/hooks/v1/hook-schedule.json#
	HookScheduleResponse struct {

		// The next time this hook's task is scheduled to be created. This property
		// is only present if there is a scheduled next time. Some hooks don't have
		// any schedules.
		//
		// See http://schemas.taskcluster.net/hooks/v1/hook-schedule.json#/properties/nextScheduledDate
		NextScheduledDate tcclient.Time `json:"nextScheduledDate,omitempty"`

		// See http://schemas.taskcluster.net/hooks/v1/hook-schedule.json#/properties/schedule
		Schedule Schedule `json:"schedule"`
	}

	// A snapshot of the current status of a hook.
	//
	// See http://schemas.taskcluster.net/hooks/v1/hook-status.json#
	HookStatusResponse struct {

		// Information about the last time this hook fired.  This property is only present
		// if the hook has fired at least once.
		//
		// See http://schemas.taskcluster.net/hooks/v1/hook-status.json#/properties/lastFire
		LastFire json.RawMessage `json:"lastFire"`

		// The next time this hook's task is scheduled to be created. This property
		// is only present if there is a scheduled next time. Some hooks don't have
		// any schedules.
		//
		// See http://schemas.taskcluster.net/hooks/v1/hook-status.json#/properties/nextScheduledDate
		NextScheduledDate tcclient.Time `json:"nextScheduledDate,omitempty"`
	}

	// Information about no firing of the hook (e.g., a new hook)
	//
	// See http://schemas.taskcluster.net/hooks/v1/hook-status.json#/properties/lastFire/oneOf[2]
	NoFire struct {

		// Possible values:
		//   * "no-fire"
		//
		// See http://schemas.taskcluster.net/hooks/v1/hook-status.json#/properties/lastFire/oneOf[2]/properties/result
		Result string `json:"result"`
	}

	// A list of cron-style definitions to represent a set of moments in (UTC) time.
	// If several patterns are specified, a given moment in time represented by
	// more than one pattern is considered only to be counted once, in other words
	// it is allowed for the cron patterns to overlap; duplicates are redundant.
	//
	// Default:    []
	//
	// See http://schemas.taskcluster.net/hooks/v1/schedule.json#
	Schedule []string

	// Information about a successful firing of the hook
	//
	// See http://schemas.taskcluster.net/hooks/v1/hook-status.json#/properties/lastFire/oneOf[0]
	SuccessfulFire struct {

		// Possible values:
		//   * "success"
		//
		// See http://schemas.taskcluster.net/hooks/v1/hook-status.json#/properties/lastFire/oneOf[0]/properties/result
		Result string `json:"result"`

		// The task created
		//
		// Syntax:     ^[A-Za-z0-9_-]{8}[Q-T][A-Za-z0-9_-][CGKOSWaeimquy26-][A-Za-z0-9_-]{10}[AQgw]$
		//
		// See http://schemas.taskcluster.net/hooks/v1/hook-status.json#/properties/lastFire/oneOf[0]/properties/taskId
		TaskID string `json:"taskId"`

		// The time the task was created.  This will not necessarily match `task.created`.
		//
		// See http://schemas.taskcluster.net/hooks/v1/hook-status.json#/properties/lastFire/oneOf[0]/properties/time
		Time tcclient.Time `json:"time"`
	}

	// A representation of **task status** as known by the queue
	//
	// See http://schemas.taskcluster.net/hooks/v1/task-status.json#
	TaskStatusStructure struct {

		// See http://schemas.taskcluster.net/hooks/v1/task-status.json#/properties/status
		Status struct {

			// Use of this field is deprecated; use `deadline: {$fromNow: ..}` in the task template instead.
			//
			// Default:    "1 day"
			//
			// See http://schemas.taskcluster.net/hooks/v1/task-status.json#/properties/status/properties/deadline
			Deadline string `json:"deadline"`

			// Use of this field is deprecated; use `expires: {$fromNow: ..}` in the task template instead.
			//
			// Default:    "3 months"
			//
			// See http://schemas.taskcluster.net/hooks/v1/task-status.json#/properties/status/properties/expires
			Expires string `json:"expires"`

			// Unique identifier for the provisioner that this task must be scheduled on
			//
			// Syntax:     ^([a-zA-Z0-9-_]*)$
			// Min length: 1
			// Max length: 22
			//
			// See http://schemas.taskcluster.net/hooks/v1/task-status.json#/properties/status/properties/provisionerId
			ProvisionerID string `json:"provisionerId"`

			// Number of retries left for the task in case of infrastructure issues
			//
			// Mininum:    0
			// Maximum:    999
			//
			// See http://schemas.taskcluster.net/hooks/v1/task-status.json#/properties/status/properties/retriesLeft
			RetriesLeft int `json:"retriesLeft"`

			// List of runs, ordered so that index `i` has `runId == i`
			//
			// See http://schemas.taskcluster.net/hooks/v1/task-status.json#/properties/status/properties/runs
			Runs []struct {

				// Reason for the creation of this run,
				// **more reasons may be added in the future**.
				//
				// Possible values:
				//   * "scheduled"
				//   * "retry"
				//   * "task-retry"
				//   * "rerun"
				//   * "exception"
				//
				// See http://schemas.taskcluster.net/hooks/v1/task-status.json#/properties/status/properties/runs/items/properties/reasonCreated
				ReasonCreated string `json:"reasonCreated"`

				// Reason that run was resolved, this is mainly
				// useful for runs resolved as `exception`.
				// Note, **more reasons may be added in the future**, also this
				// property is only available after the run is resolved.
				//
				// Possible values:
				//   * "completed"
				//   * "failed"
				//   * "deadline-exceeded"
				//   * "canceled"
				//   * "superseded"
				//   * "claim-expired"
				//   * "worker-shutdown"
				//   * "malformed-payload"
				//   * "resource-unavailable"
				//   * "internal-error"
				//   * "intermittent-task"
				//
				// See http://schemas.taskcluster.net/hooks/v1/task-status.json#/properties/status/properties/runs/items/properties/reasonResolved
				ReasonResolved string `json:"reasonResolved,omitempty"`

				// Date-time at which this run was resolved, ie. when the run changed
				// state from `running` to either `completed`, `failed` or `exception`.
				// This property is only present after the run as been resolved.
				//
				// See http://schemas.taskcluster.net/hooks/v1/task-status.json#/properties/status/properties/runs/items/properties/resolved
				Resolved tcclient.Time `json:"resolved,omitempty"`

				// Id of this task run, `run-id`s always starts from `0`
				//
				// Mininum:    0
				// Maximum:    1000
				//
				// See http://schemas.taskcluster.net/hooks/v1/task-status.json#/properties/status/properties/runs/items/properties/runId
				RunID int `json:"runId"`

				// Date-time at which this run was scheduled, ie. when the run was
				// created in state `pending`.
				//
				// See http://schemas.taskcluster.net/hooks/v1/task-status.json#/properties/status/properties/runs/items/properties/scheduled
				Scheduled tcclient.Time `json:"scheduled"`

				// Date-time at which this run was claimed, ie. when the run changed
				// state from `pending` to `running`. This property is only present
				// after the run has been claimed.
				//
				// See http://schemas.taskcluster.net/hooks/v1/task-status.json#/properties/status/properties/runs/items/properties/started
				Started tcclient.Time `json:"started,omitempty"`

				// State of this run
				//
				// Possible values:
				//   * "pending"
				//   * "running"
				//   * "completed"
				//   * "failed"
				//   * "exception"
				//
				// See http://schemas.taskcluster.net/hooks/v1/task-status.json#/properties/status/properties/runs/items/properties/state
				State string `json:"state"`

				// Time at which the run expires and is resolved as `failed`, if the
				// run isn't reclaimed. Note, only present after the run has been
				// claimed.
				//
				// See http://schemas.taskcluster.net/hooks/v1/task-status.json#/properties/status/properties/runs/items/properties/takenUntil
				TakenUntil tcclient.Time `json:"takenUntil,omitempty"`

				// Identifier for group that worker who executes this run is a part of,
				// this identifier is mainly used for efficient routing.
				// Note, this property is only present after the run is claimed.
				//
				// Syntax:     ^([a-zA-Z0-9-_]*)$
				// Min length: 1
				// Max length: 22
				//
				// See http://schemas.taskcluster.net/hooks/v1/task-status.json#/properties/status/properties/runs/items/properties/workerGroup
				WorkerGroup string `json:"workerGroup,omitempty"`

				// Identifier for worker evaluating this run within given
				// `workerGroup`. Note, this property is only available after the run
				// has been claimed.
				//
				// Syntax:     ^([a-zA-Z0-9-_]*)$
				// Min length: 1
				// Max length: 22
				//
				// See http://schemas.taskcluster.net/hooks/v1/task-status.json#/properties/status/properties/runs/items/properties/workerId
				WorkerID string `json:"workerId,omitempty"`
			} `json:"runs"`

			// Identifier for the scheduler that _defined_ this task.
			//
			// Syntax:     ^([a-zA-Z0-9-_]*)$
			// Min length: 1
			// Max length: 22
			//
			// See http://schemas.taskcluster.net/hooks/v1/task-status.json#/properties/status/properties/schedulerId
			SchedulerID string `json:"schedulerId"`

			// State of this task. This is just an auxiliary property derived from state
			// of latests run, or `unscheduled` if none.
			//
			// Possible values:
			//   * "unscheduled"
			//   * "pending"
			//   * "running"
			//   * "completed"
			//   * "failed"
			//   * "exception"
			//
			// See http://schemas.taskcluster.net/hooks/v1/task-status.json#/properties/status/properties/state
			State string `json:"state"`

			// Identifier for a group of tasks scheduled together with this task, by
			// scheduler identified by `schedulerId`. For tasks scheduled by the
			// task-graph scheduler, this is the `taskGraphId`.
			//
			// Syntax:     ^[A-Za-z0-9_-]{8}[Q-T][A-Za-z0-9_-][CGKOSWaeimquy26-][A-Za-z0-9_-]{10}[AQgw]$
			//
			// See http://schemas.taskcluster.net/hooks/v1/task-status.json#/properties/status/properties/taskGroupId
			TaskGroupID string `json:"taskGroupId"`

			// Unique task identifier, this is UUID encoded as
			// [URL-safe base64](http://tools.ietf.org/html/rfc4648#section-5) and
			// stripped of `=` padding.
			//
			// Syntax:     ^[A-Za-z0-9_-]{8}[Q-T][A-Za-z0-9_-][CGKOSWaeimquy26-][A-Za-z0-9_-]{10}[AQgw]$
			//
			// See http://schemas.taskcluster.net/hooks/v1/task-status.json#/properties/status/properties/taskId
			TaskID string `json:"taskId"`

			// Identifier for worker type within the specified provisioner
			//
			// Syntax:     ^([a-zA-Z0-9-_]*)$
			// Min length: 1
			// Max length: 22
			//
			// See http://schemas.taskcluster.net/hooks/v1/task-status.json#/properties/status/properties/workerType
			WorkerType string `json:"workerType"`
		} `json:"status"`
	}

	// Trigger context
	//
	// See http://schemas.taskcluster.net/hooks/v1/trigger-context.json#
	TriggerContext json.RawMessage

	// Secret token for a trigger
	//
	// See http://schemas.taskcluster.net/hooks/v1/trigger-token-response.json#
	TriggerTokenResponse struct {

		// See http://schemas.taskcluster.net/hooks/v1/trigger-token-response.json#/properties/token
		Token string `json:"token"`
	}
)

// MarshalJSON calls json.RawMessage method of the same name. Required since
// TriggerContext is of type json.RawMessage...
func (this *TriggerContext) MarshalJSON() ([]byte, error) {
	x := json.RawMessage(*this)
	return (&x).MarshalJSON()
}

// UnmarshalJSON is a copy of the json.RawMessage implementation.
func (this *TriggerContext) UnmarshalJSON(data []byte) error {
	if this == nil {
		return errors.New("TriggerContext: UnmarshalJSON on nil pointer")
	}
	*this = append((*this)[0:0], data...)
	return nil
}
