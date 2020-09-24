// Code generated by smithy-go-codegen DO NOT EDIT.

package sqs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Remove cost allocation tags from the specified Amazon SQS queue. For an
// overview, see Tagging Your Amazon SQS Queues
// (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-queue-tags.html)
// in the Amazon Simple Queue Service Developer Guide. Cross-account permissions
// don't apply to this action. For more information, see Grant Cross-Account
// Permissions to a Role and a User Name
// (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-customer-managed-policy-examples.html#grant-cross-account-permissions-to-role-and-user-name)
// in the Amazon Simple Queue Service Developer Guide.
func (c *Client) UntagQueue(ctx context.Context, params *UntagQueueInput, optFns ...func(*Options)) (*UntagQueueOutput, error) {
	stack := middleware.NewStack("UntagQueue", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpUntagQueueMiddlewares(stack)
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
	addOpUntagQueueValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUntagQueue(options.Region), middleware.Before)
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
			OperationName: "UntagQueue",
			Err:           err,
		}
	}
	out := result.(*UntagQueueOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UntagQueueInput struct {
	// The URL of the queue.
	QueueUrl *string
	// The list of tags to be removed from the specified queue.
	TagKeys []*string
}

type UntagQueueOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpUntagQueueMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpUntagQueue{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpUntagQueue{}, middleware.After)
}

func newServiceMetadataMiddleware_opUntagQueue(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sqs",
		OperationName: "UntagQueue",
	}
}
