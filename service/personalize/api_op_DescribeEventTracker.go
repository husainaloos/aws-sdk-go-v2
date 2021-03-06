// Code generated by smithy-go-codegen DO NOT EDIT.

package personalize

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/personalize/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes an event tracker. The response includes the trackingId and status of
// the event tracker. For more information on event trackers, see
// CreateEventTracker ().
func (c *Client) DescribeEventTracker(ctx context.Context, params *DescribeEventTrackerInput, optFns ...func(*Options)) (*DescribeEventTrackerOutput, error) {
	if params == nil {
		params = &DescribeEventTrackerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeEventTracker", params, optFns, addOperationDescribeEventTrackerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeEventTrackerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeEventTrackerInput struct {

	// The Amazon Resource Name (ARN) of the event tracker to describe.
	//
	// This member is required.
	EventTrackerArn *string
}

type DescribeEventTrackerOutput struct {

	// An object that describes the event tracker.
	EventTracker *types.EventTracker

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeEventTrackerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeEventTracker{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeEventTracker{}, middleware.After)
	if err != nil {
		return err
	}
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	addResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	addRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpDescribeEventTrackerValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeEventTracker(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeEventTracker(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "personalize",
		OperationName: "DescribeEventTracker",
	}
}
