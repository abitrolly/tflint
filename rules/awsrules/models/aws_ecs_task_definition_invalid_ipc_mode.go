// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/wata727/tflint/tflint"
)

// AwsEcsTaskDefinitionInvalidIpcModeRule checks the pattern is valid
type AwsEcsTaskDefinitionInvalidIpcModeRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsEcsTaskDefinitionInvalidIpcModeRule returns new rule with default attributes
func NewAwsEcsTaskDefinitionInvalidIpcModeRule() *AwsEcsTaskDefinitionInvalidIpcModeRule {
	return &AwsEcsTaskDefinitionInvalidIpcModeRule{
		resourceType:  "aws_ecs_task_definition",
		attributeName: "ipc_mode",
		enum: []string{
			"host",
			"task",
			"none",
		},
	}
}

// Name returns the rule name
func (r *AwsEcsTaskDefinitionInvalidIpcModeRule) Name() string {
	return "aws_ecs_task_definition_invalid_ipc_mode"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsEcsTaskDefinitionInvalidIpcModeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsEcsTaskDefinitionInvalidIpcModeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsEcsTaskDefinitionInvalidIpcModeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsEcsTaskDefinitionInvalidIpcModeRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssue(
					r,
					`ipc_mode is not a valid value`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
