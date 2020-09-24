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

// Describes the credit option for CPU usage of the specified burstable performance
// instances. The credit options are standard and unlimited. If you do not specify
// an instance ID, Amazon EC2 returns burstable performance instances with the
// unlimited credit option, as well as instances that were previously configured as
// T2, T3, and T3a with the unlimited credit option. For example, if you resize a
// T2 instance, while it is configured as unlimited, to an M4 instance, Amazon EC2
// returns the M4 instance. If you specify one or more instance IDs, Amazon EC2
// returns the credit option (standard or unlimited) of those instances. If you
// specify an instance ID that is not valid, such as an instance that is not a
// burstable performance instance, an error is returned. Recently terminated
// instances might appear in the returned results. This interval is usually less
// than one hour. If an Availability Zone is experiencing a service disruption and
// you specify instance IDs in the affected zone, or do not specify any instance
// IDs at all, the call fails. If you specify only instance IDs in an unaffected
// zone, the call works normally. For more information, see Burstable performance
// instances
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/burstable-performance-instances.html)
// in the Amazon Elastic Compute Cloud User Guide.
func (c *Client) DescribeInstanceCreditSpecifications(ctx context.Context, params *DescribeInstanceCreditSpecificationsInput, optFns ...func(*Options)) (*DescribeInstanceCreditSpecificationsOutput, error) {
	stack := middleware.NewStack("DescribeInstanceCreditSpecifications", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpDescribeInstanceCreditSpecificationsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeInstanceCreditSpecifications(options.Region), middleware.Before)
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
			OperationName: "DescribeInstanceCreditSpecifications",
			Err:           err,
		}
	}
	out := result.(*DescribeInstanceCreditSpecificationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeInstanceCreditSpecificationsInput struct {
	// The filters.
	//
	//     * instance-id - The ID of the instance.
	Filters []*types.Filter
	// The token to retrieve the next page of results.
	NextToken *string
	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
	// The maximum number of results to return in a single call. To retrieve the
	// remaining results, make another call with the returned NextToken value. This
	// value can be between 5 and 1000. You cannot specify this parameter and the
	// instance IDs parameter in the same call.
	MaxResults *int32
	// The instance IDs. Default: Describes all your instances. Constraints: Maximum
	// 1000 explicitly specified instance IDs.
	InstanceIds []*string
}

type DescribeInstanceCreditSpecificationsOutput struct {
	// Information about the credit option for CPU usage of an instance.
	InstanceCreditSpecifications []*types.InstanceCreditSpecification
	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpDescribeInstanceCreditSpecificationsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpDescribeInstanceCreditSpecifications{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeInstanceCreditSpecifications{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeInstanceCreditSpecifications(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeInstanceCreditSpecifications",
	}
}
