// Code generated by smithy-go-codegen DO NOT EDIT.

package globalaccelerator

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/globalaccelerator/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describe the attributes of an accelerator. To see an AWS CLI example of
// describing the attributes of an accelerator, scroll down to Example.
func (c *Client) DescribeAcceleratorAttributes(ctx context.Context, params *DescribeAcceleratorAttributesInput, optFns ...func(*Options)) (*DescribeAcceleratorAttributesOutput, error) {
	stack := middleware.NewStack("DescribeAcceleratorAttributes", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeAcceleratorAttributesMiddlewares(stack)
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
	addOpDescribeAcceleratorAttributesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAcceleratorAttributes(options.Region), middleware.Before)

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
			OperationName: "DescribeAcceleratorAttributes",
			Err:           err,
		}
	}
	out := result.(*DescribeAcceleratorAttributesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAcceleratorAttributesInput struct {
	// The Amazon Resource Name (ARN) of the accelerator with the attributes that you
	// want to describe.
	AcceleratorArn *string
}

type DescribeAcceleratorAttributesOutput struct {
	// The attributes of the accelerator.
	AcceleratorAttributes *types.AcceleratorAttributes

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeAcceleratorAttributesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeAcceleratorAttributes{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeAcceleratorAttributes{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeAcceleratorAttributes(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "globalaccelerator",
		OperationName: "DescribeAcceleratorAttributes",
	}
}