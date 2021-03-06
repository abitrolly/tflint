// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/wata727/tflint/tflint"
)

// AwsStoragegatewayGatewayInvalidGatewayTimezoneRule checks the pattern is valid
type AwsStoragegatewayGatewayInvalidGatewayTimezoneRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsStoragegatewayGatewayInvalidGatewayTimezoneRule returns new rule with default attributes
func NewAwsStoragegatewayGatewayInvalidGatewayTimezoneRule() *AwsStoragegatewayGatewayInvalidGatewayTimezoneRule {
	return &AwsStoragegatewayGatewayInvalidGatewayTimezoneRule{
		resourceType:  "aws_storagegateway_gateway",
		attributeName: "gateway_timezone",
		max:           10,
		min:           3,
	}
}

// Name returns the rule name
func (r *AwsStoragegatewayGatewayInvalidGatewayTimezoneRule) Name() string {
	return "aws_storagegateway_gateway_invalid_gateway_timezone"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsStoragegatewayGatewayInvalidGatewayTimezoneRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsStoragegatewayGatewayInvalidGatewayTimezoneRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsStoragegatewayGatewayInvalidGatewayTimezoneRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsStoragegatewayGatewayInvalidGatewayTimezoneRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"gateway_timezone must be 10 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"gateway_timezone must be 3 characters or higher",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
