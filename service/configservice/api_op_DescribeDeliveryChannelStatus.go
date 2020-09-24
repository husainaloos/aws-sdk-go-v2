// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the current status of the specified delivery channel. If a delivery
// channel is not specified, this action returns the current status of all delivery
// channels associated with the account. Currently, you can specify only one
// delivery channel per region in your account.
func (c *Client) DescribeDeliveryChannelStatus(ctx context.Context, params *DescribeDeliveryChannelStatusInput, optFns ...func(*Options)) (*DescribeDeliveryChannelStatusOutput, error) {
	stack := middleware.NewStack("DescribeDeliveryChannelStatus", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeDeliveryChannelStatusMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDeliveryChannelStatus(options.Region), middleware.Before)
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
			OperationName: "DescribeDeliveryChannelStatus",
			Err:           err,
		}
	}
	out := result.(*DescribeDeliveryChannelStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for the DeliveryChannelStatus () action.
type DescribeDeliveryChannelStatusInput struct {
	// A list of delivery channel names.
	DeliveryChannelNames []*string
}

// The output for the DescribeDeliveryChannelStatus () action.
type DescribeDeliveryChannelStatusOutput struct {
	// A list that contains the status of a specified delivery channel.
	DeliveryChannelsStatus []*types.DeliveryChannelStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeDeliveryChannelStatusMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeDeliveryChannelStatus{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeDeliveryChannelStatus{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeDeliveryChannelStatus(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "config",
		OperationName: "DescribeDeliveryChannelStatus",
	}
}
