package detector

import (
	"fmt"

	"github.com/wata727/tflint/issue"
)

type AwsDBInstanceInvalidTypeDetector struct {
	*Detector
}

func (d *Detector) CreateAwsDBInstanceInvalidTypeDetector() *AwsDBInstanceInvalidTypeDetector {
	return &AwsDBInstanceInvalidTypeDetector{d}
}

func (d *AwsDBInstanceInvalidTypeDetector) Detect(issues *[]*issue.Issue) {
	var validInstanceType = map[string]bool{
		"db.t2.micro":    true,
		"db.t2.small":    true,
		"db.t2.medium":   true,
		"db.t2.large":    true,
		"db.m4.large":    true,
		"db.m4.xlarge":   true,
		"db.m4.2xlarge":  true,
		"db.m4.4xlarge":  true,
		"db.m4.10xlarge": true,
		"db.m3.medium":   true,
		"db.m3.large":    true,
		"db.m3.xlarge":   true,
		"db.m3.2xlarge":  true,
		"db.r3.large":    true,
		"db.r3.xlarge":   true,
		"db.r3.2xlarge":  true,
		"db.r3.4xlarge":  true,
		"db.r3.8xlarge":  true,
		"db.t1.micro":    true,
		"db.m1.small":    true,
		"db.m1.medium":   true,
		"db.m1.large":    true,
		"db.m1.xlarge":   true,
		"db.m2.xlarge":   true,
		"db.m2.2xlarge":  true,
		"db.m2.4xlarge":  true,
		"db.cr1.8xlarge": true,
	}

	for filename, list := range d.ListMap {
		for _, item := range list.Filter("resource", "aws_db_instance").Items {
			instanceTypeToken, err := hclLiteralToken(item, "instance_class")
			if err != nil {
				d.Logger.Error(err)
				continue
			}
			instanceType, err := d.evalToString(instanceTypeToken.Text)
			if err != nil {
				d.Logger.Error(err)
				continue
			}

			if !validInstanceType[instanceType] {
				issue := &issue.Issue{
					Type:    "ERROR",
					Message: fmt.Sprintf("\"%s\" is invalid instance type.", instanceType),
					Line:    instanceTypeToken.Pos.Line,
					File:    filename,
				}
				*issues = append(*issues, issue)
			}
		}
	}
}