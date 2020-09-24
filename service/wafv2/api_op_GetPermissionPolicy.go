// Code generated by smithy-go-codegen DO NOT EDIT.

package wafv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the IAM policy that is attached to the specified rule group. You must be
// the owner of the rule group to perform this operation.
func (c *Client) GetPermissionPolicy(ctx context.Context, params *GetPermissionPolicyInput, optFns ...func(*Options)) (*GetPermissionPolicyOutput, error) {
	stack := middleware.NewStack("GetPermissionPolicy", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetPermissionPolicyMiddlewares(stack)
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
	addOpGetPermissionPolicyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetPermissionPolicy(options.Region), middleware.Before)
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
			OperationName: "GetPermissionPolicy",
			Err:           err,
		}
	}
	out := result.(*GetPermissionPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetPermissionPolicyInput struct {
	// The Amazon Resource Name (ARN) of the rule group for which you want to get the
	// policy.
	ResourceArn *string
}

type GetPermissionPolicyOutput struct {
	// The IAM policy that is attached to the specified rule group.
	Policy *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetPermissionPolicyMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetPermissionPolicy{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetPermissionPolicy{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetPermissionPolicy(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "wafv2",
		OperationName: "GetPermissionPolicy",
	}
}
