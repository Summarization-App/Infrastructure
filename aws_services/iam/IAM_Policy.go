package aws_services

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func CreateIAMPolicy(ctx *pulumi.Context, bucket *s3.Bucket) (*iam.Policy, error) {

	policy, err := iam.NewPolicy(ctx, "myPolicy", &iam.PolicyArgs{
		Description: pulumi.String("My custom IAM policy"),
		Policy: pulumi.Sprintf(`{
			"Version": "2012-10-17",
			"Statement": [
				{
					"Effect": "Allow",
					"Action": "s3:GetObject",
					"Resource": "%s/*"
				},
				{
					"Effect": "Allow",
					"Action": "ec2:Describe*",
					"Resource": "*"
				}
			]
		}`, bucket.Arn).ToStringOutput(),
	})

	return policy, err

}