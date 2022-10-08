// Code generated by actionlint/scripts/generate-availability. DO NOT EDIT.

package actionlint

// WorkflowKeyAvailability returns contexts and special functions availability of the given workflow key.
// 1st return value indicates what contexts are available. Empty slice means any contexts are available.
// 2nd return value indicates what special functions are available. Empty slice means no special functions are available.
// The 'key' parameter should represents a workflow key like "jobs.<job_id>.concurrency".
//
// This function was generated from https://docs.github.com/en/actions/learn-github-actions/contexts#context-availability.
// See the script for more details: https://github.com/rhysd/actionlint/blob/main/scripts/generate-availability/
func WorkflowKeyAvailability(key string) ([]string, []string) {
	switch key {
	case "jobs.<job_id>.outputs.<output_id>":
		return []string{"env", "github", "inputs", "job", "matrix", "needs", "runner", "secrets", "steps", "strategy"}, []string{}
	case "jobs.<job_id>.steps.continue-on-error", "jobs.<job_id>.steps.env", "jobs.<job_id>.steps.name", "jobs.<job_id>.steps.run", "jobs.<job_id>.steps.timeout-minutes", "jobs.<job_id>.steps.with", "jobs.<job_id>.steps.working-directory":
		return []string{"env", "github", "inputs", "job", "matrix", "needs", "runner", "secrets", "steps", "strategy"}, []string{"hashfiles"}
	case "jobs.<job_id>.container.env.<env_id>", "jobs.<job_id>.services.<service_id>.env.<env_id>":
		return []string{"env", "github", "inputs", "job", "matrix", "needs", "runner", "secrets", "strategy"}, []string{}
	case "jobs.<job_id>.environment.url":
		return []string{"env", "github", "inputs", "job", "matrix", "needs", "runner", "steps", "strategy"}, []string{}
	case "jobs.<job_id>.steps.if":
		return []string{"env", "github", "inputs", "job", "matrix", "needs", "runner", "steps", "strategy"}, []string{"always", "cancelled", "failure", "hashfiles", "success"}
	case "jobs.<job_id>.container", "jobs.<job_id>.container.credentials", "jobs.<job_id>.services.<service_id>.credentials":
		return []string{"env", "github", "inputs", "matrix", "needs", "secrets", "strategy"}, []string{}
	case "jobs.<job_id>.defaults.run":
		return []string{"env", "github", "inputs", "matrix", "needs", "strategy"}, []string{}
	case "concurrency", "on.workflow_call.inputs.<inputs_id>.default":
		return []string{"github", "inputs"}, []string{}
	case "on.workflow_call.outputs.<output_id>.value":
		return []string{"github", "inputs", "jobs"}, []string{}
	case "jobs.<job_id>.env", "jobs.<job_id>.secrets.<secrets_id>":
		return []string{"github", "inputs", "matrix", "needs", "secrets", "strategy"}, []string{}
	case "jobs.<job_id>.concurrency", "jobs.<job_id>.continue-on-error", "jobs.<job_id>.environment", "jobs.<job_id>.name", "jobs.<job_id>.runs-on", "jobs.<job_id>.services", "jobs.<job_id>.timeout-minutes", "jobs.<job_id>.with.<with_id>":
		return []string{"github", "inputs", "matrix", "needs", "strategy"}, []string{}
	case "jobs.<job_id>.strategy":
		return []string{"github", "inputs", "needs"}, []string{}
	case "jobs.<job_id>.if":
		return []string{"github", "inputs", "needs"}, []string{"always", "cancelled", "failure", "success"}
	case "env":
		return []string{"github", "inputs", "secrets"}, []string{}
	default:
		return nil, nil
	}
}

// SpecialFunctionNames is a map from special function name to available workflow keys.
// Some functions are only available at specific positions. This variable is useful when you want to
// know which functions are special and what workflow keys support them.
//
// This function was generated from https://docs.github.com/en/actions/learn-github-actions/contexts#context-availability.
// See the script for more details: https://github.com/rhysd/actionlint/blob/main/scripts/generate-availability/
var SpecialFunctionNames = map[string][]string{"always": []string{"jobs.<job_id>.if", "jobs.<job_id>.steps.if"}, "cancelled": []string{"jobs.<job_id>.if", "jobs.<job_id>.steps.if"}, "failure": []string{"jobs.<job_id>.if", "jobs.<job_id>.steps.if"}, "hashfiles": []string{"jobs.<job_id>.steps.continue-on-error", "jobs.<job_id>.steps.env", "jobs.<job_id>.steps.if", "jobs.<job_id>.steps.name", "jobs.<job_id>.steps.run", "jobs.<job_id>.steps.timeout-minutes", "jobs.<job_id>.steps.with", "jobs.<job_id>.steps.working-directory"}, "success": []string{"jobs.<job_id>.if", "jobs.<job_id>.steps.if"}}

// For test
var allWorkflowKeys = []string{"concurrency", "env", "jobs.<job_id>.concurrency", "jobs.<job_id>.container", "jobs.<job_id>.container.credentials", "jobs.<job_id>.container.env.<env_id>", "jobs.<job_id>.continue-on-error", "jobs.<job_id>.defaults.run", "jobs.<job_id>.env", "jobs.<job_id>.environment", "jobs.<job_id>.environment.url", "jobs.<job_id>.if", "jobs.<job_id>.name", "jobs.<job_id>.outputs.<output_id>", "jobs.<job_id>.runs-on", "jobs.<job_id>.secrets.<secrets_id>", "jobs.<job_id>.services", "jobs.<job_id>.services.<service_id>.credentials", "jobs.<job_id>.services.<service_id>.env.<env_id>", "jobs.<job_id>.steps.continue-on-error", "jobs.<job_id>.steps.env", "jobs.<job_id>.steps.if", "jobs.<job_id>.steps.name", "jobs.<job_id>.steps.run", "jobs.<job_id>.steps.timeout-minutes", "jobs.<job_id>.steps.with", "jobs.<job_id>.steps.working-directory", "jobs.<job_id>.strategy", "jobs.<job_id>.timeout-minutes", "jobs.<job_id>.with.<with_id>", "on.workflow_call.inputs.<inputs_id>.default", "on.workflow_call.outputs.<output_id>.value"}
