// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/wata727/tflint/tflint"
)

// AwsElasticBeanstalkEnvironmentInvalidNameRule checks the pattern is valid
type AwsElasticBeanstalkEnvironmentInvalidNameRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsElasticBeanstalkEnvironmentInvalidNameRule returns new rule with default attributes
func NewAwsElasticBeanstalkEnvironmentInvalidNameRule() *AwsElasticBeanstalkEnvironmentInvalidNameRule {
	return &AwsElasticBeanstalkEnvironmentInvalidNameRule{
		resourceType:  "aws_elastic_beanstalk_environment",
		attributeName: "name",
		max:           40,
		min:           4,
	}
}

// Name returns the rule name
func (r *AwsElasticBeanstalkEnvironmentInvalidNameRule) Name() string {
	return "aws_elastic_beanstalk_environment_invalid_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsElasticBeanstalkEnvironmentInvalidNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsElasticBeanstalkEnvironmentInvalidNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsElasticBeanstalkEnvironmentInvalidNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsElasticBeanstalkEnvironmentInvalidNameRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"name must be 40 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"name must be 4 characters or higher",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
