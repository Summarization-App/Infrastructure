package eks

import (
	oldec2 "github.com/pulumi/pulumi-aws/sdk/v5/go/aws/ec2"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi-eks/sdk/go/eks"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func CreateEKSCluster2(ctx *pulumi.Context, vpc *ec2.Vpc, subnets []*ec2.Subnet, securityGroup *oldec2.SecurityGroup) (*eks.Cluster, error) {

	// instanceProfile, err := iam.NewInstanceProfile(ctx, "Instance-profile-1", &iam.InstanceProfileArgs{
	// 	Role: role.Name,
	// })


	// cluster, err := eks.NewCluster(ctx, "EKS Cluster", &eks.ClusterArgs{
	// 	SkipDefaultNodeGroup: pulumi.BoolRef(true),
	// 	InstanceRoles: nil,
	// })

	cluster, err := eks.NewCluster(ctx, "cluster", &eks.ClusterArgs{
		DesiredCapacity: pulumi.Int(5),
		MinSize:         pulumi.Int(3),
		MaxSize:         pulumi.Int(5),
		// EnabledClusterLogTypes: pulumi.StringArray{
		// 	pulumi.String("api"),
		// 	pulumi.String("audit"),
		// 	pulumi.String("authenticator"),
		// },
		SubnetIds: pulumi.StringArray{
			subnets[0].ID(),
			subnets[1].ID(),
			subnets[2].ID(),
		},
		ClusterSecurityGroup: securityGroup,
	})

	// eks.NewNodeGroup(ctx, "Fixed Node Group", &eks.NodeGroupArgs{
	// 	Cluster: cluster,
	// 	InstanceType:    pulumi.String("t2.nano"),
	// 	DesiredCapacity: pulumi.Int(2),
	// 	MinSize:         pulumi.Int(1),
	// 	MaxSize:         pulumi.Int(3),
	// 	SpotPrice:       pulumi.String("1"),
	// 	Labels: map[string]string{
	// 		"ondemand": "true",
	// 	},
	// 	InstanceProfile: instanceProfile,
	// })

	ctx.Export("KubeConfig", cluster.Kubeconfig);

	return cluster, err
}