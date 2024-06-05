// This file was auto-generated by Fern from our API Definition.

package api

type ListTestSuiteTestCasesRequest struct {
	// Number of results to return per page.
	Limit *int `json:"-"`
	// The initial index from which to return the results.
	Offset *int `json:"-"`
}

type UpsertTestSuiteTestCaseRequest struct {
	// The ID of the Test Case to upsert. If specified and a match is found, the existing Test Case will be updated. If specified and no match is found, a Test Case will be created with the provided ID. If not provided, a new Test Case will be created with an auto-generated ID.
	UpsertTestSuiteTestCaseRequestId *string `json:"id,omitempty"`
	Label                            *string `json:"label,omitempty"`
	// Values for each of the Test Case's input variables
	InputValues []*NamedTestCaseVariableValueRequest `json:"input_values,omitempty"`
	// Values for each of the Test Case's evaluation variables
	EvaluationValues []*NamedTestCaseVariableValueRequest `json:"evaluation_values,omitempty"`
}
