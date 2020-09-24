// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Describes properties of scheduled actions.
func (c *Client) DescribeScheduledActions(ctx context.Context, params *DescribeScheduledActionsInput, optFns ...func(*Options)) (*DescribeScheduledActionsOutput, error) {
	stack := middleware.NewStack("DescribeScheduledActions", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpDescribeScheduledActionsMiddlewares(stack)
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
	addOpDescribeScheduledActionsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeScheduledActions(options.Region), middleware.Before)
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
			OperationName: "DescribeScheduledActions",
			Err:           err,
		}
	}
	out := result.(*DescribeScheduledActionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeScheduledActionsInput struct {
	// The start time in UTC of the scheduled actions to retrieve. Only active
	// scheduled actions that have invocations after this time are retrieved.
	StartTime *time.Time
	// If true, retrieve only active scheduled actions. If false, retrieve only
	// disabled scheduled actions.
	Active *bool
	// The name of the scheduled action to retrieve.
	ScheduledActionName *string
	// An optional parameter that specifies the starting point to return a set of
	// response records. When the results of a DescribeScheduledActions () request
	// exceed the value specified in MaxRecords, AWS returns a value in the Marker
	// field of the response. You can retrieve the next set of response records by
	// providing the returned marker value in the Marker parameter and retrying the
	// request.
	Marker *string
	// The maximum number of response records to return in each call. If the number of
	// remaining response records exceeds the specified MaxRecords value, a value is
	// returned in a marker field of the response. You can retrieve the next set of
	// records by retrying the command with the returned marker value. Default: 100
	// Constraints: minimum 20, maximum 100.
	MaxRecords *int32
	// The end time in UTC of the scheduled action to retrieve. Only active scheduled
	// actions that have invocations before this time are retrieved.
	EndTime *time.Time
	// The type of the scheduled actions to retrieve.
	TargetActionType types.ScheduledActionTypeValues
	// List of scheduled action filters.
	Filters []*types.ScheduledActionFilter
}

type DescribeScheduledActionsOutput struct {
	// List of retrieved scheduled actions.
	ScheduledActions []*types.ScheduledAction
	// An optional parameter that specifies the starting point to return a set of
	// response records. When the results of a DescribeScheduledActions () request
	// exceed the value specified in MaxRecords, AWS returns a value in the Marker
	// field of the response. You can retrieve the next set of response records by
	// providing the returned marker value in the Marker parameter and retrying the
	// request.
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpDescribeScheduledActionsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpDescribeScheduledActions{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeScheduledActions{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeScheduledActions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift",
		OperationName: "DescribeScheduledActions",
	}
}
