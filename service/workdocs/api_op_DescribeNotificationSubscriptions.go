// Code generated by smithy-go-codegen DO NOT EDIT.

package workdocs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/workdocs/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the specified notification subscriptions.
func (c *Client) DescribeNotificationSubscriptions(ctx context.Context, params *DescribeNotificationSubscriptionsInput, optFns ...func(*Options)) (*DescribeNotificationSubscriptionsOutput, error) {
	stack := middleware.NewStack("DescribeNotificationSubscriptions", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDescribeNotificationSubscriptionsMiddlewares(stack)
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
	addOpDescribeNotificationSubscriptionsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeNotificationSubscriptions(options.Region), middleware.Before)
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
			OperationName: "DescribeNotificationSubscriptions",
			Err:           err,
		}
	}
	out := result.(*DescribeNotificationSubscriptionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeNotificationSubscriptionsInput struct {
	// The marker for the next set of results. (You received this marker from a
	// previous call.)
	Marker *string
	// The ID of the organization.
	OrganizationId *string
	// The maximum number of items to return with this call.
	Limit *int32
}

type DescribeNotificationSubscriptionsOutput struct {
	// The subscriptions.
	Subscriptions []*types.Subscription
	// The marker to use when requesting the next set of results. If there are no
	// additional results, the string is empty.
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDescribeNotificationSubscriptionsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDescribeNotificationSubscriptions{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeNotificationSubscriptions{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeNotificationSubscriptions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workdocs",
		OperationName: "DescribeNotificationSubscriptions",
	}
}
