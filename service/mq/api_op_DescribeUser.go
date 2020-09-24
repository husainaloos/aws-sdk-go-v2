// Code generated by smithy-go-codegen DO NOT EDIT.

package mq

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mq/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns information about an ActiveMQ user.
func (c *Client) DescribeUser(ctx context.Context, params *DescribeUserInput, optFns ...func(*Options)) (*DescribeUserOutput, error) {
	stack := middleware.NewStack("DescribeUser", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDescribeUserMiddlewares(stack)
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
	addOpDescribeUserValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeUser(options.Region), middleware.Before)
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
			OperationName: "DescribeUser",
			Err:           err,
		}
	}
	out := result.(*DescribeUserOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeUserInput struct {
	// The username of the ActiveMQ user. This value can contain only alphanumeric
	// characters, dashes, periods, underscores, and tildes (- . _ ~). This value must
	// be 2-100 characters long.
	Username *string
	// The unique ID that Amazon MQ generates for the broker.
	BrokerId *string
}

type DescribeUserOutput struct {
	// Enables access to the the ActiveMQ Web Console for the ActiveMQ user.
	ConsoleAccess *bool
	// The list of groups (20 maximum) to which the ActiveMQ user belongs. This value
	// can contain only alphanumeric characters, dashes, periods, underscores, and
	// tildes (- . _ ~). This value must be 2-100 characters long.
	Groups []*string
	// Required. The username of the ActiveMQ user. This value can contain only
	// alphanumeric characters, dashes, periods, underscores, and tildes (- . _ ~).
	// This value must be 2-100 characters long.
	Username *string
	// Required. The unique ID that Amazon MQ generates for the broker.
	BrokerId *string
	// The status of the changes pending for the ActiveMQ user.
	Pending *types.UserPendingChanges

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDescribeUserMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDescribeUser{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeUser{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeUser(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mq",
		OperationName: "DescribeUser",
	}
}
