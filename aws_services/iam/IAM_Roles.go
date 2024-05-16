package aws_services

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func CreateRole(ctx *pulumi.Context, policy *iam.Policy, name string) (*iam.Role, error)  {

	role, err1 := iam.NewRole(ctx, name, &iam.RoleArgs{
		Name:             pulumi.String(name),
		AssumeRolePolicy: pulumi.String(`{
			"Version": "2012-10-17",
			"Statement": [
				{
					"Effect": "Allow",
					"Principal": {
						"Service": 
						[
							"codepipeline.amazonaws.com",
							"ec2.amazonaws.com"
						]
					},
					"Action": "sts:AssumeRole"
				}
			]
		}`),
		Tags: pulumi.StringMap{
			"Tech": pulumi.String(name),
		},
	})

	iam.NewRolePolicyAttachment(ctx, name, &iam.RolePolicyAttachmentArgs{
		PolicyArn: policy.Arn,
		Role: role.Name,
	})

	return role, err1

}