// Code generated by smithy-go-codegen DO NOT EDIT.

package shield

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/shield/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the details of a Protection () object.
func (c *Client) DescribeProtection(ctx context.Context, params *DescribeProtectionInput, optFns ...func(*Options)) (*DescribeProtectionOutput, error) {
	if params == nil {
		params = &DescribeProtectionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeProtection", params, optFns, addOperationDescribeProtectionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeProtectionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeProtectionInput struct {

	// The unique identifier (ID) for the Protection () object that is described. When
	// submitting the DescribeProtection request you must provide either the
	// ResourceArn or the ProtectionID, but not both.
	ProtectionId *string

	// The ARN (Amazon Resource Name) of the AWS resource for the Protection () object
	// that is described. When submitting the DescribeProtection request you must
	// provide either the ResourceArn or the ProtectionID, but not both.
	ResourceArn *string
}

type DescribeProtectionOutput struct {

	// The Protection () object that is described.
	Protection *types.Protection

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeProtectionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeProtection{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeProtection{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeProtection(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeProtection(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "shield",
		OperationName: "DescribeProtection",
	}
}
