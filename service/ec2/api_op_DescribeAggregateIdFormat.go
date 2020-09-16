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

// Describes the longer ID format settings for all resource types in a specific
// Region. This request is useful for performing a quick audit to determine whether
// a specific Region is fully opted in for longer IDs (17-character IDs).  <p>This
// request only returns information about resource types that support longer
// IDs.</p> <p>The following resource types support longer IDs: <code>bundle</code>
// | <code>conversion-task</code> | <code>customer-gateway</code> |
// <code>dhcp-options</code> | <code>elastic-ip-allocation</code> |
// <code>elastic-ip-association</code> | <code>export-task</code> |
// <code>flow-log</code> | <code>image</code> | <code>import-task</code> |
// <code>instance</code> | <code>internet-gateway</code> | <code>network-acl</code>
// | <code>network-acl-association</code> | <code>network-interface</code> |
// <code>network-interface-attachment</code> | <code>prefix-list</code> |
// <code>reservation</code> | <code>route-table</code> |
// <code>route-table-association</code> | <code>security-group</code> |
// <code>snapshot</code> | <code>subnet</code> |
// <code>subnet-cidr-block-association</code> | <code>volume</code> |
// <code>vpc</code> | <code>vpc-cidr-block-association</code> |
// <code>vpc-endpoint</code> | <code>vpc-peering-connection</code> |
// <code>vpn-connection</code> | <code>vpn-gateway</code>.</p>
func (c *Client) DescribeAggregateIdFormat(ctx context.Context, params *DescribeAggregateIdFormatInput, optFns ...func(*Options)) (*DescribeAggregateIdFormatOutput, error) {
	stack := middleware.NewStack("DescribeAggregateIdFormat", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpDescribeAggregateIdFormatMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAggregateIdFormat(options.Region), middleware.Before)

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
			OperationName: "DescribeAggregateIdFormat",
			Err:           err,
		}
	}
	out := result.(*DescribeAggregateIdFormatOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAggregateIdFormatInput struct {
	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
}

type DescribeAggregateIdFormatOutput struct {
	// Information about each resource's ID format.
	Statuses []*types.IdFormat
	// Indicates whether all resource types in the Region are configured to use longer
	// IDs. This value is only true if all users are configured to use longer IDs for
	// all resources types in the Region.
	UseLongIdsAggregated *bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpDescribeAggregateIdFormatMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpDescribeAggregateIdFormat{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeAggregateIdFormat{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeAggregateIdFormat(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeAggregateIdFormat",
	}
}