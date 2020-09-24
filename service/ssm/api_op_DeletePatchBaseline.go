// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes a patch baseline.
func (c *Client) DeletePatchBaseline(ctx context.Context, params *DeletePatchBaselineInput, optFns ...func(*Options)) (*DeletePatchBaselineOutput, error) {
	stack := middleware.NewStack("DeletePatchBaseline", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDeletePatchBaselineMiddlewares(stack)
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
	addOpDeletePatchBaselineValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeletePatchBaseline(options.Region), middleware.Before)
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
			OperationName: "DeletePatchBaseline",
			Err:           err,
		}
	}
	out := result.(*DeletePatchBaselineOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeletePatchBaselineInput struct {
	// The ID of the patch baseline to delete.
	BaselineId *string
}

type DeletePatchBaselineOutput struct {
	// The ID of the deleted patch baseline.
	BaselineId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDeletePatchBaselineMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDeletePatchBaseline{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeletePatchBaseline{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeletePatchBaseline(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "DeletePatchBaseline",
	}
}
