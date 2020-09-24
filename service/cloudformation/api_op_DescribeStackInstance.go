// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudformation

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the stack instance that's associated with the specified stack set, AWS
// account, and Region. For a list of stack instances that are associated with a
// specific stack set, use ListStackInstances ().
func (c *Client) DescribeStackInstance(ctx context.Context, params *DescribeStackInstanceInput, optFns ...func(*Options)) (*DescribeStackInstanceOutput, error) {
	stack := middleware.NewStack("DescribeStackInstance", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpDescribeStackInstanceMiddlewares(stack)
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
	addOpDescribeStackInstanceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeStackInstance(options.Region), middleware.Before)
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
			OperationName: "DescribeStackInstance",
			Err:           err,
		}
	}
	out := result.(*DescribeStackInstanceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeStackInstanceInput struct {
	// The name or the unique stack ID of the stack set that you want to get stack
	// instance information for.
	StackSetName *string
	// The ID of an AWS account that's associated with this stack instance.
	StackInstanceAccount *string
	// The name of a Region that's associated with this stack instance.
	StackInstanceRegion *string
}

type DescribeStackInstanceOutput struct {
	// The stack instance that matches the specified request parameters.
	StackInstance *types.StackInstance

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpDescribeStackInstanceMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpDescribeStackInstance{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeStackInstance{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeStackInstance(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudformation",
		OperationName: "DescribeStackInstance",
	}
}
