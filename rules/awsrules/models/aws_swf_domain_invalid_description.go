// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/wata727/tflint/tflint"
)

// AwsSwfDomainInvalidDescriptionRule checks the pattern is valid
type AwsSwfDomainInvalidDescriptionRule struct {
	resourceType  string
	attributeName string
	max           int
}

// NewAwsSwfDomainInvalidDescriptionRule returns new rule with default attributes
func NewAwsSwfDomainInvalidDescriptionRule() *AwsSwfDomainInvalidDescriptionRule {
	return &AwsSwfDomainInvalidDescriptionRule{
		resourceType:  "aws_swf_domain",
		attributeName: "description",
		max:           1024,
	}
}

// Name returns the rule name
func (r *AwsSwfDomainInvalidDescriptionRule) Name() string {
	return "aws_swf_domain_invalid_description"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsSwfDomainInvalidDescriptionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsSwfDomainInvalidDescriptionRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsSwfDomainInvalidDescriptionRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsSwfDomainInvalidDescriptionRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"description must be 1024 characters or less",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
