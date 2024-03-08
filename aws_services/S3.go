package awsservices

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func CreateS3Bucket(ctx *pulumi.Context) (*s3.Bucket, error) {

	bucket, err := s3.NewBucket(ctx, "my-bucket", nil)

	return bucket, err
	
}