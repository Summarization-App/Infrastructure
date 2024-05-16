package eks

import (
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/eks"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func CreateEKSCluster(ctx *pulumi.Context, vpc *ec2.Vpc, subnets []*ec2.Subnet, role *iam.Role, ec2NodeRole *iam.Role) (*eks.Cluster, error) {

	cluster, err := eks.NewCluster(ctx, "My_EKS_Cluster", &eks.ClusterArgs{
		Name:    pulumi.String("My_EKS_Cluster"),
		RoleArn: role.Arn,
		VpcConfig: &eks.ClusterVpcConfigArgs{
			// VpcId: vpc.ID(),
			SubnetIds: pulumi.StringArray{
				subnets[0].ID(),
				subnets[1].ID(),
				subnets[2].ID(),
			},
			EndpointPublicAccess: pulumi.Bool(true),
		},		
		
	})

	// eks.NewNodeGroup(ctx, "Node_Group", &eks.NodeGroupArgs{
	// 	ClusterName: cluster.Name,
	// 	NodeRoleArn: ec2NodeRole.Arn,
	// 	NodeGroupName: pulumi.String("My_Node_Group"),
	// 	SubnetIds: pulumi.StringArray{
	// 		subnets[0].ID(),
	// 		subnets[1].ID(),
	// 		subnets[2].ID(),
	// 	},
	// 	ScalingConfig: &eks.NodeGroupScalingConfigArgs{
	// 		DesiredSize: pulumi.Int(2),
	// 		MaxSize: pulumi.Int(3),
	// 		MinSize: pulumi.Int(1),
	// 	},
	// 	InstanceTypes: pulumi.StringArray{
	// 		pulumi.String("t2.medium"),
	// 	},	
	// })

	eks.NewFargateProfile(ctx, "FargateProfile", &eks.FargateProfileArgs{
		ClusterName: cluster.Name,
		FargateProfileName: pulumi.String("Fargate_Profile_Pulumi"),
		PodExecutionRoleArn: ec2NodeRole.Arn,
		SubnetIds: pulumi.StringArray{
			subnets[0].ID(),
			subnets[1].ID(),
			subnets[2].ID(),
		},
		Selectors: eks.FargateProfileSelectorArray{
			&eks.FargateProfileSelectorArgs{
				Namespace: pulumi.String("default"),
			},
		},
	})

	// eksProvider, err := kubernetes.NewProvider(ctx, "eks-provider", &kubernetes.ProviderArgs{
	// 	Kubeconfig: cluster.GetKube,
	// })



	// opts := &helm.v3.ChartOptions{
	// 	Chart:   "path/to/helm/chart", // Path to submodule
	// 	Version: "1.2.3",              // Version of the Helm chart
	// 	Values:  types.Values{},       // Optional: Values to override default Helm chart values
	// }

	// opts := &helm.ChartArgs{
	// 	Chart: pulumi.StringInput(pulumi.String("./Helm_Chart")),
	// 	Version: pulumi.StringInput(pulumi.String("1.2.3")),
	// 	Values: types.Values{},
	// }

	// helm.NewChart(ctx, "kubernetes-nginx", helm.ChartArgs{
	// 	Chart: opts.Chart,
	// 	Version: opts.Version,
	// 	Transformations: []yaml.Transformation{
	// 		// func(state *helm.TransformedState) *helm.TransformedState {
	// 		// 	state.Kubeconfig = cluster.Kubeconfig
	// 		// 	return state
	// 		// },
	// 		func(state *helm.TransformedState) *helm.TransformedState {
	// 			state.Kubeconfig = cluster.Kubeconfig
	// 			return state
	// 		},
	// 	},

	// })


	return cluster, err

}