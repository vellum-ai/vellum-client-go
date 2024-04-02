// This file was auto-generated by Fern from our API Definition.

package api

type TestSuiteRunCreateRequest struct {
	// The ID of the Test Suite to run
	TestSuiteId *string `json:"test_suite_id,omitempty"`
	// Configuration that defines how the Test Suite should be run
	ExecConfig *TestSuiteRunExecConfigRequest `json:"exec_config,omitempty"`
}

type ListTestSuiteRunExecutionsRequest struct {
	// Number of results to return per page.
	Limit *int `json:"-"`
	// The initial index from which to return the results.
	Offset *int `json:"-"`
}