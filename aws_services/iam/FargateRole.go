package aws_services

import (
	"encoding/json"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func CreateFargateRole(ctx *pulumi.Context) (*iam.Role, error)  {
	
	assumeRolePolicyDocument := map[string]interface{}{
        "Version": "2012-10-17",
        "Statement": []map[string]interface{}{
            {
                "Effect": "Allow",
                "Action": "sts:AssumeRole",
                "Principal": map[string]interface{}{
                    "Service": "eks-fargate-pods.amazonaws.com",
                },
            },
        },
    }

	assumeRolePolicyJSON, _ := json.Marshal(assumeRolePolicyDocument)

	ec2NodeRole, err := iam.NewRole(ctx, "ec2_node", &iam.RoleArgs{
		Name: pulumi.String("EKS-Node-Role"),
		AssumeRolePolicy: pulumi.String(string(assumeRolePolicyJSON)),
	})

	eksNodePolicy, _ := json.Marshal(map[string]interface{}{
		"Version": "2012-10-17",
		"Statement": []map[string]interface{}{
			{
				"Effect":   "Allow",
				"Action":   []string{
					"eks:DescribeCluster", 
					"eks:ListNodegroups", 
					"eks:CreateNodegroup"},
				"Resource": "*",
			},
		},
	})

	eks_describeRole, _ := iam.NewPolicy(ctx, "eks_policy", &iam.PolicyArgs{
		Name: pulumi.String("EKS_Policy"),
		Path: pulumi.String("/"),
		Description: pulumi.String("Created using Pulumi"),
		Policy: pulumi.String(eksNodePolicy),
	})

	iam.NewRolePolicyAttachment(ctx, "attaching_EKS_Node_Policy", &iam.RolePolicyAttachmentArgs{
		PolicyArn: eks_describeRole.Arn,
		Role: ec2NodeRole.Name,
	})

	iam.NewRolePolicyAttachment(ctx, "example-AmazonEKSWorkerNodePolicy", &iam.RolePolicyAttachmentArgs{
		PolicyArn: pulumi.String("arn:aws:iam::aws:policy/AmazonEKSWorkerNodePolicy"),
		Role:      ec2NodeRole.Name,
	})

	iam.NewRolePolicyAttachment(ctx, "example-AmazonEKS_CNI_Policy", &iam.RolePolicyAttachmentArgs{
		PolicyArn: pulumi.String("arn:aws:iam::aws:policy/AmazonEKS_CNI_Policy"),
		Role:      ec2NodeRole.Name,
	})

	// _, err = iam.NewRolePolicyAttachment(ctx, "example-AmazonEC2ContainerRegistryReadOnly", &iam.RolePolicyAttachmentArgs{
	// 	PolicyArn: pulumi.String("arn:aws:iam::aws:policy/AmazonEC2ContainerRegistryReadOnly"),
	// 	Role:      example.Name,
	// })

	return ec2NodeRole, err

}