// This file was auto-generated by Fern from our API Definition.

package api

type TestSuiteRunCreateRequest struct {
	// The ID of the Test Suite to run. Must provide either this or test_suite_id.
	TestSuiteId *string `json:"test_suite_id,omitempty" url:"-"`
	// The name of the Test Suite to run. Must provide either this or test_suite_id.
	TestSuiteName *string `json:"test_suite_name,omitempty" url:"-"`
	// Configuration that defines how the Test Suite should be run
	ExecConfig *TestSuiteRunExecConfigRequest `json:"exec_config,omitempty" url:"-"`
}

type TestSuiteRunsListExecutionsRequest struct {
	// The response fields to expand for more information.
	//
	// - 'results.metric_results.metric_label' expands the metric label for each metric result.
	// - 'results.metric_results.metric_definition' expands the metric definition for each metric result.
	// - 'results.metric_results.metric_definition.name' expands the metric definition name for each metric result.
	Expand []*string `json:"-" url:"expand,omitempty"`
	// Number of results to return per page.
	Limit *int `json:"-" url:"limit,omitempty"`
	// The initial index from which to return the results.
	Offset *int `json:"-" url:"offset,omitempty"`
}
