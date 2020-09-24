// Code generated by smithy-go-codegen DO NOT EDIT.

package sesv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// List the dedicated IP addresses that are associated with your AWS account.
func (c *Client) GetDedicatedIps(ctx context.Context, params *GetDedicatedIpsInput, optFns ...func(*Options)) (*GetDedicatedIpsOutput, error) {
	stack := middleware.NewStack("GetDedicatedIps", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetDedicatedIpsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetDedicatedIps(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)

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
			OperationName: "GetDedicatedIps",
			Err:           err,
		}
	}
	out := result.(*GetDedicatedIpsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to obtain more information about dedicated IP pools.
type GetDedicatedIpsInput struct {
	// The number of results to show in a single call to GetDedicatedIpsRequest. If the
	// number of results is larger than the number you specified in this parameter,
	// then the response includes a NextToken element, which you can use to obtain
	// additional results.
	PageSize *int32
	// A token returned from a previous call to GetDedicatedIps to indicate the
	// position of the dedicated IP pool in the list of IP pools.
	NextToken *string
	// The name of the IP pool that the dedicated IP address is associated with.
	PoolName *string
}

// Information about the dedicated IP addresses that are associated with your AWS
// account.
type GetDedicatedIpsOutput struct {
	// A token that indicates that there are additional dedicated IP addresses to list.
	// To view additional addresses, issue another request to GetDedicatedIps, passing
	// this token in the NextToken parameter.
	NextToken *string
	// A list of dedicated IP addresses that are associated with your AWS account.
	DedicatedIps []*types.DedicatedIp

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetDedicatedIpsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetDedicatedIps{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetDedicatedIps{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetDedicatedIps(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "GetDedicatedIps",
	}
}
