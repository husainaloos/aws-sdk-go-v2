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

// Deletes one or more specified VPC endpoints. Deleting a gateway endpoint also
// deletes the endpoint routes in the route tables that were associated with the
// endpoint. Deleting an interface endpoint deletes the endpoint network
// interfaces.
func (c *Client) DeleteVpcEndpoints(ctx context.Context, params *DeleteVpcEndpointsInput, optFns ...func(*Options)) (*DeleteVpcEndpointsOutput, error) {
	stack := middleware.NewStack("DeleteVpcEndpoints", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpDeleteVpcEndpointsMiddlewares(stack)
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
	addOpDeleteVpcEndpointsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteVpcEndpoints(options.Region), middleware.Before)
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
			OperationName: "DeleteVpcEndpoints",
			Err:           err,
		}
	}
	out := result.(*DeleteVpcEndpointsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for DeleteVpcEndpoints.
type DeleteVpcEndpointsInput struct {
	// One or more VPC endpoint IDs.
	VpcEndpointIds []*string
	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
}

// Contains the output of DeleteVpcEndpoints.
type DeleteVpcEndpointsOutput struct {
	// Information about the VPC endpoints that were not successfully deleted.
	Unsuccessful []*types.UnsuccessfulItem

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpDeleteVpcEndpointsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpDeleteVpcEndpoints{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpDeleteVpcEndpoints{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteVpcEndpoints(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DeleteVpcEndpoints",
	}
}
