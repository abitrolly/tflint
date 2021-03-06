// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"testing"

	"github.com/wata727/tflint/tflint"
)

func Test_AwsDatasyncLocationNfsInvalidSubdirectoryRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected tflint.Issues
	}{
		{
			Name: "It includes invalid characters",
			Content: `
resource "aws_datasync_location_nfs" "foo" {
	subdirectory = "/exported^path"
}`,
			Expected: tflint.Issues{
				{
					Rule:    NewAwsDatasyncLocationNfsInvalidSubdirectoryRule(),
					Message: `subdirectory does not match valid pattern ^[a-zA-Z0-9_\-\./]*$`,
				},
			},
		},
		{
			Name: "It is valid",
			Content: `
resource "aws_datasync_location_nfs" "foo" {
	subdirectory = "/exported/path"
}`,
			Expected: tflint.Issues{},
		},
	}

	rule := NewAwsDatasyncLocationNfsInvalidSubdirectoryRule()

	for _, tc := range cases {
		runner := tflint.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		tflint.AssertIssuesWithoutRange(t, tc.Expected, runner.Issues)
	}
}
