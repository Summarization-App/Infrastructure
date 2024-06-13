package main

import (
	eks "infrastructure/aws_services/eks"
	iam "infrastructure/aws_services/iam"
	"infrastructure/aws_services/subnets"
	vpc "infrastructure/aws_services/vpc"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create an AWS resource (S3 Bucket)
		// bucket, err := awsservices.CreateS3Bucket(ctx)

		// if(err != nil){
		// 	return err
		// }

		// // Export the name of the bucket
		// ctx.Export("bucketName", bucket.ID())
		// return nil
		cfg := config.New(ctx, "")

		vpc, err := vpc.CreateVPC(ctx, cfg)

		if err != nil {
			return err
		}

		subnets, errors := subnets.CreateSubnets(ctx, cfg, vpc)

		if errors[0] != nil {
			return errors[0]
		}

		// securityGroup, error := securitygroups.Security_Group_EKS(ctx, vpc)

		// if error != nil {
		// 	return error
		// }

		eksRole, err := iam.EKSRole(ctx)

		if err != nil {
			return err
		}

		eks_fargate, err := iam.CreateFargateRole(ctx)

		// ec2NodeRole, err := iam.CreateEC2Role(ctx)

		if err != nil {
			return err
		}

		// eks_cluster, err := eks.CreateEKSCluster(ctx, vpc, subnets, eksRole, ec2NodeRole)

		eks_cluster, err := eks.CreateEKSCluster(ctx, vpc, subnets, eksRole, eks_fargate)
		// eks_cluster, err := eks.CreateEKSCluster2(ctx, vpc, subnets, securityGroup)

		if err != nil {
			return err
		}

		ctx.Export("vpc name", vpc.ID())
		ctx.Export("First Subnet", subnets[0].ID())
		ctx.Export("EKS Name", eks_cluster.Arn)

		return nil

	})
}
