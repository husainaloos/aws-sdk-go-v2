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

// Returns AWS resource descriptions for running and deleted stacks. If StackName
// is specified, all the associated resources that are part of the stack are
// returned. If PhysicalResourceId is specified, the associated resources of the
// stack that the resource belongs to are returned. Only the first 100 resources
// will be returned. If your stack has more resources than this, you should use
// ListStackResources instead. For deleted stacks, DescribeStackResources returns
// resource information for up to 90 days after the stack has been deleted. You
// must specify either StackName or PhysicalResourceId, but not both. In addition,
// you can specify LogicalResourceId to filter the returned result. For more
// information about resources, the LogicalResourceId and PhysicalResourceId, go to
// the AWS CloudFormation User Guide
// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/). A
// ValidationError is returned if you specify both StackName and PhysicalResourceId
// in the same request.
func (c *Client) DescribeStackResources(ctx context.Context, params *DescribeStackResourcesInput, optFns ...func(*Options)) (*DescribeStackResourcesOutput, error) {
	stack := middleware.NewStack("DescribeStackResources", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpDescribeStackResourcesMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeStackResources(options.Region), middleware.Before)
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
			OperationName: "DescribeStackResources",
			Err:           err,
		}
	}
	out := result.(*DescribeStackResourcesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for DescribeStackResources () action.
type DescribeStackResourcesInput struct {
	// The name or unique identifier that corresponds to a physical instance ID of a
	// resource supported by AWS CloudFormation. For example, for an Amazon Elastic
	// Compute Cloud (EC2) instance, PhysicalResourceId corresponds to the InstanceId.
	// You can pass the EC2 InstanceId to DescribeStackResources to find which stack
	// the instance belongs to and what other resources are part of the stack.
	// Required: Conditional. If you do not specify PhysicalResourceId, you must
	// specify StackName. Default: There is no default value.
	PhysicalResourceId *string
	// The logical name of the resource as specified in the template. Default: There is
	// no default value.
	LogicalResourceId *string
	// The name or the unique stack ID that is associated with the stack, which are not
	// always interchangeable:
	//
	//     * Running stacks: You can specify either the
	// stack's name or its unique stack ID.
	//
	//     * Deleted stacks: You must specify the
	// unique stack ID.
	//
	// Default: There is no default value. Required: Conditional. If
	// you do not specify StackName, you must specify PhysicalResourceId.
	StackName *string
}

// The output for a DescribeStackResources () action.
type DescribeStackResourcesOutput struct {
	// A list of StackResource structures.
	StackResources []*types.StackResource

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpDescribeStackResourcesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpDescribeStackResources{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeStackResources{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeStackResources(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudformation",
		OperationName: "DescribeStackResources",
	}
}
