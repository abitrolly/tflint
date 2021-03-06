// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/wata727/tflint/tflint"
)

// AwsGlueTriggerInvalidTypeRule checks the pattern is valid
type AwsGlueTriggerInvalidTypeRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsGlueTriggerInvalidTypeRule returns new rule with default attributes
func NewAwsGlueTriggerInvalidTypeRule() *AwsGlueTriggerInvalidTypeRule {
	return &AwsGlueTriggerInvalidTypeRule{
		resourceType:  "aws_glue_trigger",
		attributeName: "type",
		enum: []string{
			"SCHEDULED",
			"CONDITIONAL",
			"ON_DEMAND",
		},
	}
}

// Name returns the rule name
func (r *AwsGlueTriggerInvalidTypeRule) Name() string {
	return "aws_glue_trigger_invalid_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsGlueTriggerInvalidTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsGlueTriggerInvalidTypeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsGlueTriggerInvalidTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsGlueTriggerInvalidTypeRule) Check(runner *tflint.Runner) error {
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
					`type is not a valid value`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
