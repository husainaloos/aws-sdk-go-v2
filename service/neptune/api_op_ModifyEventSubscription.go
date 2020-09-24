// Code generated by smithy-go-codegen DO NOT EDIT.

package neptune

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/neptune/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Modifies an existing event notification subscription. Note that you can't modify
// the source identifiers using this call; to change source identifiers for a
// subscription, use the AddSourceIdentifierToSubscription () and
// RemoveSourceIdentifierFromSubscription () calls. You can see a list of the event
// categories for a given SourceType by using the DescribeEventCategories action.
func (c *Client) ModifyEventSubscription(ctx context.Context, params *ModifyEventSubscriptionInput, optFns ...func(*Options)) (*ModifyEventSubscriptionOutput, error) {
	stack := middleware.NewStack("ModifyEventSubscription", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpModifyEventSubscriptionMiddlewares(stack)
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
	addOpModifyEventSubscriptionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opModifyEventSubscription(options.Region), middleware.Before)
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
			OperationName: "ModifyEventSubscription",
			Err:           err,
		}
	}
	out := result.(*ModifyEventSubscriptionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyEventSubscriptionInput struct {
	// The name of the event notification subscription.
	SubscriptionName *string
	// The type of source that is generating the events. For example, if you want to be
	// notified of events generated by a DB instance, you would set this parameter to
	// db-instance. if this value is not specified, all events are returned. Valid
	// values: db-instance | db-parameter-group | db-security-group | db-snapshot
	SourceType *string
	// A Boolean value; set to true to activate the subscription.
	Enabled *bool
	// The Amazon Resource Name (ARN) of the SNS topic created for event notification.
	// The ARN is created by Amazon SNS when you create a topic and subscribe to it.
	SnsTopicArn *string
	// A list of event categories for a SourceType that you want to subscribe to. You
	// can see a list of the categories for a given SourceType  by using the
	// <b>DescribeEventCategories</b> action.</p>
	EventCategories []*string
}

type ModifyEventSubscriptionOutput struct {
	// Contains the results of a successful invocation of the
	// DescribeEventSubscriptions () action.
	EventSubscription *types.EventSubscription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpModifyEventSubscriptionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpModifyEventSubscription{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpModifyEventSubscription{}, middleware.After)
}

func newServiceMetadataMiddleware_opModifyEventSubscription(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "ModifyEventSubscription",
	}
}
