// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Adds the specified ingress rules to a security group. An inbound rule permits
// instances to receive traffic from the specified IPv4 or IPv6 CIDR address
// ranges, or from the instances associated with the specified destination security
// groups. You specify a protocol for each rule (for example, TCP). For TCP and
// UDP, you must also specify the destination port or port range. For ICMP/ICMPv6,
// you must also specify the ICMP/ICMPv6 type and code. You can use -1 to mean all
// types or all codes. Rule changes are propagated to instances within the security
// group as quickly as possible. However, a small delay might occur. For more
// information about VPC security group limits, see Amazon VPC Limits
// (https://docs.aws.amazon.com/vpc/latest/userguide/amazon-vpc-limits.html).
func (c *Client) AuthorizeSecurityGroupIngress(ctx context.Context, params *AuthorizeSecurityGroupIngressInput, optFns ...func(*Options)) (*AuthorizeSecurityGroupIngressOutput, error) {
	stack := middleware.NewStack("AuthorizeSecurityGroupIngress", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpAuthorizeSecurityGroupIngressMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAuthorizeSecurityGroupIngress(options.Region), middleware.Before)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     ServiceID,
			OperationName: "AuthorizeSecurityGroupIngress",
			Err:           err,
		}
	}
	out := result.(*AuthorizeSecurityGroupIngressOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AuthorizeSecurityGroupIngressInput struct {
	// The end of port range for the TCP and UDP protocols, or an ICMP code number. For
	// the ICMP code number, use -1 to specify all codes. If you specify all ICMP
	// types, you must specify all codes. Alternatively, use a set of IP permissions to
	// specify multiple rules and a description for the rule.
	ToPort *int32
	// [EC2-Classic, default VPC] The name of the security group. You must specify
	// either the security group ID or the security group name in the request.
	GroupName *string
	// The start of port range for the TCP and UDP protocols, or an ICMP type number.
	// For the ICMP type number, use -1 to specify all types. If you specify all ICMP
	// types, you must specify all codes. Alternatively, use a set of IP permissions to
	// specify multiple rules and a description for the rule.
	FromPort *int32
	// [nondefault VPC] The AWS account ID for the source security group, if the source
	// security group is in a different account. You can't specify this parameter in
	// combination with the following parameters: the CIDR IP address range, the IP
	// protocol, the start of the port range, and the end of the port range. Creates
	// rules that grant full ICMP, UDP, and TCP access. To create a rule with a
	// specific IP protocol and port range, use a set of IP permissions instead.
	SourceSecurityGroupOwnerId *string
	// [EC2-Classic, default VPC] The name of the source security group. You can't
	// specify this parameter in combination with the following parameters: the CIDR IP
	// address range, the start of the port range, the IP protocol, and the end of the
	// port range. Creates rules that grant full ICMP, UDP, and TCP access. To create a
	// rule with a specific IP protocol and port range, use a set of IP permissions
	// instead. For EC2-VPC, the source security group must be in the same VPC.
	SourceSecurityGroupName *string
	// The IPv4 address range, in CIDR format. You can't specify this parameter when
	// specifying a source security group. To specify an IPv6 address range, use a set
	// of IP permissions. Alternatively, use a set of IP permissions to specify
	// multiple rules and a description for the rule.
	CidrIp *string
	// The IP protocol name (tcp, udp, icmp) or number (see Protocol Numbers
	// (http://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml)). To
	// specify icmpv6, use a set of IP permissions. [VPC only] Use -1 to specify all
	// protocols. If you specify -1 or a protocol other than tcp, udp, or icmp, traffic
	// on all ports is allowed, regardless of any ports you specify. Alternatively, use
	// a set of IP permissions to specify multiple rules and a description for the
	// rule.
	IpProtocol *string
	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
	// The sets of IP permissions.
	IpPermissions []*types.IpPermission
	// The ID of the security group. You must specify either the security group ID or
	// the security group name in the request. For security groups in a nondefault VPC,
	// you must specify the security group ID.
	GroupId *string
}

type AuthorizeSecurityGroupIngressOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpAuthorizeSecurityGroupIngressMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpAuthorizeSecurityGroupIngress{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpAuthorizeSecurityGroupIngress{}, middleware.After)
}

func newServiceMetadataMiddleware_opAuthorizeSecurityGroupIngress(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "AuthorizeSecurityGroupIngress",
	}
}