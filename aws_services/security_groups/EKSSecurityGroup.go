package securitygroups

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/vpc"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func Security_Group_EKS(ctx *pulumi.Context, inputVpc *ec2.Vpc) (*ec2.SecurityGroup, error) {

	allowTls, err := ec2.NewSecurityGroup(ctx, "allow_tls", &ec2.SecurityGroupArgs{
		Name:        pulumi.String("allow_tls"),
		Description: pulumi.String("Allow TLS inbound traffic and all outbound traffic"),
		VpcId:       inputVpc.ID().ToIDOutput(),
		Tags: pulumi.StringMap{
			"Name": pulumi.String("allow_tls"),
		},
	})

	_, err = vpc.NewSecurityGroupIngressRule(ctx, "allow_tls_ipv4", &vpc.SecurityGroupIngressRuleArgs{
		SecurityGroupId: allowTls.ID(),
		CidrIpv4:        inputVpc.CidrBlock.ToStringOutput(),
		FromPort:        pulumi.Int(443),
		IpProtocol:      pulumi.String("tcp"),
		ToPort:          pulumi.Int(443),
	})

	_, err = vpc.NewSecurityGroupIngressRule(ctx, "allow_tls_ipv6", &vpc.SecurityGroupIngressRuleArgs{
		SecurityGroupId: allowTls.ID(),
		CidrIpv6:        inputVpc.Ipv6CidrBlock.ToStringOutput(),
		FromPort:        pulumi.Int(443),
		IpProtocol:      pulumi.String("tcp"),
		ToPort:          pulumi.Int(443),
	})

	_, err = vpc.NewSecurityGroupEgressRule(ctx, "allow_all_traffic_ipv4", &vpc.SecurityGroupEgressRuleArgs{
		SecurityGroupId: allowTls.ID(),
		CidrIpv4:        pulumi.String("0.0.0.0/0"),
		IpProtocol:      pulumi.String("-1"),
	})

	_, err = vpc.NewSecurityGroupEgressRule(ctx, "allow_all_traffic_ipv6", &vpc.SecurityGroupEgressRuleArgs{
		SecurityGroupId: allowTls.ID(),
		CidrIpv6:        pulumi.String("::/0"),
		IpProtocol:      pulumi.String("-1"),
	})

	return allowTls, err

}
